package gcd

import (
	//"encoding/json"
	"fmt"
	//"github.com/wirepair/gcd/gcdprotogen/types"
	"sync"
	"testing"
	"time"
)

func TestPageFPS(t *testing.T) {
	debugger := NewChromeDebugger()
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess()
	time.Sleep(1 * time.Second)
	targets := debugger.GetTargets()
	if len(targets) <= 0 {
		t.Fatalf("invalid number of targets, got: %d\n", len(targets))
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("page: %s\n", targets[0].Target.Url)
	page := targets[0].Page()

	fmt.Printf("Sending enable...")
	_, err := page.SetShowFPSCounter(true)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	cookies, newErr := page.GetCookies()
	fmt.Printf("%v %v\n", cookies, newErr)
}

func TestPageNavigate(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	debugger := NewChromeDebugger()
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess()
	targets := debugger.GetTargets()
	target := targets[0]
	target.Subscribe("Page.loadEventFired", func(targ *ChromeTarget, v []byte) {
		wg.Done()
	})
	page := target.Page()
	page.Enable()
	ret, err := page.Navigate("http://www.veracode.com")
	if err != nil {
		t.Fatalf("Error navigating: %s\n", err)
	}
	fmt.Printf("ret: %#v\n", ret)
	wg.Wait()
}

func TestPageDeadLock(t *testing.T) {
	//var frameId *types.ChromePageFrameId
	var wg sync.WaitGroup
	wg.Add(1)
	debugger := NewChromeDebugger()
	// start process, specify a tmp profile path so we get a fresh profiled browser
	// set port 9222 as the debug port
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess()     // exit when done
	targets := debugger.GetTargets() // get the 'targets' or tabs/background processes
	target := targets[0]             // take the first one
	fmt.Printf("targetInfo: %v\n", target.Target)
	page := target.Page() // get the Page API

	//subscribe to page load
	target.Subscribe("Page.loadEventFired", func(targ *ChromeTarget, v []byte) {
		fmt.Printf("FIRED frameStoppedLoading EVENT")
		dom := target.DOM()
		nodeId, err := dom.GetNodeForLocation(0, 0)
		_, err = dom.SetOuterHTML(nodeId, "<html><body><h1>Veracodeの皆さん<br>おはようございます！Hackathon 6!</h1></body></html>")
		fmt.Printf("SetOuterHTML....\n")
		if err != nil {
			fmt.Printf("GOT ERR IN SETDOC: %s\n", err)
		}
		time.Sleep(3 * time.Second)
		wg.Done() // page loaded, we can exit now

	})
	target.Subscribe("Page.frameStoppedLoading", func(targ *ChromeTarget, v []byte) {

	})

	page.Enable()                                        // Enable so we can recv events
	ret, err := page.Navigate("http://www.veracode.com") // navigate
	if err != nil {
		t.Fatalf("Error navigating: %s\n", err)
	}

	//frameId = ret
	//page.AddScriptToEvaluateOnLoad("document.write('<h1>Veracodeの皆さん<br>おはよございます！Hackathon 6!</h1>')")
	fmt.Printf("\n\nret: %s\n\n", ret)

	wg.Wait() // wait for page load
}
