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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	revisionOS          = "win" // doesn't really matter
	revisionEndpoint    = "https://omahaproxy.appspot.com/all.json"
	browserProtocolFile = "https://chromium.googlesource.com/chromium/src/+/%s/third_party/WebKit/Source/core/inspector/browser_protocol.json?format=text"
	jsProtocolFile      = "https://chromium.googlesource.com/v8/v8/+/%s/src/inspector/js_protocol.json?format=text"
)

type ChromiumRevision []struct {
	Os       string `json:"os"`
	Versions []struct {
		BranchCommit       string `json:"branch_commit"`
		BranchBasePosition string `json:"branch_base_position"`
		SkiaCommit         string `json:"skia_commit"`
		V8Version          string `json:"v8_version"`
		PreviousVersion    string `json:"previous_version"`
		V8Commit           string `json:"v8_commit"`
		TrueBranch         string `json:"true_branch"`
		PreviousReldate    string `json:"previous_reldate"`
		BranchBaseCommit   string `json:"branch_base_commit"`
		Version            string `json:"version"`
		CurrentReldate     string `json:"current_reldate"`
		CurrentVersion     string `json:"current_version"`
		Os                 string `json:"os"`
		Channel            string `json:"channel"`
		ChromiumCommit     string `json:"chromium_commit"`
	} `json:"versions"`
}

type RevisionInfo struct {
	Channel  string
	Version  string
	Branch   string
	JsBranch string
}

func getApiRevision(channel string) *RevisionInfo {
	versionInfo := getRevisionData(channel)
	if versionInfo == nil {
		log.Fatalf("Error finding version information from %s", revisionEndpoint)
	}
	download(fmt.Sprintf(browserProtocolFile, versionInfo.Branch), fmt.Sprintf(jsProtocolFile, versionInfo.JsBranch))
	return versionInfo
}

func getRevisionData(channel string) *RevisionInfo {
	var revision ChromiumRevision
	revisionData := getRemoteFile(revisionEndpoint)

	if err := json.Unmarshal(revisionData, &revision); err != nil {
		log.Fatalf("Error getting revision information: %s\n", err)
	}

	for _, revisions := range revision {
		if revisions.Os != revisionOS {
			continue
		}

		for _, versions := range revisions.Versions {
			if versions.Channel != channel {
				continue
			}
			return &RevisionInfo{Channel: versions.Channel, Branch: versions.BranchCommit, JsBranch: versions.V8Commit, Version: versions.Version}
		}
	}
	return nil
}

func download(browserFile, jsFile string) {
	browserData := fixBrokenBool(decodeProtocol(getRemoteFile(browserFile)))
	jsData := fixBrokenBool(decodeProtocol(getRemoteFile(jsFile)))

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

// because google's git viewer is broken, downloading ?format=json, we have to
// base64 decode the text response.
func decodeProtocol(data []byte) []byte {
	decoded, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		log.Fatalf("error base64 decoding data: %s %s\n", err, string(data))
	}
	return decoded
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
	f.Write(protocolData)
	f.Sync()
	f.Close()
}
