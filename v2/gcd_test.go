package gcd

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"testing"
	"time"

	"github.com/wirepair/gcd/v2/gcdapi"
)

var (
	debugger                 *Gcd
	testListener             net.Listener
	testSkipNetworkIntercept bool
	testPath                 string
	testDir                  string
	testPort                 string
	testServerAddr           string
	testCtx                  = context.Background()
)

func init() {
	switch runtime.GOOS {
	case "windows":
		flag.StringVar(&testPath, "chrome", "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe", "path to chrome")
		flag.StringVar(&testDir, "dir", "C:\\temp\\gcd\\", "user directory")
	case "darwin":
		flag.StringVar(&testPath, "chrome", "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", "path to chrome")
		flag.StringVar(&testDir, "dir", "/tmp/gcd/", "user directory")
	case "linux":
		flag.StringVar(&testPath, "chrome", "/usr/bin/chromium-browser", "path to chrome")
		flag.StringVar(&testDir, "dir", "/tmp/gcd/", "user directory")
	}
	flag.StringVar(&testPort, "port", "9222", "Debugger port")

}

func TestMain(m *testing.M) {
	flag.Parse()
	testServer()
	os.MkdirAll(testDir, 0755)
	ret := m.Run()
	testCleanUp()
	os.Exit(ret)
}

func testCleanUp() {
	testListener.Close()
	os.RemoveAll(testDir)
}

func TestDeleteProfileOnExit(t *testing.T) {
	if runtime.GOOS == "windows" {
		//t.Skip("windows will hold on to the process handle too long")
	}

	debugger := NewChromeDebugger(WithDeleteProfileOnExit(),
		WithFlags([]string{"--headless"}),
	)

	profileDir := testRandomTempDir(t)
	err := debugger.StartProcess(testPath, profileDir, testRandomPort(t))
	if err != nil {
		t.Fatalf("error starting chrome: %s\n", err)
	}
	debugger.ExitProcess()
	time.Sleep(3 * time.Second)
	if stat, err := os.Stat(profileDir); err == nil {
		t.Fatalf("error temporary profileDir still exists: %s\n", stat.Name())
	}
}

func TestGetPages(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	newTab, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error creating tab: %s\n", err)
	}

	targets, err := debugger.GetTargets()
	if err != nil {
		t.Fatalf("error getting targets: %s\n", err)
	}

	if len(targets) <= 0 {
		t.Fatalf("invalid number of targets, got: %d\n", len(targets))
	}

	err = debugger.CloseTab(newTab)
	if err != nil {
		t.Fatalf("error CloseTab: %s\n", err)
	}

	t.Logf("page: %s\n", targets[0].Target.Url)
}

func TestEnv(t *testing.T) {
	var ok bool
	debugger = NewChromeDebugger(WithEnvironmentVars([]string{"hello=youze", "zoinks=scoob"}), WithFlags([]string{"--headless"}))
	debugger.StartProcess(testPath, testRandomTempDir(t), testRandomPort(t))
	defer debugger.ExitProcess()

	t.Logf("%#v\n", debugger.chromeCmd.Env)
	for _, v := range debugger.chromeCmd.Env {
		if v == "hello=youze" {
			ok = true
		}
	}
	if !ok {
		t.Fatalf("error finding our environment vars in chrome process")
	}
}

func TestGoRoutineCount(t *testing.T) {
	startRoutines := runtime.NumGoroutine()

	testDefaultStartup(t)
	defer debugger.ExitProcess()

	_, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error creating new tab")
	}

	debugger.ExitProcess()
	// let exit process clean up profile directories and what not
	time.Sleep(time.Second * 2)
	remaining := runtime.NumGoroutine()
	if remaining > startRoutines {
		pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
		t.Fatalf("error expected remaining go routines (%d) to less than or equal start routines (%d)\n", remaining, startRoutines)
	}
}

