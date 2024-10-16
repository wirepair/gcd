// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains FileSystem functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// No Description.
type FileSystemFile struct {
	Name         string  `json:"name"`         //
	LastModified float64 `json:"lastModified"` // Timestamp
	Size         float64 `json:"size"`         // Size in bytes
	Type         string  `json:"type"`         //
}

// No Description.
type FileSystemDirectory struct {
	Name              string            `json:"name"`              //
	NestedDirectories []string          `json:"nestedDirectories"` //
	NestedFiles       []*FileSystemFile `json:"nestedFiles"`       // Files that are directly nested under this directory.
}

// No Description.
type FileSystemBucketFileSystemLocator struct {
	StorageKey     string   `json:"storageKey"`           // Storage key
	BucketName     string   `json:"bucketName,omitempty"` // Bucket name. Not passing a `bucketName` will retrieve the default Bucket. (https://developer.mozilla.org/en-US/docs/Web/API/Storage_API#storage_buckets)
	PathComponents []string `json:"pathComponents"`       // Path to the directory using each path component as an array item.
}

type FileSystem struct {
	target gcdmessage.ChromeTargeter
}

func NewFileSystem(target gcdmessage.ChromeTargeter) *FileSystem {
	c := &FileSystem{target: target}
	return c
}

type FileSystemGetDirectoryParams struct {
	//
	BucketFileSystemLocator *FileSystemBucketFileSystemLocator `json:"bucketFileSystemLocator"`
}

// GetDirectoryWithParams -
// Returns -  directory - Returns the directory object at the path.
func (c *FileSystem) GetDirectoryWithParams(ctx context.Context, v *FileSystemGetDirectoryParams) (*FileSystemDirectory, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FileSystem.getDirectory", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Directory *FileSystemDirectory
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Directory, nil
}

// GetDirectory -
// bucketFileSystemLocator -
// Returns -  directory - Returns the directory object at the path.
func (c *FileSystem) GetDirectory(ctx context.Context, bucketFileSystemLocator *FileSystemBucketFileSystemLocator) (*FileSystemDirectory, error) {
	var v FileSystemGetDirectoryParams
	v.BucketFileSystemLocator = bucketFileSystemLocator
	return c.GetDirectoryWithParams(ctx, &v)
}
