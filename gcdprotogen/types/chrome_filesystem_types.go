// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the FileSystem types.
// API Version: 1.1
package types

// Represents a browser side file or directory.
type ChromeFileSystemEntry struct {
	Url          string                  `json:"url"`                    // filesystem: URL for the entry.
	Name         string                  `json:"name"`                   // The name of the file or directory.
	IsDirectory  bool                    `json:"isDirectory"`            // True if the entry is a directory.
	MimeType     string                  `json:"mimeType,omitempty"`     // MIME type of the entry, available for a file only.
	ResourceType *ChromePageResourceType `json:"resourceType,omitempty"` // ResourceType of the entry, available for a file only.
	IsTextFile   bool                    `json:"isTextFile,omitempty"`   // True if the entry is a text file.
}

// Represents metadata of a file or entry.
type ChromeFileSystemMetadata struct {
	ModificationTime float64 `json:"modificationTime"` // Modification time.
	Size             float64 `json:"size"`             // File size. This field is always zero for directories.
}
