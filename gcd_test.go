package gcd

import (
	"encoding/json"
	"fmt"
	"github.com/wirepair/gcd/gcdprotogen/types"
	"log"
	"testing"
	"time"
)

type EventData struct {
	Method string         `json:"method"`
	Params *ConsoleParams `json:"params"`
}

type ConsoleParams struct {
	Message *types.ChromeConsoleConsoleMessage `json:"message"`
}

const (
	path = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	dir  = "C:\\tmp\\"
	port = "9222"
)

func TestChromeDebugger(t *testing.T) {
	debugger := NewChromeDebugger()
	debugger.StartProcess(path, dir, port)
	time.Sleep(2 * time.Second)
	debugger.ExitProcess()
}

func TestGetPages(t *testing.T) {
	debugger := NewChromeDebugger()
	debugger.StartProcess(path, dir, port)
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
	debugger.StartProcess(path, dir, port)
	defer debugger.ExitProcess()
	time.Sleep(1 * time.Second)
	targets := debugger.GetTargets()
	if len(targets) <= 0 {
		t.Fatalf("invalid number of targets, got: %d\n", len(targets))
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("page: %s\n", targets[0].Target.Url)
	console := targets[0].Console()

	targets[0].Subscribe("Console.messageAdded", func(target *ChromeTarget, v []byte) {
		msg := &EventData{}
		err := json.Unmarshal(v, msg)
		if err != nil {
			log.Fatalf("error unmarshalling event data: %v\n", err)
		}
		log.Printf("METHOD: %s\n", msg.Method)
		eventData := msg.Params.Message
		fmt.Printf("Got event: %v\n", eventData)
		fmt.Printf("Timestamp: %f\n", *eventData.Timestamp)
	})

	fmt.Printf("Sending enable...")
	resp, err := console.Enable()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Printf("Sent enable... resp: %v\n", resp)
	time.Sleep(1 * time.Second)
	fmt.Printf("Sending ClearMessages...")
	resp, _ = console.ClearMessages()
	fmt.Printf("Sent ClearMessages... resp: %v\n", resp)
	time.Sleep(1 * time.Second)
	fmt.Printf("Sending Disable...")
	resp, _ = console.Disable()
	fmt.Printf("Sent Disable... resp: %v\n", resp)
	time.Sleep(1 * time.Second)
}

func TestNewTab(t *testing.T) {
	debugger := NewChromeDebugger()
	debugger.StartProcess(path, dir, port)
	defer debugger.ExitProcess()
	debugger.NewTab()
	time.Sleep(2 * time.Second)
}

func TestCloseTab(t *testing.T) {
	debugger := NewChromeDebugger()
	debugger.StartProcess(path, dir, port)
	defer debugger.ExitProcess()
	target := debugger.NewTab()
	time.Sleep(1 * time.Second)
	debugger.CloseTab(target)
	time.Sleep(2 * time.Second)
}
