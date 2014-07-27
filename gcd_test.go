package gcd

import (
	"testing"
	"time"
)

func TestChromeDebugger(t *testing.T) {
	debugger := NewChromeDebugger()
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	time.Sleep(5 * time.Second)
	debugger.ExitProcess()
}

func TestConnectChromeDebugger(t *testing.T) {
	debugger := NewChromeDebugger()
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess()
	time.Sleep(3 * time.Second)
	debugger.Connect()
	time.Sleep(2 * time.Second)
}
