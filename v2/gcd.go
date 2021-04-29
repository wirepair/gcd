/*
The MIT License (MIT)

Copyright (c) 2020 isaac dawson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package gcd

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/wirepair/gcd/v2/gcdapi"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var GCDVERSION = "v2.1.1"

var (
	ErrNoTabAvailable = errors.New("no available tab found")
)

// When we get an error reading the body from the debugger api endpoint
type GcdBodyReadErr struct {
	Message string
}

func (g *GcdBodyReadErr) Error() string {
	return "error reading response body: " + g.Message
}

// Failure to unmarshal the JSON response from debugger API
type GcdDecodingErr struct {
	Message string
}

func (g *GcdDecodingErr) Error() string {
	return "error decoding inspectable page: " + g.Message
}

type TerminatedHandler func(reason string)

// The Google Chrome Debugger
type Gcd struct {
	processLock *sync.RWMutex // make go -race detection happy for process

	timeout           time.Duration // how much time to wait for debugger port to open up
	chromeProcess     *os.Process
	chromeCmd         *exec.Cmd
	terminatedHandler TerminatedHandler
	port              string
	host              string
	addr              string
	profileDir        string
	deleteProfile     bool
	readyCh           chan struct{}
	apiEndpoint       string
	flags             []string
	env               []string
	chomeApiVersion   string
	ctx               context.Context
	logger            Log
	debugEvents       bool
	debug             bool
}

// Give it a friendly name.
func NewChromeDebugger(opts ...func(*Gcd)) *Gcd {
	c := &Gcd{processLock: &sync.RWMutex{}}
	c.timeout = 15
	c.host = "localhost"
	c.readyCh = make(chan struct{})
	c.terminatedHandler = nil
	c.flags = make([]string, 0)
	c.env = make([]string, 0)
	c.ctx = context.Background()
	c.logger = LogDiscarder{}

	for _, o := range opts {
		o(c)
	}
	return c
}

// WithTerminationHandler Pass a handler to be notified when the chrome process exits.
func WithTerminationHandler(handler TerminatedHandler) func(*Gcd) {
	return func(g *Gcd) {
		g.terminatedHandler = handler
	}
}

// WithDebugPortTimeout for how long we should wait for debug port to become available.
func WithDebugPortTimeout(timeout time.Duration) func(*Gcd) {
	return func(g *Gcd) {
		g.timeout = timeout
	}
}

// WithFlags allows caller to add additional startup flags to the chrome process
func WithFlags(flags []string) func(*Gcd) {
	return func(g *Gcd) {
		g.flags = append(g.flags, flags...)
	}
}

// WithEnvironmentVars for the chrome process, useful for Xvfb etc.
func WithEnvironmentVars(vars []string) func(*Gcd) {
	return func(g *Gcd) {
		g.env = append(g.env, vars...)
	}
}

func WithDeleteProfileOnExit() func(*Gcd) {
	return func(g *Gcd) {
		g.deleteProfile = true
	}
}

func WithLogger(l Log) func(*Gcd) {
	return func(g *Gcd) {
		g.logger = l
	}
}

func WithContext(ctx context.Context) func(*Gcd) {
	return func(g *Gcd) {
		g.ctx = ctx
	}
}

func WithEventDebugging() func(*Gcd) {
	return func(g *Gcd) {
		g.debugEvents = true
	}
}

func WithInternalDebugMessages() func(*Gcd) {
	return func(g *Gcd) {
		g.debug = true
	}
}

// Port that the debugger is listening on
func (c *Gcd) Port() string {
	return c.port
}

// Host that the debugger is listening on
func (c *Gcd) Host() string {
	return c.host
}

// StartProcess the process
// exePath - the path to the executable
// userDir - the user directory to start from so we get a fresh profile
// port - The port to listen on.
func (c *Gcd) StartProcess(exePath, userDir, port string) error {
	c.port = port
	c.addr = fmt.Sprintf("%s:%s", c.host, c.port)
	c.profileDir = userDir
	c.apiEndpoint = fmt.Sprintf("http://%s/json", c.addr)
	// profile directory
	c.flags = append(c.flags, fmt.Sprintf("--user-data-dir=%s", c.profileDir))
	// debug port to use
	c.flags = append(c.flags, fmt.Sprintf("--remote-debugging-port=%s", port))
	// bypass first run check
	c.flags = append(c.flags, "--no-first-run")
	// bypass default browser check
	c.flags = append(c.flags, "--no-default-browser-check")

	c.chromeCmd = exec.Command(exePath, c.flags...)
	// add custom environment variables.
	c.chromeCmd.Env = os.Environ()
	c.chromeCmd.Env = append(c.chromeCmd.Env, c.env...)

	return c.startProcess()
}

// StartProcessCustom lets you pass in the exec.Cmd to use
func (c *Gcd) StartProcessCustom(cmd *exec.Cmd, userDir, port string) error {
	c.port = port
	c.addr = fmt.Sprintf("%s:%s", c.host, c.port)
	c.profileDir = userDir
	c.apiEndpoint = fmt.Sprintf("http://%s/json", c.addr)
	c.chromeCmd = cmd

	return c.startProcess()
}

// startProcess starts the process and waits for the debugger port to be ready
func (c *Gcd) startProcess() error {
	go func() {
		err := c.chromeCmd.Start()
		if err != nil {
			msg := fmt.Sprintf("failed to start chrome: %s", err)
			if c.terminatedHandler != nil {
				c.terminatedHandler(msg)
			} else {
				c.logger.Println(msg)
			}
		}
		c.processLock.Lock()
		c.chromeProcess = c.chromeCmd.Process
		c.processLock.Unlock()
		err = c.chromeCmd.Wait()

		c.removeProfileDir()

		closeMessage := "exited"
		if err != nil {
			closeMessage = err.Error()
		}
		if c.terminatedHandler != nil {
			c.terminatedHandler(closeMessage)
		}
	}()

	var err error
	go func(err error) {
		err = c.probeDebugPort()
	}(err)
	<-c.readyCh

	return err
}

// ExitProcess kills the process
func (c *Gcd) ExitProcess() error {
	c.processLock.Lock()
	err := c.chromeProcess.Kill()
	c.processLock.Unlock()
	return err
}

// PID of the started process
func (c *Gcd) PID() int {
	c.processLock.Lock()
	pid := c.chromeProcess.Pid
	c.processLock.Unlock()
	return pid
}

// removeProfileDir if deleteProfile is true
func (c *Gcd) removeProfileDir() {
	if c.deleteProfile {
		// let chrome shutdown completely before deleting
		time.Sleep(1 * time.Second)
		if err := os.RemoveAll(c.profileDir); err != nil {
			c.logger.Println("error deleting profile directory", err)
		}
	}
}

// ConnectToInstance connects to a running chrome instance without starting a local process
// Host - The host destination.
// Port - The port to listen on.
func (c *Gcd) ConnectToInstance(host string, port string) error {
	c.host = host
	c.port = port
	c.addr = fmt.Sprintf("%s:%s", c.host, c.port)
	c.apiEndpoint = fmt.Sprintf("http://%s/json", c.addr)

	var err error
	go func(err error) {
		err = c.probeDebugPort()
	}(err)
	<-c.readyCh

	return err
}

// GetTargets primary tabs/processes to work with. Each will have their own references
// to the underlying API components (such as Page, Debugger, DOM etc).
func (c *Gcd) GetTargets() ([]*ChromeTarget, error) {
	empty := make(map[string]struct{}, 0)
	return c.GetNewTargets(empty)
}

// GetNewTargets gets a list of current tabs and creates new chrome targets returning a list
// provided they weren't in the knownIds list. Note it is an error to attempt
// to create a new chrome target from one that already exists.
func (c *Gcd) GetNewTargets(knownIds map[string]struct{}) ([]*ChromeTarget, error) {
	connectableTargets, err := c.getConnectableTargets()
	if err != nil {
		return nil, err
	}

	chromeTargets := make([]*ChromeTarget, 0)
	for _, connectableTarget := range connectableTargets {
		if _, ok := knownIds[connectableTarget.Id]; !ok {
			target, err := openChromeTarget(c, connectableTarget)
			if err != nil {
				return nil, err
			}
			chromeTargets = append(chromeTargets, target)
		}
	}
	return chromeTargets, nil
}

func (c *Gcd) getConnectableTargets() ([]*TargetInfo, error) {
	// some times it takes a while to get results, so retry 4x
	for i := 0; i < 4; i++ {
		resp, err := http.Get(c.apiEndpoint)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, errRead := ioutil.ReadAll(resp.Body)
		if errRead != nil {
			return nil, &GcdBodyReadErr{Message: errRead.Error()}
		}

		targets := make([]*TargetInfo, 0)
		err = json.Unmarshal(body, &targets)
		if err != nil {
			return nil, &GcdDecodingErr{Message: err.Error()}
		}

		connectableTargets := make([]*TargetInfo, 0)
		for _, v := range targets {
			if v.WebSocketDebuggerUrl != "" {
				connectableTargets = append(connectableTargets, v)
			}
		}

		if len(connectableTargets) > 0 {
			return connectableTargets, nil
		}
		time.Sleep(time.Millisecond * 350)
	}
	return nil, ErrNoTabAvailable
}

// NewTab a new empty tab, returns the chrome target.
func (c *Gcd) NewTab() (*ChromeTarget, error) {
	resp, err := http.Get(c.apiEndpoint + "/new")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return nil, &GcdBodyReadErr{Message: errRead.Error()}
	}

	tabTarget := &TargetInfo{}
	err = json.Unmarshal(body, &tabTarget)
	if err != nil {
		return nil, &GcdDecodingErr{Message: err.Error()}
	}
	return openChromeTarget(c, tabTarget)
}

// GetFirstTab returns the first tab created, to be called when
// first started, otherwise you will get a random tab returned.
func (c *Gcd) GetFirstTab() (*ChromeTarget, error) {
	connectableTargets, err := c.getConnectableTargets()
	if err != nil {
		return nil, err
	}

	for _, tabTarget := range connectableTargets {
		if tabTarget.Type == "page" {
			return openChromeTarget(c, tabTarget)
		}
	}
	return nil, ErrNoTabAvailable
}

// GetRevision of chrome
func (c *Gcd) GetRevision() string {
	return gcdapi.CHROME_VERSION
}

// CloseTab closes the target tab.
func (c *Gcd) CloseTab(target *ChromeTarget) error {
	resp, err := http.Get(fmt.Sprintf("%s/close/%s", c.apiEndpoint, target.Target.Id))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, errRead := ioutil.ReadAll(resp.Body)
	return errRead
}

// ActivateTab (focus) the tab.
func (c *Gcd) ActivateTab(target *ChromeTarget) error {
	resp, err := http.Get(fmt.Sprintf("%s/activate/%s", c.apiEndpoint, target.Target.Id))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, errRead := ioutil.ReadAll(resp.Body)
	return errRead
}

// probes the debugger report and signals when it's available.
func (c *Gcd) probeDebugPort() error {
	ticker := time.NewTicker(time.Millisecond * 100)
	timeoutTicker := time.NewTicker(time.Second * c.timeout)

	defer func() {
		ticker.Stop()
		timeoutTicker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			resp, err := http.Get(c.apiEndpoint)
			if err != nil {
				continue
			}
			defer resp.Body.Close()
			c.readyCh <- struct{}{}
			return nil
		case <-timeoutTicker.C:
			return fmt.Errorf("Unable to contact debugger at %s after %d seconds, gave up", c.apiEndpoint, c.timeout)
		}
	}
}
