package gcd

import (
	//"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
)

type ChromeDebugger struct {
	sync.RWMutex  // for locking pages (i.e. websocket clients)
	Targets       []*ChromeTarget
	chromeProcess *os.Process
	port          string
}

func NewChromeDebugger() *ChromeDebugger {
	c := &ChromeDebugger{}
	c.Targets = make([]*ChromeTarget, 0)
	return c
}

func (c *ChromeDebugger) StartProcess(exePath, userDir, port string) {
	c.port = port
	dir := fmt.Sprintf("--user-data-dir=%s", userDir)
	debugPort := fmt.Sprintf("--remote-debugging-port=%s", port)
	cmd := exec.Command(exePath, dir, debugPort)
	go func() {
		err := cmd.Start()
		if err != nil {
			log.Fatalf("error starting chrome process: %s", err)
		}
		c.chromeProcess = cmd.Process
		err = cmd.Wait()
		log.Printf("Chrome exited")
	}()
}

func (c *ChromeDebugger) ExitProcess() error {
	return c.chromeProcess.Kill()
}

func (c *ChromeDebugger) GetTargets() []*ChromeTarget {
	resp, err := http.Get(fmt.Sprintf("http://localhost:%s/json", c.port))
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	defer resp.Body.Close()
	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		log.Fatalf("error reading body: %v\n", errRead)
	}
	targets := make([]*TargetInfo, 0)
	err = json.Unmarshal(body, &targets)
	if err != nil {
		fmt.Printf("body: %s\n", string(body))
		log.Fatalf("error decoding inspectable pages: %v\n", err)
	}
	for _, v := range targets {
		target := newChromeTarget(c.port, v)
		c.Targets = append(c.Targets, target)
	}
	return c.Targets
}