func TestProcessKilled(t *testing.T) {
	doneCh := make(chan struct{})

	terminatedHandler := func(reason string) {
		t.Logf("reason: %s\n", reason)
		doneCh <- struct{}{}
	}

	testDefaultStartup(t, WithTerminationHandler(terminatedHandler))
	shutdown := time.NewTimer(time.Second * 4)
	timeout := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-doneCh:
			goto DONE
		case <-shutdown.C:
			debugger.ExitProcess()
		case <-timeout.C:
			t.Fatalf("timed out waiting for termination")
		}
	}
DONE:
}

func TestTargetCrashed(t *testing.T) {

	testDefaultStartup(t)
	defer debugger.ExitProcess()

	doneCh := make(chan struct{})
	go testTimeoutListener(t, doneCh, 5, "timed out waiting for crashed to be handled")

	targetCrashedFn := func(targ *ChromeTarget, payload []byte) {
		t.Logf("reason: %s\n", string(payload))
		close(doneCh)
	}

	tab, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error creating new tab")
	}

	tab.Inspector.Enable(testCtx)
	tab.Subscribe("Inspector.targetCrashed", targetCrashedFn)

	navParams := &gcdapi.PageNavigateParams{Url: "chrome://crash", TransitionType: "typed"}
	tab.Page.NavigateWithParams(testCtx, navParams)
	<-doneCh
}

func TestEvents(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}
	console := target.Console

	doneCh := make(chan struct{}, 1)
	target.Subscribe("Console.messageAdded", func(target *ChromeTarget, v []byte) {
		target.Unsubscribe("Console.messageAdded")
		msg := &gcdapi.ConsoleMessageAddedEvent{}
		err := json.Unmarshal(v, msg)
		if err != nil {
			t.Fatalf("error unmarshalling event data: %v\n", err)
		}
		t.Logf("METHOD: %s\n", msg.Method)
		eventData := msg.Params.Message
		t.Logf("Got event: %v\n", eventData)
		close(doneCh)
	})

	_, err = console.Enable(testCtx)
	if err != nil {
		t.Fatalf("error sending enable: %s\n", err)
	}

	navParams := &gcdapi.PageNavigateParams{Url: testServerAddr + "console_log.html", TransitionType: "typed"}
	if _, _, _, err := target.Page.NavigateWithParams(testCtx, navParams); err != nil {
		t.Fatalf("error attempting to navigate: %s\n", err)
	}

	go testTimeoutListener(t, doneCh, 5, "console message")

	<-doneCh
}

func TestEvaluate(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()
	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}

	doneCh := make(chan struct{}, 1)
	target.Subscribe("Runtime.executionContextCreated", func(target *ChromeTarget, v []byte) {
		//target.Unsubscribe("Console.messageAdded")
		msg := &gcdapi.RuntimeExecutionContextCreatedEvent{}
		err := json.Unmarshal(v, msg)
		if err != nil {
			t.Fatalf("error unmarshalling event data: %v\n", err)
		}

		if msg.Params.Context.Origin != testServerAddr[:len(testServerAddr)-1] {
			return
		}
		scriptSource := "document.location.href"
		objectGroup := "gcdtest"
		awaitPromise := false
		includeCommandLineAPI := true
		contextId := msg.Params.Context.Id
		silent := true
		returnByValue := false
		generatePreview := true
		userGestures := true
		evalParams := &gcdapi.RuntimeEvaluateParams{Expression: scriptSource, ObjectGroup: objectGroup, IncludeCommandLineAPI: includeCommandLineAPI, Silent: silent, ContextId: contextId, ReturnByValue: returnByValue, GeneratePreview: generatePreview, UserGesture: userGestures, AwaitPromise: awaitPromise}
		rro, exception, err := target.Runtime.EvaluateWithParams(testCtx, evalParams)
		if err != nil {
			t.Fatalf("error evaluating: %s %#v\n", err, exception)
		}

		if val, ok := rro.Value.(string); ok {
			if val != testServerAddr {
				t.Fatalf("invalid location returned expected %s got %s\n", testServerAddr, val)
			}
		} else {
			t.Fatalf("error rro.Value was not a string")
		}
		close(doneCh)
	})
	target.Runtime.Enable(testCtx)

	navParams := &gcdapi.PageNavigateParams{Url: testServerAddr, TransitionType: "typed"}
	target.Page.NavigateWithParams(testCtx, navParams)
	<-doneCh
}

