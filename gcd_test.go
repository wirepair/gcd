package gcd

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/wirepair/gcd/gcdapi"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	debugger       *Gcd
	testListener   net.Listener
	testPath       string
	testDir        string
	testPort       string
	testServerAddr string
)

func init() {
	flag.StringVar(&testPath, "chrome", "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "path to Xvfb")
	flag.StringVar(&testDir, "dir", "C:\\temp\\", "user directory")
	flag.StringVar(&testPort, "port", "9222", "Debugger port")
}

func TestMain(m *testing.M) {
	flag.Parse()
	testServer()

	ret := m.Run()
	testCleanUp()
	os.Exit(ret)
}

func testCleanUp() {
	testListener.Close()

}

type ConsoleMessageEventHeader struct {
	Method string                           `json:"method,omitempty"`
	Params *gcdapi.ConsoleMessageAddedEvent `json:"params,omitempty"`
}

func TestGetPages(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	targets, _ := debugger.GetTargets()
	if len(targets) <= 0 {
		t.Fatalf("invalid number of targets, got: %d\n", len(targets))
	}
	t.Logf("page: %s\n", targets[0].Target.Url)
}

func TestEnv(t *testing.T) {
	var ok bool
	debugger = NewChromeDebugger()
	debugger.AddEnvironmentVars([]string{"hello=youze", "zoinks=scoob"})
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
		msg := &ConsoleMessageEventHeader{}
		err := json.Unmarshal(v, msg)
		if err != nil {
			t.Fatalf("error unmarshalling event data: %v\n", err)
		}
		t.Logf("METHOD: %s\n", msg.Method)
		eventData := msg.Params.Message
		t.Logf("Got event: %v\n", eventData)
		t.Logf("Timestamp: %f\n", eventData.Timestamp)
		doneCh <- struct{}{}
	})

	_, err = console.Enable()
	if err != nil {
		t.Fatalf("error sending enable: %s\n", err)
	}

	if _, err := target.Page.Navigate(testServerAddr + "console_log.html"); err != nil {
		t.Fatalf("error attempting to navigate: %s\n", err)
	}

	go testTimeoutListener(t, 5, "console message")

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
	if _, err := network.Enable(); err != nil {
		t.Fatalf("error enabling network")
	}
	ret, err = network.CanClearBrowserCache()
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

	wg := &sync.WaitGroup{}
	wg.Add(1)
	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}
	if _, err := target.Network.Enable(); err != nil {
		t.Fatalf("error enabling network %s\n", err)
	}

	if _, err := target.Page.Enable(); err != nil {
		t.Fatalf("error enabling page: %s\n", err)
	}

	target.Subscribe("Page.loadEventFired", func(target *ChromeTarget, payload []byte) {
		var ok bool
		t.Logf("page load event fired\n")
		cookies, err := target.Network.GetCookies()
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
		wg.Done()
	})

	_, err = target.Page.Navigate(testServerAddr + "cookie.html")
	if err != nil {
		t.Fatalf("error navigating to cookie page: %s\n", err)
	}

	go testTimeoutListener(t, 7, "waiting for page load to get cookies")
	t.Logf("waiting for loadEventFired")
	wg.Wait()
}

// UTILITY FUNCTIONS

func testDefaultStartup(t *testing.T) {
	debugger = NewChromeDebugger()
	debugger.StartProcess(testPath, testRandomTempDir(t), testRandomPort(t))
}

func testServer() {
	testListener, _ = net.Listen("tcp", ":0")
	_, testServerPort, _ := net.SplitHostPort(testListener.Addr().String())
	testServerAddr = fmt.Sprintf("http://localhost:%s/", testServerPort)
	go http.Serve(testListener, http.FileServer(http.Dir("testdata/")))
}

func testTimeoutListener(t *testing.T, seconds time.Duration, message string) {
	timeout := time.NewTimer(seconds * time.Second)
	for {
		select {
		case <-timeout.C:
			t.Fatalf("timed out waiting for %s", message)
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
