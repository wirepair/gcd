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
	pages := debugger.GetPages()
	if len(pages) <= 0 {
		t.Fatalf("invalid number of pages, got: %d\n", len(pages))
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("page: %s\n", pages[0].Page.Url)
}

func TestConsole(t *testing.T) {
	debugger := NewChromeDebugger()
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess()
	time.Sleep(1 * time.Second)
	pages := debugger.GetPages()
	if len(pages) <= 0 {
		t.Fatalf("invalid number of pages, got: %d\n", len(pages))
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("page: %s\n", pages[0].Page.Url)
	console := pages[0].Console
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
