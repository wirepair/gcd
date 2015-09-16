package main

import (
	"github.com/wirepair/gcd"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	debugger := gcd.NewChromeDebugger()
	// start process, specify a tmp profile path so we get a fresh profiled browser
	// set port 9222 as the debug port
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess()          // exit when done
	targets, err := debugger.GetTargets() // get the 'targets' or tabs/background processes
	if err != nil {
		log.Fatalf("error getting targets: %s\n", err)
	}
	target := targets[0] // take the first one

	//subscribe to page load
	target.Subscribe("Page.loadEventFired", func(targ *gcd.ChromeTarget, v []byte) {
		wg.Done() // page loaded, we can exit now
		// if you wanted to inspect the full response data, you could do that here
	})
	// get the Page API and enable it
	if _, err := target.Page.Enable(); err != nil {
		log.Fatalf("error getting page: %s\n", err)
	}
	ret, err := target.Page.Navigate("http://www.veracode.com") // navigate
	if err != nil {
		log.Fatalf("Error navigating: %s\n", err)
	}
	log.Printf("ret: %#v\n", ret)
	wg.Wait() // wait for page load
}
