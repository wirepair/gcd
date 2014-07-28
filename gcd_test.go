package gcd

import (
	"fmt"
	"testing"
	"time"
)

func TestChromeDebugger(t *testing.T) {
	debugger := NewChromeDebugger()
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	time.Sleep(2 * time.Second)
	debugger.ExitProcess()
}

func TestGetPages(t *testing.T) {
	debugger := NewChromeDebugger()
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess()
	time.Sleep(2 * time.Second)
	targets := debugger.GetTargets()
	if len(targets) <= 0 {
		t.Fatalf("invalid number of targets, got: %d\n", len(targets))
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("page: %s\n", targets[0].Target.Url)
}

func TestConsole(t *testing.T) {
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
	console := targets[0].Console()
	fmt.Printf("Sending enable...")
	console.Enable()
	fmt.Printf("Sent enable...")
	time.Sleep(1 * time.Second)
	fmt.Printf("Sending ClearMessages...")
	console.ClearMessages()
	fmt.Printf("Sent ClearMessages...")
	time.Sleep(1 * time.Second)
	fmt.Printf("Sending Disable...")
	console.Disable()
	fmt.Printf("Sent Disable...")
	time.Sleep(1 * time.Second)
}
