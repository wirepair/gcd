/*
The MIT License (MIT)

Copyright (c) 2016 isaac dawson

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

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	browserProtocolFile = "https://raw.githubusercontent.com/ChromeDevTools/devtools-protocol/refs/heads/master/json/browser_protocol.json"
	jsProtocolFile      = "https://raw.githubusercontent.com/ChromeDevTools/devtools-protocol/refs/heads/master/json/js_protocol.json"
)

func download(browserFile, jsFile string) {
	browserData := fixBrokenBool(getRemoteFile(browserFile))
	jsData := fixBrokenBool(getRemoteFile(jsFile))

	api := &ProtoDebuggerApi{}
	copyApi := &ProtoDebuggerApi{}

	if err := json.Unmarshal(browserData, api); err != nil {
		log.Fatalf("error unmarshalling browser data: %s\n", err)
	}
	log.Printf("len: %d\n", len(api.Domains))

	// add js data to our api
	if err := json.Unmarshal(jsData, copyApi); err != nil {
		log.Fatalf("error unmarshalling js data: %s\n", err)
	}

	// append Domain entries
	api.Domains = append(api.Domains, copyApi.Domains...)
	log.Printf("len: %d\n", len(api.Domains))
	merged, err := json.Marshal(api)
	if err != nil {
		log.Fatalf("error marshaling merged protocol data: %s\n", err)
	}
	writeFile(merged)
}

func getRemoteFile(fileName string) []byte {
	var data []byte

	log.Printf("requesting: %s\n", fileName)
	resp, err := http.Get(fileName)
	if err != nil {
		log.Fatalf("error requesting %s\n", fileName)
	}
	defer resp.Body.Close()

	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Fatalf("error reading data from response: %s\n", err)
	}

	return data
}

func fixBrokenBool(data []byte) []byte {
	return bytes.Replace(data, []byte("\"true\""), []byte("true"), -1)
}

func writeFile(protocolData []byte) {
	var f *os.File
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("error opening file %s for writing: %s\n", file, err)
	}
	protocolData = bytes.Replace(protocolData, []byte("\\n"), []byte(" "), -1) // remove newlines
	f.Write(protocolData)
	f.Sync()
	f.Close()
}