func TestSimpleReturn(t *testing.T) {
	var ret bool
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}
	network := target.Network
	if _, err := network.Enable(testCtx, -1, -1, -1); err != nil {
		t.Fatalf("error enabling network")
	}
	ret, err = network.CanClearBrowserCache(testCtx)
	if err != nil {
		t.Fatalf("error getting response to clearing browser cache: %s\n", err)
	}
	if !ret {
		t.Fatalf("we should have got true for can clear browser cache\n")
	}
}

func TestSimpleReturnReturnsGoError(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}

	css := target.CSS
	if _, err := css.Enable(testCtx); err == nil {
		t.Fatalf("expected a go error since css.Enable requires DOM.Enable first.")
	}
}

func TestDOMEnableWithWhiteSpace(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()
	ctx := context.Background()
	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}

	page := target.Page
	if _, err := page.Enable(ctx); err != nil {
		t.Fatalf("failed to enable page: %s\n", err)
	}

	dom := target.DOM
	if _, err := dom.EnableWithParams(ctx, &gcdapi.DOMEnableParams{}); err != nil {
		t.Fatalf("got error enabling DOM: %s\n", err)
	}

	doneCh := make(chan struct{})
	target.Subscribe("Page.loadEventFired", func(target *ChromeTarget, payload []byte) {

		doc, err := target.DOM.GetDocument(context.Background(), -1, true)
		if err != nil {
			t.Fatalf("failed to get doc: %s\n", err)
		}

		if doc.ChildNodeCount != 2 {
			t.Fatalf("failed to get 2 child nodes, got %d\n", doc.ChildNodeCount)
		}
		doneCh <- struct{}{}
	})

	navParams := &gcdapi.PageNavigateParams{Url: testServerAddr + "cookie.html", TransitionType: "typed"}
	_, _, _, err = target.Page.NavigateWithParams(testCtx, navParams)
	if err != nil {
		t.Fatalf("failed to load test page: %s\n", err)
	}

	<-doneCh
}

// Tests that the ctx canceled doesn't cause the wsconn to get stuck in a loop in windows
func TestCtxCancel(t *testing.T) {
	testDefaultStartup(t, WithEventDebugging(), WithEventDebugging())
	defer debugger.ExitProcess()

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}

	network := target.Network
	if _, err := network.Enable(testCtx, -1, -1, -1); err != nil {
		t.Fatalf("error enabling network")
	}
	ctx, cancel := context.WithCancel(testCtx)
	cancel()
	if _, err := network.CanClearBrowserCache(ctx); err == nil {
		t.Fatal(err)
	}

	if _, err := network.CanClearBrowserCache(ctx); err == nil {
		t.Fatal(err)
	}
}

func TestSimpleReturnWithParams(t *testing.T) {
	var ret bool
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}
	network := target.Network

	networkParams := &gcdapi.NetworkEnableParams{
		MaxTotalBufferSize:    -1,
		MaxResourceBufferSize: -1,
	}

	if _, err := network.EnableWithParams(testCtx, networkParams); err != nil {
		t.Fatalf("error enabling network")
	}
	ret, err = network.CanClearBrowserCache(testCtx)
	if err != nil {
		t.Fatalf("error getting response to clearing browser cache: %s\n", err)
	}
	if !ret {
		t.Fatalf("we should have got true for can clear browser cache\n")
	}
}

