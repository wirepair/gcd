package main

import (
	"github.com/wirepair/gcd"
	"log"
)

func main() {
	debugger := gcd.NewChromeDebugger()
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess()
	target, err := debugger.NewTab()
	if err != nil {
		log.Fatalf("error getting new tab: %s\n", err)
	}
	dom := target.DOM
	r, err := dom.GetDocument()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	log.Printf("NodeID: %d Node Name: %s, URL: %s\n", r.NodeId, r.NodeName, r.DocumentURL)
}
