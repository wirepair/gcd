// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains FileSystem functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Represents a browser side file or directory.
type FileSystemEntry struct {
	Url          string `json:"url"`                    // filesystem: URL for the entry.
	Name         string `json:"name"`                   // The name of the file or directory.
	IsDirectory  bool   `json:"isDirectory"`            // True if the entry is a directory.
	MimeType     string `json:"mimeType,omitempty"`     // MIME type of the entry, available for a file only.
	ResourceType string `json:"resourceType,omitempty"` // ResourceType of the entry, available for a file only. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Other
	IsTextFile   bool   `json:"isTextFile,omitempty"`   // True if the entry is a text file.
}

// Represents metadata of a file or entry.
type FileSystemMetadata struct {
	ModificationTime float64 `json:"modificationTime"` // Modification time.
	Size             float64 `json:"size"`             // File size. This field is always zero for directories.
}

type FileSystem struct {
	target gcdmessage.ChromeTargeter
}

func NewFileSystem(target gcdmessage.ChromeTargeter) *FileSystem {
	c := &FileSystem{target: target}
	return c
}

// Enables events from backend.
func (c *FileSystem) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FileSystem.enable"})
}

// Disables events from backend.
func (c *FileSystem) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FileSystem.disable"})
}

// RequestFileSystemRoot - Returns root directory of the FileSystem, if exists.
// origin - Security origin of requesting FileSystem. One of frames in current page needs to have this security origin.
// type - FileSystem type of requesting FileSystem.
// Returns -  errorCode - 0, if no error. Otherwise, errorCode is set to FileError::ErrorCode value. root - Contains root of the requested FileSystem if the command completed successfully.
func (c *FileSystem) RequestFileSystemRoot(origin string, theType string) (int, *FileSystemEntry, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["origin"] = origin
	paramRequest["type"] = theType
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FileSystem.requestFileSystemRoot"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ErrorCode int
			Root      *FileSystemEntry
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, nil, err
	}

	return chromeData.Result.ErrorCode, chromeData.Result.Root, nil
}

// RequestDirectoryContent - Returns content of the directory.
// url - URL of the directory that the frontend is requesting to read from.
// Returns -  errorCode - 0, if no error. Otherwise, errorCode is set to FileError::ErrorCode value. entries - Contains all entries on directory if the command completed successfully.
func (c *FileSystem) RequestDirectoryContent(url string) (int, []*FileSystemEntry, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FileSystem.requestDirectoryContent"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ErrorCode int
			Entries   []*FileSystemEntry
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, nil, err
	}

	return chromeData.Result.ErrorCode, chromeData.Result.Entries, nil
}

// RequestMetadata - Returns metadata of the entry.
// url - URL of the entry that the frontend is requesting to get metadata from.
// Returns -  errorCode - 0, if no error. Otherwise, errorCode is set to FileError::ErrorCode value. metadata - Contains metadata of the entry if the command completed successfully.
func (c *FileSystem) RequestMetadata(url string) (int, *FileSystemMetadata, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FileSystem.requestMetadata"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ErrorCode int
			Metadata  *FileSystemMetadata
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, nil, err
	}

	return chromeData.Result.ErrorCode, chromeData.Result.Metadata, nil
}

// RequestFileContent - Returns content of the file. Result should be sliced into [start, end).
// url - URL of the file that the frontend is requesting to read from.
// readAsText - True if the content should be read as text, otherwise the result will be returned as base64 encoded text.
// start - Specifies the start of range to read.
// end - Specifies the end of range to read exclusively.
// charset - Overrides charset of the content when content is served as text.
// Returns -  errorCode - 0, if no error. Otherwise, errorCode is set to FileError::ErrorCode value. content - Content of the file. charset - Charset of the content if it is served as text.
func (c *FileSystem) RequestFileContent(url string, readAsText bool, start int, end int, charset string) (int, string, string, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["url"] = url
	paramRequest["readAsText"] = readAsText
	paramRequest["start"] = start
	paramRequest["end"] = end
	paramRequest["charset"] = charset
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FileSystem.requestFileContent"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ErrorCode int
			Content   string
			Charset   string
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, "", "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, "", "", err
	}

	return chromeData.Result.ErrorCode, chromeData.Result.Content, chromeData.Result.Charset, nil
}

// DeleteEntry - Deletes specified entry. If the entry is a directory, the agent deletes children recursively.
// url - URL of the entry to delete.
// Returns -  errorCode - 0, if no error. Otherwise errorCode is set to FileError::ErrorCode value.
func (c *FileSystem) DeleteEntry(url string) (int, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FileSystem.deleteEntry"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ErrorCode int
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, err
	}

	return chromeData.Result.ErrorCode, nil
}