// tests getting a complex object back from inside a fired event that we subscribed to.
func TestComplexReturn(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	doneCh := make(chan struct{}, 1)
	go testTimeoutListener(t, doneCh, 7, "waiting for page load to get cookies")
	target, err := debugger.NewTab()

	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}
	if _, err := target.Network.Enable(testCtx, -1, -1, -1); err != nil {
		t.Fatalf("error enabling network %s\n", err)
	}

	if _, err := target.Page.Enable(testCtx); err != nil {
		t.Fatalf("error enabling page: %s\n", err)
	}

	target.Subscribe("Page.loadEventFired", func(target *ChromeTarget, payload []byte) {
		var ok bool
		t.Logf("page load event fired\n")
		cookies, err := target.Network.GetCookies(testCtx, []string{testServerAddr})
		if err != nil {
			t.Fatalf("error getting cookies!")
		}
		for _, v := range cookies {
			t.Logf("got cookies: %#v\n", v)
			if v.Name == "HEYA" {
				ok = true
				break
			}
		}
		if !ok {
			t.Fatalf("error finding our cookie value!")
		}
		close(doneCh)
	})

	navParams := &gcdapi.PageNavigateParams{Url: testServerAddr + "cookie.html", TransitionType: "typed"}
	_, _, _, err = target.Page.NavigateWithParams(testCtx, navParams)
	if err != nil {
		t.Fatalf("error navigating to cookie page: %s\n", err)
	}

	t.Logf("waiting for loadEventFired")
	<-doneCh
}

func TestConnectToInstance(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	doneCh := make(chan struct{})

	go testTimeoutListener(t, doneCh, 15, "timed out waiting for remote connection")
	go func() {
		remoteDebugger := NewChromeDebugger()
		remoteDebugger.ConnectToInstance(debugger.host, debugger.port)

		_, err := remoteDebugger.NewTab()
		if err != nil {
			t.Fatalf("error creating new tab")
		}

		targets, error := remoteDebugger.GetTargets()
		if error != nil {
			t.Fatalf("cannot get targets: %s \n", error)
		}
		if len(targets) <= 0 {
			t.Fatalf("invalid number of targets, got: %d\n", len(targets))
		}
		for _, target := range targets {
			t.Logf("page: %s\n", target.Target.Url)
		}
		close(doneCh)
	}()
	<-doneCh
}

func TestLocalExtension(t *testing.T) {
	testExtensionStartup(t)
	if err := debugger.ConnectToInstance(debugger.host, debugger.port); err != nil {
		t.Fatalf("failed to connect: %s\n", err)
	}

	defer debugger.ExitProcess()

	doneCh := make(chan struct{})

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error creating new tab")
	}

	if _, err := target.Page.Enable(testCtx); err != nil {
		t.Fatalf("error enabling page: %s\n", err)
	}

	target.Subscribe("Page.loadEventFired", func(target *ChromeTarget, payload []byte) {
		t.Logf("page load event fired\n")
		close(doneCh)
	})

	if _, err := target.Network.Enable(testCtx, -1, -1, -1); err != nil {
		t.Fatalf("error enabling network: %s\n", err)
	}

	params := &gcdapi.PageNavigateParams{Url: "http://www.google.com"}
	_, _, _, err = target.Page.NavigateWithParams(testCtx, params)
	if err != nil {
		t.Fatalf("error navigating: %s\n", err)
	}

	go testTimeoutListener(t, doneCh, 15, "timed out waiting for remote connection")
	<-doneCh
}

func TestContextCancel(t *testing.T) {
	testDefaultStartup(t, WithEventDebugging(), WithInternalDebugMessages(), WithLogger(&DebugLogger{}))
	defer debugger.ExitProcess()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting first tab")
	}

	if _, err := target.Page.Enable(testCtx); err != nil {
		t.Fatalf("error enabling page: %s\n", err)
	}

	if _, err := target.Page.WaitForDebugger(ctx); err == nil {
		t.Fatalf("did not get cancelation")
	}
}

