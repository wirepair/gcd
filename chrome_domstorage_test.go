package gcd

import (
	//"encoding/json"
	"fmt"
	"github.com/wirepair/gcd/gcdprotogen/types"
	//"log"
	"testing"
	"time"
)

func TestDOMStorage(t *testing.T) {
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
	storage := targets[0].DOMStorage()

	fmt.Printf("Sending enable...")
	resp, err := storage.Enable()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Printf("Sent enable... resp: %v\n", resp)
	time.Sleep(1 * time.Second)
	fmt.Printf("Sending GetDOMStorageItems...")
	i := &types.ChromeDOMStorageStorageId{SecurityOrigin: "www.google.com", IsLocalStorage: true}
	itemResp, _ := storage.GetDOMStorageItems(i)
	fmt.Printf("Sent GetDOMStorageItems... itemResp: %v\n", itemResp)
	time.Sleep(1 * time.Second)
	fmt.Printf("Sending Disable...")
	resp, _ = storage.Disable()
	fmt.Printf("Sent Disable... resp: %v\n", resp)
	time.Sleep(1 * time.Second)
}
