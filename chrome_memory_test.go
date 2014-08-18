package gcd

import (
	//"encoding/json"
	"fmt"
	//"log"
	"testing"
	"time"
)

func TestMemory(t *testing.T) {
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
	memory := targets[0].Memory()

	fmt.Printf("Sending getDOMCounters...")
	doc, nodes, event, err := memory.GetDOMCounters()
	if err != nil {
		t.Fatalf("error getting dom counters: %v\n", err)
	}
	fmt.Printf("doc: %d nodes: %d event: %d", doc, nodes, event)

}