func TestNetworkIntercept(t *testing.T) {

	testDefaultStartup(t, WithEventDebugging(), WithInternalDebugMessages(), WithLogger(&DebugLogger{}))
	defer debugger.ExitProcess()

	doneCh := make(chan struct{})

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}

	go testTimeoutListener(t, doneCh, 5, "timed out waiting for requestIntercepted")
	ctx := context.Background()
	if _, err := target.Fetch.Enable(ctx, []*gcdapi.FetchRequestPattern{
		{
			UrlPattern:   "*",
			RequestStage: "Request",
		},
	}, false); err != nil {
		t.Fatalf("error enabling fetch: %s", err)
	}

	target.Subscribe("Fetch.requestPaused", func(target *ChromeTarget, payload []byte) {
		close(doneCh)

		pausedEvent := &gcdapi.FetchRequestPausedEvent{}
		if err := json.Unmarshal(payload, pausedEvent); err != nil {
			t.Fatalf("error unmarshal: %s\n", err)
		}
		requestHeaders := pausedEvent.Params.Request.Headers
		fetchHeaders := make([]*gcdapi.FetchHeaderEntry, 0)
		for k, v := range requestHeaders {
			value := ""
			switch t := v.(type) {
			case string:
				value = t
			case []string:
				value = strings.Join(t, "")
			}
			fetchHeaders = append(fetchHeaders, &gcdapi.FetchHeaderEntry{Name: k, Value: value})
		}

		p := &gcdapi.FetchContinueRequestParams{
			RequestId: pausedEvent.Params.RequestId,
			Url:       pausedEvent.Params.Request.Url,
			Method:    pausedEvent.Method,
			PostData:  pausedEvent.Params.Request.PostData,
			Headers:   fetchHeaders,
		}
		target.Fetch.ContinueRequestWithParams(ctx, p)
	})

	params := &gcdapi.PageNavigateParams{Url: "http://www.example.com"}
	_, _, _, err = target.Page.NavigateWithParams(testCtx, params)
	if err != nil {
		t.Fatalf("error navigating: %s\n", err)
	}

	<-doneCh
}

func TestCloseTab(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()
	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting first tab: %v\n", err)
	}
	if err := debugger.CloseTab(target); err != nil {
		t.Fatalf("error closing tab")
	}
}

type testLogger struct {
	called bool
}

func (t *testLogger) Println(msg ...interface{}) {
	t.called = true
}

func TestCustomLogger(t *testing.T) {
	customLogger := &testLogger{}
	testDefaultStartup(t, WithLogger(customLogger), WithEventDebugging())
	defer debugger.ExitProcess()
	tab, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %v\n", err)
	}

	if _, err = tab.Page.Enable(context.TODO()); err != nil {
		t.Fatalf("error using custom logger")
	}

	if !customLogger.called {
		t.Fatalf("custom logger was not called")
	}
}

// UTILITY FUNCTIONS
func testExtensionStartup(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatalf("error getting working directory: %s\n", err)
	}

	sep := string(os.PathSeparator)
	extensionPath := "--load-extension=" + path + sep + "testdata" + sep + "extension" + sep
	debugger = NewChromeDebugger(WithFlags([]string{extensionPath, "--headless"}))

	debugger.StartProcess(testPath, testRandomTempDir(t), testRandomPort(t))
}

func testDefaultStartup(t *testing.T, opts ...func(*Gcd)) {
	mergedOpts := make([]func(*Gcd), 0)
	mergedOpts = append(mergedOpts, WithDeleteProfileOnExit())
	mergedOpts = append(mergedOpts, WithFlags([]string{"--headless"}))
	mergedOpts = append(mergedOpts, opts...)
	debugger = NewChromeDebugger(mergedOpts...)

	debugger.StartProcess(testPath, testRandomTempDir(t), testRandomPort(t))
}

func testServer() {
	testListener, _ = net.Listen("tcp", ":0")
	_, testServerPort, _ := net.SplitHostPort(testListener.Addr().String())
	testServerAddr = fmt.Sprintf("http://localhost:%s/", testServerPort)
	go http.Serve(testListener, http.FileServer(http.Dir("testdata/")))
}

func testTimeoutListener(t *testing.T, closeCh chan struct{}, seconds time.Duration, message string) {
	timeout := time.NewTimer(seconds * time.Second)
	for {
		select {
		case <-closeCh:
			timeout.Stop()
			return
		case <-timeout.C:
			close(closeCh)
			t.Fatalf("timed out waiting for %s", message)
			return
		}
	}
}

func testRandomPort(t *testing.T) string {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	_, randPort, _ := net.SplitHostPort(l.Addr().String())
	l.Close()
	return randPort
}

func testRandomTempDir(t *testing.T) string {
	dir, err := ioutil.TempDir(testDir, "gcd")
	if err != nil {
		t.Errorf("error creating temp dir: %s\n", err)
	}
	return dir
}
