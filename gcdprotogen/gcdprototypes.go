package main

// Top level API protocol file
type DebuggerApi struct {
	Version *ApiVersion `json:"version"`
	Domains []*Domain   `json:"domains"`
}

// Version information
type ApiVersion struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
}

// The Domain (contains all objects, their type/commands/events)
type Domain struct {
	Domain      string     `json:"domain"`
	Description string     `json:"description,omitempty"`
	Types       []*Type    `json:"types,omitempty"`
	Commands    []*Command `json:"commands,omitempty"`
	Events      []*Event   `json:"events,omitempty"`
	Hidden      bool       `json:"hidden,omitempty"`
	Items       *Item      `json:"items,omitempty"`
}

// A Type which represents objects specific to the API method
type Type struct {
	Id          string      `json:"id"`
	Type        string      `json:"type"`
	Description string      `json:"description,omitempty"`
	Enum        []string    `json:"enum,omitempty"`
	Properties  []*Property `json:"properties,omitempty"`
	Hidden      bool        `json:"hidden,omitempty"`
	Items       *Item       `json:"items,omitempty"`
	MinItems    int64       `json:"minItems,omitempty"`
	MaxItems    int64       `json:"maxItems,omitempty"`
}

// A property & Parameter type used by both commands & types
type Property struct {
	Name        string   `json:"name"`
	Type        string   `json:"type,omitempty"`
	Description string   `json:"description,omitempty"`
	Ref         string   `json:"$ref,omitempty"`
	Optional    bool     `json:"optional,omitempty"`
	Hidden      bool     `json:"hidden,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	Items       *Item    `json:"items,omitempty"`
}

// An item used by types, properties and events.
type Item struct {
	Type        string      `json:"type,omitempty"`
	Ref         string      `json:"$ref,omitempty"`
	Properties  []*Property `json:"properties,omitempty"`
	Description string      `json:"description,omitempty"`
	Enum        []string    `json:"enum,omitempty"`
}

// The API Command call.
type Command struct {
	Name        string            `json:"name"`
	Type        string            `json:"type,omitempty"`
	Description string            `json:"description,omitempty"`
	Handlers    []string          `json:"handlers,omitempty"`
	Parameters  []*Property       `json:"parameters,omitempty"`
	Returns     []*CommandReturns `json:"returns,omitempty"`
	Hidden      bool              `json:"hidden,omitempty"`
	Async       bool              `json:"async,omitempty"`
	Redirect    string            `json:"redirect,omitempty"`
}

// Parameters to the API Command call
type CommandParameters struct {
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
	Ref  string `json:"$ref,omitempty"`
}

// The return parameters for an API call
type CommandReturns struct {
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
	Ref         string `json:"$ref,omitempty"`
	Description string `json:"description,omitempty"`
}

// An event, asynchornous events that can come in once
// enabled.
type Event struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Parameters  []*Property `json:"parameters,omitempty"`
	Deprecated  bool        `json:"deprecated,omitempty"`
	Hidden      bool        `json:"hidden,omitempty"`
}
