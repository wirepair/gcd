[![Go Report Card](https://goreportcard.com/badge/github.com/wirepair/gcd)](https://goreportcard.com/report/github.com/wirepair/gcd)

# Google Chrome Debugger (GCD)

This is primarly an auto-generated client library for communicating with a Google Chrome Browser over their [remote client debugger protocol](https://developer.chrome.com/devtools/docs/debugger-protocol). Note that their documentation is partially incorrect and does not contain a lot of the API calls that are actually available.

Because I'm lazy and there are hundereds of different custom types and API methods, this library has been automatically generated using their [protocol.json](https://code.google.com/p/chromium/codesearch#chromium/src/third_party/WebKit/Source/devtools/protocol.json&q=protocol.json&sq=package:chromium&type=cs).

The [gcdapigen](https://github.com/wirepair/gcd/tree/master/v2/gcdapigen) program was created to generate types, event types and commands for gcd.

# Changelog

See the [CHANGELOG](https://github.com/wirepair/gcd/tree/master/CHANGELOG.md).

## Dependencies

gcd requires the [gcdapi](https://github.com/wirepair/gcd/tree/master/v2/gcdapi) and [gcdmessage](https://github.com/wirepair/gcd/tree/master/v2/gcdmessage) packages. gcdapi is the auto-generated API. gcdmessage is the glue between gcd and gcdapi so we can keep the packages clean.

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

For a full list of api methods, types, event types & godocs: [Documentation](https://godoc.org/github.com/wirepair/gcd/v2/gcdapi)

## Examples

See [examples](https://github.com/wirepair/gcd/tree/master/v2/examples)

## Licensing

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
