package gcd

import (
	//"encoding/json"
	"fmt"
	//"github.com/wirepair/gcd/gcdprotogen/types"
	"sync"
	"testing"
	"time"
)

/* see gcd_test for consts.
const (
	path = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.path"
	dir  = "C:\\tmp\\"
	port = "9222"
)
*/

func TestPageNavigate(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	debugger := NewChromeDebugger()
	debugger.StartProcess(path, dir, port)
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
	debugger.StartProcess(path, dir, port)
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
	fmt.Printf("\n\nret: %s\n\n", ret)
	wg.Wait() // wait for page load
}
