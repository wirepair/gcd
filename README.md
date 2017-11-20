[![Go Report Card](https://goreportcard.com/badge/github.com/wirepair/gcd)](https://goreportcard.com/report/github.com/wirepair/gcd)

# Google Chrome Debugger (GCD)
This is primarly an auto-generated client library for communicating with a Google Chrome Browser over their [remote client debugger protocol](https://developer.chrome.com/devtools/docs/debugger-protocol). Note that their documentation is partially incorrect and does not contain a lot of the API calls that are actually available.

Because I'm lazy and there are hundereds of different custom types and API methods, this library has been automatically generated using their [protocol.json](https://code.google.com/p/chromium/codesearch#chromium/src/third_party/WebKit/Source/devtools/protocol.json&q=protocol.json&sq=package:chromium&type=cs). 

The [gcdapigen](https://github.com/wirepair/gcd/tree/master/gcdapigen) program was created to generate types, event types and commands for gcd.

# Changelog 
See the [CHANGELOG](https://github.com/wirepair/gcd/tree/master/CHANGELOG.md).

## Dependencies
gcd requires the [gcdapi](https://github.com/wirepair/gcd/tree/master/gcdapi) and [gcdmessage](https://github.com/wirepair/gcd/tree/master/gcdmessage) packages. gcdapi is the auto-generated API. gcdmessage is the glue between gcd and gcdapi so we can keep the packages clean. 

## The API
The API consists of of synchronous requests, asynchronous requests / events. Synchronous requests are handled by using non-buffered channels and methods can be called and will return once the value is available. Events are handled by subscribing the response method type and calling the API's "Enable()" such as:
```Go
    target, err := debugger.NewTab()
    if err != nil {
    	log.Fatalf("error getting new tab: %s\n", err)
	}
	console := target.Console

	target.Subscribe("Console.messageAdded", func(target *ChromeTarget, v []byte) {
		
		msg := &gcdapi.ConsoleMessageAddedEvent{}
		err := json.Unmarshal(v, msg)
		if err != nil {
			log.Fatalf("error unmarshalling event data: %v\n", err)
		}
		log.Printf("METHOD: %s\n", msg.Method)
		eventData := msg.Params.Message
		log.Printf("Got event: %s\n", eventData)
	})
	console.Enable()
	// recv events
	// console.Disable()
```

## Usage
For a full list of api methods, types, event types & godocs: [Documentation](https://godoc.org/github.com/wirepair/gcd/gcdapi)

Loading a page using the Page API.
```Go
package main

import (
	"github.com/wirepair/gcd"
	"github.com/wirepair/gcd/gcdapi"
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
	target, err := debugger.NewTab()
	if err != nil {
		log.Fatalf("error getting targets: %s\n", err)
	}

	//subscribe to page load
	target.Subscribe("Page.loadEventFired", func(targ *gcd.ChromeTarget, v []byte) {
		wg.Done() // page loaded, we can exit now
		// if you wanted to inspect the full response data, you could do that here
	})

	// get the Page API and enable it
	if _, err := target.Page.Enable(); err != nil {
		log.Fatalf("error getting page: %s\n", err)
	}

	navigateParams := &gcdapi.PageNavigateParams{Url: "http://www.veracode.com"}
	_, _, err := target.Page.NavigateWithParams(navigateParams)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	log.Printf("ret: %#v\n", ret)
	wg.Wait() // wait for page load
}

```
Getting a document from a target.
```Go
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
	r, err := dom.GetDocument(-1, true)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	log.Printf("NodeID: %d Node Name: %s, URL: %s\n", r.NodeId, r.NodeName, r.DocumentURL)
}
```

Create 5 tabs, navigate to URLs, wait for them to load, take screen shot
```Go
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/wirepair/gcd"
	"github.com/wirepair/gcd/gcdapi"
	"log"
	"net/url"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	numTabs = 5
)

var debugger *gcd.Gcd
var wg sync.WaitGroup

var path string
var dir string
var port string

func init() {
	switch runtime.GOOS {
	case "windows":
		flag.StringVar(&path, "chrome", "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "path to chrome")
		flag.StringVar(&dir, "dir", "C:\\temp\\", "user directory")
	case "darwin":
		flag.StringVar(&path, "chrome", "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", "path to chrome")
		flag.StringVar(&dir, "dir", "/tmp/", "user directory")
	case "linux":
		flag.StringVar(&path, "chrome", "/usr/bin/chromium-browser", "path to chrome")
		flag.StringVar(&dir, "dir", "/tmp/", "user directory")
	}

	flag.StringVar(&port, "port", "9222", "Debugger port")
}

func main() {
	var err error
	urls := []string{"http://www.google.com", "http://www.veracode.com", "http://www.microsoft.com", "http://bbc.co.uk", "http://www.reddit.com/r/golang"}

	flag.Parse()

	debugger = gcd.NewChromeDebugger()
	debugger.StartProcess(path, dir, port)
	defer debugger.ExitProcess()
	targets := make([]*gcd.ChromeTarget, numTabs)

	for i := 0; i < numTabs; i++ {
		wg.Add(1)
		targets[i], err = debugger.NewTab()
		if err != nil {
			log.Fatalf("error getting targets")
		}
		page := targets[i].Page
		page.Enable()
		targets[i].Subscribe("Page.loadEventFired", PageLoaded)
		// navigate
		navigateParams := &gcdapi.PageNavigateParams{Url: urls[i]}
		_, _, err := page.NavigateWithParams(navigateParams)
		if err != nil {
			log.Fatalf("error: %s\n", err)
		}
	}
	wg.Wait()
	for i := 0; i < numTabs; i++ {
		TakeScreenShot(targets[i])
	}
}

func PageLoaded(target *gcd.ChromeTarget, event []byte) {
	wg.Done()
}

func TakeScreenShot(target *gcd.ChromeTarget) {
	dom := target.DOM
	page := target.Page
	doc, err := dom.GetDocument(-1, true)
	if err != nil {
		fmt.Printf("error getting doc: %s\n", err)
		return
	}

	debugger.ActivateTab(target)
	time.Sleep(1 * time.Second) // give it a sec to paint
	u, urlErr := url.Parse(doc.DocumentURL)
	if urlErr != nil {
		fmt.Printf("error parsing url: %s\n", urlErr)
		return
	}

	fmt.Printf("Taking screen shot of: %s\n", u.Host)
	screenShotParams := &gcdapi.PageCaptureScreenshotParams{Format: "png", FromSurface: true}
	img, errCap := page.CaptureScreenshotWithParams(screenShotParams)
	if errCap != nil {
		fmt.Printf("error getting doc: %s\n", errCap)
		return
	}

	imgBytes, errDecode := base64.StdEncoding.DecodeString(img)
	if errDecode != nil {
		fmt.Printf("error decoding image: %s\n", errDecode)
		return
	}

	f, errFile := os.Create(u.Host + ".png")
	defer f.Close()
	if errFile != nil {
		fmt.Printf("error creating image file: %s\n", errFile)
		return
	}
	f.Write(imgBytes)
}

```

## Licensing
The MIT License (MIT)

Copyright (c) 2017 isaac dawson

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
