// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Runtime functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Mirror object referencing original JavaScript object.
type RuntimeRemoteObject struct {
	Type                string                `json:"type"`                          // Object type.
	Subtype             string                `json:"subtype,omitempty"`             // Object subtype hint. Specified for `object` type values only. NOTE: If you change anything here, make sure to also update `subtype` in `ObjectPreview` and `PropertyPreview` below.
	ClassName           string                `json:"className,omitempty"`           // Object class (constructor) name. Specified for `object` type values only.
	Value               interface{}           `json:"value,omitempty"`               // Remote object value in case of primitive values or JSON values (if it was requested).
	UnserializableValue string                `json:"unserializableValue,omitempty"` // Primitive value which can not be JSON-stringified does not have `value`, but gets this property.
	Description         string                `json:"description,omitempty"`         // String representation of the object.
	ObjectId            string                `json:"objectId,omitempty"`            // Unique object identifier (for non-primitive values).
	Preview             *RuntimeObjectPreview `json:"preview,omitempty"`             // Preview containing abbreviated property values. Specified for `object` type values only.
	CustomPreview       *RuntimeCustomPreview `json:"customPreview,omitempty"`       //
}

// No Description.
type RuntimeCustomPreview struct {
	Header       string `json:"header"`                 // The JSON-stringified result of formatter.header(object, config) call. It contains json ML array that represents RemoteObject.
	BodyGetterId string `json:"bodyGetterId,omitempty"` // If formatter returns true as a result of formatter.hasBody call then bodyGetterId will contain RemoteObjectId for the function that returns result of formatter.body(object, config) call. The result value is json ML array.
}

// Object containing abbreviated remote object value.
type RuntimeObjectPreview struct {
	Type        string                    `json:"type"`                  // Object type.
	Subtype     string                    `json:"subtype,omitempty"`     // Object subtype hint. Specified for `object` type values only.
	Description string                    `json:"description,omitempty"` // String representation of the object.
	Overflow    bool                      `json:"overflow"`              // True iff some of the properties or entries of the original object did not fit.
	Properties  []*RuntimePropertyPreview `json:"properties"`            // List of the properties.
	Entries     []*RuntimeEntryPreview    `json:"entries,omitempty"`     // List of the entries. Specified for `map` and `set` subtype values only.
}

// No Description.
type RuntimePropertyPreview struct {
	Name         string                `json:"name"`                   // Property name.
	Type         string                `json:"type"`                   // Object type. Accessor means that the property itself is an accessor property.
	Value        string                `json:"value,omitempty"`        // User-friendly property value string.
	ValuePreview *RuntimeObjectPreview `json:"valuePreview,omitempty"` // Nested value preview.
	Subtype      string                `json:"subtype,omitempty"`      // Object subtype hint. Specified for `object` type values only.
}

// No Description.
type RuntimeEntryPreview struct {
	Key   *RuntimeObjectPreview `json:"key,omitempty"` // Preview of the key. Specified for map-like collection entries.
	Value *RuntimeObjectPreview `json:"value"`         // Preview of the value.
}

// Object property descriptor.
type RuntimePropertyDescriptor struct {
	Name         string               `json:"name"`                // Property name or symbol description.
	Value        *RuntimeRemoteObject `json:"value,omitempty"`     // The value associated with the property.
	Writable     bool                 `json:"writable,omitempty"`  // True if the value associated with the property may be changed (data descriptors only).
	Get          *RuntimeRemoteObject `json:"get,omitempty"`       // A function which serves as a getter for the property, or `undefined` if there is no getter (accessor descriptors only).
	Set          *RuntimeRemoteObject `json:"set,omitempty"`       // A function which serves as a setter for the property, or `undefined` if there is no setter (accessor descriptors only).
	Configurable bool                 `json:"configurable"`        // True if the type of this property descriptor may be changed and if the property may be deleted from the corresponding object.
	Enumerable   bool                 `json:"enumerable"`          // True if this property shows up during enumeration of the properties on the corresponding object.
	WasThrown    bool                 `json:"wasThrown,omitempty"` // True if the result was thrown during the evaluation.
	IsOwn        bool                 `json:"isOwn,omitempty"`     // True if the property is owned for the object.
	Symbol       *RuntimeRemoteObject `json:"symbol,omitempty"`    // Property symbol object, if the property is of the `symbol` type.
}

// Object internal property descriptor. This property isn't normally visible in JavaScript code.
type RuntimeInternalPropertyDescriptor struct {
	Name  string               `json:"name"`            // Conventional property name.
	Value *RuntimeRemoteObject `json:"value,omitempty"` // The value associated with the property.
}

// Object private field descriptor.
type RuntimePrivatePropertyDescriptor struct {
	Name  string               `json:"name"`            // Private property name.
	Value *RuntimeRemoteObject `json:"value,omitempty"` // The value associated with the private property.
	Get   *RuntimeRemoteObject `json:"get,omitempty"`   // A function which serves as a getter for the private property, or `undefined` if there is no getter (accessor descriptors only).
	Set   *RuntimeRemoteObject `json:"set,omitempty"`   // A function which serves as a setter for the private property, or `undefined` if there is no setter (accessor descriptors only).
}

// Represents function call argument. Either remote object id `objectId`, primitive `value`, unserializable primitive value or neither of (for undefined) them should be specified.
type RuntimeCallArgument struct {
	Value               interface{} `json:"value,omitempty"`               // Primitive value or serializable javascript object.
	UnserializableValue string      `json:"unserializableValue,omitempty"` // Primitive value which can not be JSON-stringified.
	ObjectId            string      `json:"objectId,omitempty"`            // Remote object handle.
}

// Description of an isolated world.
type RuntimeExecutionContextDescription struct {
	Id       int                    `json:"id"`                // Unique id of the execution context. It can be used to specify in which execution context script evaluation should be performed.
	Origin   string                 `json:"origin"`            // Execution context origin.
	Name     string                 `json:"name"`              // Human readable name describing given context.
	UniqueId string                 `json:"uniqueId"`          // A system-unique execution context identifier. Unlike the id, this is unique across multiple processes, so can be reliably used to identify specific context while backend performs a cross-process navigation.
	AuxData  map[string]interface{} `json:"auxData,omitempty"` // Embedder-specific auxiliary data.
}

// Detailed information about exception (or error) that was thrown during script compilation or execution.
type RuntimeExceptionDetails struct {
	ExceptionId        int                    `json:"exceptionId"`                  // Exception id.
	Text               string                 `json:"text"`                         // Exception text, which should be used together with exception object when available.
	LineNumber         int                    `json:"lineNumber"`                   // Line number of the exception location (0-based).
	ColumnNumber       int                    `json:"columnNumber"`                 // Column number of the exception location (0-based).
	ScriptId           string                 `json:"scriptId,omitempty"`           // Script ID of the exception location.
	Url                string                 `json:"url,omitempty"`                // URL of the exception location, to be used when the script was not reported.
	StackTrace         *RuntimeStackTrace     `json:"stackTrace,omitempty"`         // JavaScript stack trace if available.
	Exception          *RuntimeRemoteObject   `json:"exception,omitempty"`          // Exception object if available.
	ExecutionContextId int                    `json:"executionContextId,omitempty"` // Identifier of the context where exception happened.
	ExceptionMetaData  map[string]interface{} `json:"exceptionMetaData,omitempty"`  // Dictionary with entries of meta data that the client associated with this exception, such as information about associated network requests, etc.
}

// Stack entry for runtime errors and assertions.
type RuntimeCallFrame struct {
	FunctionName string `json:"functionName"` // JavaScript function name.
	ScriptId     string `json:"scriptId"`     // JavaScript script id.
	Url          string `json:"url"`          // JavaScript script name or url.
	LineNumber   int    `json:"lineNumber"`   // JavaScript script line number (0-based).
	ColumnNumber int    `json:"columnNumber"` // JavaScript script column number (0-based).
}

// Call frames for assertions or error messages.
type RuntimeStackTrace struct {
	Description string               `json:"description,omitempty"` // String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
	CallFrames  []*RuntimeCallFrame  `json:"callFrames"`            // JavaScript function name.
	Parent      *RuntimeStackTrace   `json:"parent,omitempty"`      // Asynchronous JavaScript stack trace that preceded this stack, if available.
	ParentId    *RuntimeStackTraceId `json:"parentId,omitempty"`    // Asynchronous JavaScript stack trace that preceded this stack, if available.
}

// If `debuggerId` is set stack trace comes from another debugger and can be resolved there. This allows to track cross-debugger calls. See `Runtime.StackTrace` and `Debugger.paused` for usages.
type RuntimeStackTraceId struct {
	Id         string `json:"id"`                   //
	DebuggerId string `json:"debuggerId,omitempty"` //
}

// Notification is issued every time when binding is called.
type RuntimeBindingCalledEvent struct {
	Method string `json:"method"`
	Params struct {
		Name               string `json:"name"`               //
		Payload            string `json:"payload"`            //
		ExecutionContextId int    `json:"executionContextId"` // Identifier of the context where the call was made.
	} `json:"Params,omitempty"`
}

// Issued when console API was called.
type RuntimeConsoleAPICalledEvent struct {
	Method string `json:"method"`
	Params struct {
		Type               string                 `json:"type"`                 // Type of the call.
		Args               []*RuntimeRemoteObject `json:"args"`                 // Call arguments.
		ExecutionContextId int                    `json:"executionContextId"`   // Identifier of the context where the call was made.
		Timestamp          float64                `json:"timestamp"`            // Call timestamp.
		StackTrace         *RuntimeStackTrace     `json:"stackTrace,omitempty"` // Stack trace captured when the call was made. The async stack chain is automatically reported for the following call types: `assert`, `error`, `trace`, `warning`. For other types the async call chain can be retrieved using `Debugger.getStackTrace` and `stackTrace.parentId` field.
		Context            string                 `json:"context,omitempty"`    // Console context descriptor for calls on non-default console context (not console.*): 'anonymous#unique-logger-id' for call on unnamed context, 'name#unique-logger-id' for call on named context.
	} `json:"Params,omitempty"`
}

// Issued when unhandled exception was revoked.
type RuntimeExceptionRevokedEvent struct {
	Method string `json:"method"`
	Params struct {
		Reason      string `json:"reason"`      // Reason describing why exception was revoked.
		ExceptionId int    `json:"exceptionId"` // The id of revoked exception, as reported in `exceptionThrown`.
	} `json:"Params,omitempty"`
}

// Issued when exception was thrown and unhandled.
type RuntimeExceptionThrownEvent struct {
	Method string `json:"method"`
	Params struct {
		Timestamp        float64                  `json:"timestamp"`        // Timestamp of the exception.
		ExceptionDetails *RuntimeExceptionDetails `json:"exceptionDetails"` //
	} `json:"Params,omitempty"`
}

// Issued when new execution context is created.
type RuntimeExecutionContextCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Context *RuntimeExecutionContextDescription `json:"context"` // A newly created execution context.
	} `json:"Params,omitempty"`
}

// Issued when execution context is destroyed.
type RuntimeExecutionContextDestroyedEvent struct {
	Method string `json:"method"`
	Params struct {
		ExecutionContextId int `json:"executionContextId"` // Id of the destroyed context
	} `json:"Params,omitempty"`
}

// Issued when object should be inspected (for example, as a result of inspect() command line API call).
type RuntimeInspectRequestedEvent struct {
	Method string `json:"method"`
	Params struct {
		Object             *RuntimeRemoteObject   `json:"object"`                       //
		Hints              map[string]interface{} `json:"hints"`                        //
		ExecutionContextId int                    `json:"executionContextId,omitempty"` // Identifier of the context where the call was made.
	} `json:"Params,omitempty"`
}

type Runtime struct {
	target gcdmessage.ChromeTargeter
}

func NewRuntime(target gcdmessage.ChromeTargeter) *Runtime {
	c := &Runtime{target: target}
	return c
}

type RuntimeAwaitPromiseParams struct {
	// Identifier of the promise.
	PromiseObjectId string `json:"promiseObjectId"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
}

// AwaitPromiseWithParams - Add handler to promise with given promise object id.
// Returns -  result - Promise result. Will contain rejected value if promise was rejected. exceptionDetails - Exception details if stack strace is available.
func (c *Runtime) AwaitPromiseWithParams(ctx context.Context, v *RuntimeAwaitPromiseParams) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.awaitPromise", Params: v})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			Result           *RuntimeRemoteObject
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.ExceptionDetails, nil
}

// AwaitPromise - Add handler to promise with given promise object id.
// promiseObjectId - Identifier of the promise.
// returnByValue - Whether the result is expected to be a JSON object that should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// Returns -  result - Promise result. Will contain rejected value if promise was rejected. exceptionDetails - Exception details if stack strace is available.
func (c *Runtime) AwaitPromise(ctx context.Context, promiseObjectId string, returnByValue bool, generatePreview bool) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	var v RuntimeAwaitPromiseParams
	v.PromiseObjectId = promiseObjectId
	v.ReturnByValue = returnByValue
	v.GeneratePreview = generatePreview
	return c.AwaitPromiseWithParams(ctx, &v)
}

type RuntimeCallFunctionOnParams struct {
	// Declaration of the function to call.
	FunctionDeclaration string `json:"functionDeclaration"`
	// Identifier of the object to call function on. Either objectId or executionContextId should be specified.
	ObjectId string `json:"objectId,omitempty"`
	// Call arguments. All call arguments must belong to the same JavaScript world as the target object.
	Arguments []*RuntimeCallArgument `json:"arguments,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
	Silent bool `json:"silent,omitempty"`
	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// Whether execution should be treated as initiated by user in the UI.
	UserGesture bool `json:"userGesture,omitempty"`
	// Whether execution should `await` for resulting value and return once awaited promise is resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
	// Specifies execution context which global object will be used to call function on. Either executionContextId or objectId should be specified.
	ExecutionContextId int `json:"executionContextId,omitempty"`
	// Symbolic group name that can be used to release multiple objects. If objectGroup is not specified and objectId is, objectGroup will be inherited from object.
	ObjectGroup string `json:"objectGroup,omitempty"`
	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	ThrowOnSideEffect bool `json:"throwOnSideEffect,omitempty"`
}

// CallFunctionOnWithParams - Calls function with given declaration on the given object. Object group of the result is inherited from the target object.
// Returns -  result - Call result. exceptionDetails - Exception details.
func (c *Runtime) CallFunctionOnWithParams(ctx context.Context, v *RuntimeCallFunctionOnParams) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.callFunctionOn", Params: v})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			Result           *RuntimeRemoteObject
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.ExceptionDetails, nil
}

// CallFunctionOn - Calls function with given declaration on the given object. Object group of the result is inherited from the target object.
// functionDeclaration - Declaration of the function to call.
// objectId - Identifier of the object to call function on. Either objectId or executionContextId should be specified.
// arguments - Call arguments. All call arguments must belong to the same JavaScript world as the target object.
// silent - In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
// returnByValue - Whether the result is expected to be a JSON object which should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// userGesture - Whether execution should be treated as initiated by user in the UI.
// awaitPromise - Whether execution should `await` for resulting value and return once awaited promise is resolved.
// executionContextId - Specifies execution context which global object will be used to call function on. Either executionContextId or objectId should be specified.
// objectGroup - Symbolic group name that can be used to release multiple objects. If objectGroup is not specified and objectId is, objectGroup will be inherited from object.
// throwOnSideEffect - Whether to throw an exception if side effect cannot be ruled out during evaluation.
// Returns -  result - Call result. exceptionDetails - Exception details.
func (c *Runtime) CallFunctionOn(ctx context.Context, functionDeclaration string, objectId string, arguments []*RuntimeCallArgument, silent bool, returnByValue bool, generatePreview bool, userGesture bool, awaitPromise bool, executionContextId int, objectGroup string, throwOnSideEffect bool) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	var v RuntimeCallFunctionOnParams
	v.FunctionDeclaration = functionDeclaration
	v.ObjectId = objectId
	v.Arguments = arguments
	v.Silent = silent
	v.ReturnByValue = returnByValue
	v.GeneratePreview = generatePreview
	v.UserGesture = userGesture
	v.AwaitPromise = awaitPromise
	v.ExecutionContextId = executionContextId
	v.ObjectGroup = objectGroup
	v.ThrowOnSideEffect = throwOnSideEffect
	return c.CallFunctionOnWithParams(ctx, &v)
}

type RuntimeCompileScriptParams struct {
	// Expression to compile.
	Expression string `json:"expression"`
	// Source url to be set for the script.
	SourceURL string `json:"sourceURL"`
	// Specifies whether the compiled script should be persisted.
	PersistScript bool `json:"persistScript"`
	// Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ExecutionContextId int `json:"executionContextId,omitempty"`
}

// CompileScriptWithParams - Compiles expression.
// Returns -  scriptId - Id of the script. exceptionDetails - Exception details.
func (c *Runtime) CompileScriptWithParams(ctx context.Context, v *RuntimeCompileScriptParams) (string, *RuntimeExceptionDetails, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.compileScript", Params: v})
	if err != nil {
		return "", nil, err
	}

	var chromeData struct {
		Result struct {
			ScriptId         string
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return "", nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", nil, err
	}

	return chromeData.Result.ScriptId, chromeData.Result.ExceptionDetails, nil
}

// CompileScript - Compiles expression.
// expression - Expression to compile.
// sourceURL - Source url to be set for the script.
// persistScript - Specifies whether the compiled script should be persisted.
// executionContextId - Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
// Returns -  scriptId - Id of the script. exceptionDetails - Exception details.
func (c *Runtime) CompileScript(ctx context.Context, expression string, sourceURL string, persistScript bool, executionContextId int) (string, *RuntimeExceptionDetails, error) {
	var v RuntimeCompileScriptParams
	v.Expression = expression
	v.SourceURL = sourceURL
	v.PersistScript = persistScript
	v.ExecutionContextId = executionContextId
	return c.CompileScriptWithParams(ctx, &v)
}

// Disables reporting of execution contexts creation.
func (c *Runtime) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.disable"})
}

// Discards collected exceptions and console API calls.
func (c *Runtime) DiscardConsoleEntries(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.discardConsoleEntries"})
}

// Enables reporting of execution contexts creation by means of `executionContextCreated` event. When the reporting gets enabled the event will be sent immediately for each existing execution context.
func (c *Runtime) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.enable"})
}

type RuntimeEvaluateParams struct {
	// Expression to evaluate.
	Expression string `json:"expression"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
	Silent bool `json:"silent,omitempty"`
	// Specifies in which execution context to perform evaluation. If the parameter is omitted the evaluation will be performed in the context of the inspected page. This is mutually exclusive with `uniqueContextId`, which offers an alternative way to identify the execution context that is more reliable in a multi-process environment.
	ContextId int `json:"contextId,omitempty"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// Whether execution should be treated as initiated by user in the UI.
	UserGesture bool `json:"userGesture,omitempty"`
	// Whether execution should `await` for resulting value and return once awaited promise is resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
	// Whether to throw an exception if side effect cannot be ruled out during evaluation. This implies `disableBreaks` below.
	ThrowOnSideEffect bool `json:"throwOnSideEffect,omitempty"`
	// Terminate execution after timing out (number of milliseconds).
	Timeout float64 `json:"timeout,omitempty"`
	// Disable breakpoints during execution.
	DisableBreaks bool `json:"disableBreaks,omitempty"`
	// Setting this flag to true enables `let` re-declaration and top-level `await`. Note that `let` variables can only be re-declared if they originate from `replMode` themselves.
	ReplMode bool `json:"replMode,omitempty"`
	// The Content Security Policy (CSP) for the target might block 'unsafe-eval' which includes eval(), Function(), setTimeout() and setInterval() when called with non-callable arguments. This flag bypasses CSP for this evaluation and allows unsafe-eval. Defaults to true.
	AllowUnsafeEvalBlockedByCSP bool `json:"allowUnsafeEvalBlockedByCSP,omitempty"`
	// An alternative way to specify the execution context to evaluate in. Compared to contextId that may be reused across processes, this is guaranteed to be system-unique, so it can be used to prevent accidental evaluation of the expression in context different than intended (e.g. as a result of navigation across process boundaries). This is mutually exclusive with `contextId`.
	UniqueContextId string `json:"uniqueContextId,omitempty"`
}

// EvaluateWithParams - Evaluates expression on global object.
// Returns -  result - Evaluation result. exceptionDetails - Exception details.
func (c *Runtime) EvaluateWithParams(ctx context.Context, v *RuntimeEvaluateParams) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.evaluate", Params: v})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			Result           *RuntimeRemoteObject
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.ExceptionDetails, nil
}

// Evaluate - Evaluates expression on global object.
// expression - Expression to evaluate.
// objectGroup - Symbolic group name that can be used to release multiple objects.
// includeCommandLineAPI - Determines whether Command Line API should be available during the evaluation.
// silent - In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
// contextId - Specifies in which execution context to perform evaluation. If the parameter is omitted the evaluation will be performed in the context of the inspected page. This is mutually exclusive with `uniqueContextId`, which offers an alternative way to identify the execution context that is more reliable in a multi-process environment.
// returnByValue - Whether the result is expected to be a JSON object that should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// userGesture - Whether execution should be treated as initiated by user in the UI.
// awaitPromise - Whether execution should `await` for resulting value and return once awaited promise is resolved.
// throwOnSideEffect - Whether to throw an exception if side effect cannot be ruled out during evaluation. This implies `disableBreaks` below.
// timeout - Terminate execution after timing out (number of milliseconds).
// disableBreaks - Disable breakpoints during execution.
// replMode - Setting this flag to true enables `let` re-declaration and top-level `await`. Note that `let` variables can only be re-declared if they originate from `replMode` themselves.
// allowUnsafeEvalBlockedByCSP - The Content Security Policy (CSP) for the target might block 'unsafe-eval' which includes eval(), Function(), setTimeout() and setInterval() when called with non-callable arguments. This flag bypasses CSP for this evaluation and allows unsafe-eval. Defaults to true.
// uniqueContextId - An alternative way to specify the execution context to evaluate in. Compared to contextId that may be reused across processes, this is guaranteed to be system-unique, so it can be used to prevent accidental evaluation of the expression in context different than intended (e.g. as a result of navigation across process boundaries). This is mutually exclusive with `contextId`.
// Returns -  result - Evaluation result. exceptionDetails - Exception details.
func (c *Runtime) Evaluate(ctx context.Context, expression string, objectGroup string, includeCommandLineAPI bool, silent bool, contextId int, returnByValue bool, generatePreview bool, userGesture bool, awaitPromise bool, throwOnSideEffect bool, timeout float64, disableBreaks bool, replMode bool, allowUnsafeEvalBlockedByCSP bool, uniqueContextId string) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	var v RuntimeEvaluateParams
	v.Expression = expression
	v.ObjectGroup = objectGroup
	v.IncludeCommandLineAPI = includeCommandLineAPI
	v.Silent = silent
	v.ContextId = contextId
	v.ReturnByValue = returnByValue
	v.GeneratePreview = generatePreview
	v.UserGesture = userGesture
	v.AwaitPromise = awaitPromise
	v.ThrowOnSideEffect = throwOnSideEffect
	v.Timeout = timeout
	v.DisableBreaks = disableBreaks
	v.ReplMode = replMode
	v.AllowUnsafeEvalBlockedByCSP = allowUnsafeEvalBlockedByCSP
	v.UniqueContextId = uniqueContextId
	return c.EvaluateWithParams(ctx, &v)
}

// GetIsolateId - Returns the isolate id.
// Returns -  id - The isolate id.
func (c *Runtime) GetIsolateId(ctx context.Context) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.getIsolateId"})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Id string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.Id, nil
}

// GetHeapUsage - Returns the JavaScript heap usage. It is the total usage of the corresponding isolate not scoped to a particular Runtime.
// Returns -  usedSize - Used heap size in bytes. totalSize - Allocated heap size in bytes.
func (c *Runtime) GetHeapUsage(ctx context.Context) (float64, float64, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.getHeapUsage"})
	if err != nil {
		return 0, 0, err
	}

	var chromeData struct {
		Result struct {
			UsedSize  float64
			TotalSize float64
		}
	}

	if resp == nil {
		return 0, 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return 0, 0, err
	}

	return chromeData.Result.UsedSize, chromeData.Result.TotalSize, nil
}

type RuntimeGetPropertiesParams struct {
	// Identifier of the object to return properties for.
	ObjectId string `json:"objectId"`
	// If true, returns properties belonging only to the element itself, not to its prototype chain.
	OwnProperties bool `json:"ownProperties,omitempty"`
	// If true, returns accessor properties (with getter/setter) only; internal properties are not returned either.
	AccessorPropertiesOnly bool `json:"accessorPropertiesOnly,omitempty"`
	// Whether preview should be generated for the results.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// If true, returns non-indexed properties only.
	NonIndexedPropertiesOnly bool `json:"nonIndexedPropertiesOnly,omitempty"`
}

// GetPropertiesWithParams - Returns properties of a given object. Object group of the result is inherited from the target object.
// Returns -  result - Object properties. internalProperties - Internal object properties (only of the element itself). privateProperties - Object private properties. exceptionDetails - Exception details.
func (c *Runtime) GetPropertiesWithParams(ctx context.Context, v *RuntimeGetPropertiesParams) ([]*RuntimePropertyDescriptor, []*RuntimeInternalPropertyDescriptor, []*RuntimePrivatePropertyDescriptor, *RuntimeExceptionDetails, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.getProperties", Params: v})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			Result             []*RuntimePropertyDescriptor
			InternalProperties []*RuntimeInternalPropertyDescriptor
			PrivateProperties  []*RuntimePrivatePropertyDescriptor
			ExceptionDetails   *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, nil, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.InternalProperties, chromeData.Result.PrivateProperties, chromeData.Result.ExceptionDetails, nil
}

// GetProperties - Returns properties of a given object. Object group of the result is inherited from the target object.
// objectId - Identifier of the object to return properties for.
// ownProperties - If true, returns properties belonging only to the element itself, not to its prototype chain.
// accessorPropertiesOnly - If true, returns accessor properties (with getter/setter) only; internal properties are not returned either.
// generatePreview - Whether preview should be generated for the results.
// nonIndexedPropertiesOnly - If true, returns non-indexed properties only.
// Returns -  result - Object properties. internalProperties - Internal object properties (only of the element itself). privateProperties - Object private properties. exceptionDetails - Exception details.
func (c *Runtime) GetProperties(ctx context.Context, objectId string, ownProperties bool, accessorPropertiesOnly bool, generatePreview bool, nonIndexedPropertiesOnly bool) ([]*RuntimePropertyDescriptor, []*RuntimeInternalPropertyDescriptor, []*RuntimePrivatePropertyDescriptor, *RuntimeExceptionDetails, error) {
	var v RuntimeGetPropertiesParams
	v.ObjectId = objectId
	v.OwnProperties = ownProperties
	v.AccessorPropertiesOnly = accessorPropertiesOnly
	v.GeneratePreview = generatePreview
	v.NonIndexedPropertiesOnly = nonIndexedPropertiesOnly
	return c.GetPropertiesWithParams(ctx, &v)
}

type RuntimeGlobalLexicalScopeNamesParams struct {
	// Specifies in which execution context to lookup global scope variables.
	ExecutionContextId int `json:"executionContextId,omitempty"`
}

// GlobalLexicalScopeNamesWithParams - Returns all let, const and class variables from global scope.
// Returns -  names -
func (c *Runtime) GlobalLexicalScopeNamesWithParams(ctx context.Context, v *RuntimeGlobalLexicalScopeNamesParams) ([]string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.globalLexicalScopeNames", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Names []string
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Names, nil
}

// GlobalLexicalScopeNames - Returns all let, const and class variables from global scope.
// executionContextId - Specifies in which execution context to lookup global scope variables.
// Returns -  names -
func (c *Runtime) GlobalLexicalScopeNames(ctx context.Context, executionContextId int) ([]string, error) {
	var v RuntimeGlobalLexicalScopeNamesParams
	v.ExecutionContextId = executionContextId
	return c.GlobalLexicalScopeNamesWithParams(ctx, &v)
}

type RuntimeQueryObjectsParams struct {
	// Identifier of the prototype to return objects for.
	PrototypeObjectId string `json:"prototypeObjectId"`
	// Symbolic group name that can be used to release the results.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

// QueryObjectsWithParams -
// Returns -  objects - Array with objects.
func (c *Runtime) QueryObjectsWithParams(ctx context.Context, v *RuntimeQueryObjectsParams) (*RuntimeRemoteObject, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.queryObjects", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Objects *RuntimeRemoteObject
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Objects, nil
}

// QueryObjects -
// prototypeObjectId - Identifier of the prototype to return objects for.
// objectGroup - Symbolic group name that can be used to release the results.
// Returns -  objects - Array with objects.
func (c *Runtime) QueryObjects(ctx context.Context, prototypeObjectId string, objectGroup string) (*RuntimeRemoteObject, error) {
	var v RuntimeQueryObjectsParams
	v.PrototypeObjectId = prototypeObjectId
	v.ObjectGroup = objectGroup
	return c.QueryObjectsWithParams(ctx, &v)
}

type RuntimeReleaseObjectParams struct {
	// Identifier of the object to release.
	ObjectId string `json:"objectId"`
}

// ReleaseObjectWithParams - Releases remote object with given id.
func (c *Runtime) ReleaseObjectWithParams(ctx context.Context, v *RuntimeReleaseObjectParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.releaseObject", Params: v})
}

// ReleaseObject - Releases remote object with given id.
// objectId - Identifier of the object to release.
func (c *Runtime) ReleaseObject(ctx context.Context, objectId string) (*gcdmessage.ChromeResponse, error) {
	var v RuntimeReleaseObjectParams
	v.ObjectId = objectId
	return c.ReleaseObjectWithParams(ctx, &v)
}

type RuntimeReleaseObjectGroupParams struct {
	// Symbolic object group name.
	ObjectGroup string `json:"objectGroup"`
}

// ReleaseObjectGroupWithParams - Releases all remote objects that belong to a given group.
func (c *Runtime) ReleaseObjectGroupWithParams(ctx context.Context, v *RuntimeReleaseObjectGroupParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.releaseObjectGroup", Params: v})
}

// ReleaseObjectGroup - Releases all remote objects that belong to a given group.
// objectGroup - Symbolic object group name.
func (c *Runtime) ReleaseObjectGroup(ctx context.Context, objectGroup string) (*gcdmessage.ChromeResponse, error) {
	var v RuntimeReleaseObjectGroupParams
	v.ObjectGroup = objectGroup
	return c.ReleaseObjectGroupWithParams(ctx, &v)
}

// Tells inspected instance to run if it was waiting for debugger to attach.
func (c *Runtime) RunIfWaitingForDebugger(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.runIfWaitingForDebugger"})
}

type RuntimeRunScriptParams struct {
	// Id of the script to run.
	ScriptId string `json:"scriptId"`
	// Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ExecutionContextId int `json:"executionContextId,omitempty"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
	Silent bool `json:"silent,omitempty"`
	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`
	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// Whether execution should `await` for resulting value and return once awaited promise is resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
}

// RunScriptWithParams - Runs script with given id in a given context.
// Returns -  result - Run result. exceptionDetails - Exception details.
func (c *Runtime) RunScriptWithParams(ctx context.Context, v *RuntimeRunScriptParams) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.runScript", Params: v})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			Result           *RuntimeRemoteObject
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.ExceptionDetails, nil
}

// RunScript - Runs script with given id in a given context.
// scriptId - Id of the script to run.
// executionContextId - Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
// objectGroup - Symbolic group name that can be used to release multiple objects.
// silent - In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
// includeCommandLineAPI - Determines whether Command Line API should be available during the evaluation.
// returnByValue - Whether the result is expected to be a JSON object which should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// awaitPromise - Whether execution should `await` for resulting value and return once awaited promise is resolved.
// Returns -  result - Run result. exceptionDetails - Exception details.
func (c *Runtime) RunScript(ctx context.Context, scriptId string, executionContextId int, objectGroup string, silent bool, includeCommandLineAPI bool, returnByValue bool, generatePreview bool, awaitPromise bool) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	var v RuntimeRunScriptParams
	v.ScriptId = scriptId
	v.ExecutionContextId = executionContextId
	v.ObjectGroup = objectGroup
	v.Silent = silent
	v.IncludeCommandLineAPI = includeCommandLineAPI
	v.ReturnByValue = returnByValue
	v.GeneratePreview = generatePreview
	v.AwaitPromise = awaitPromise
	return c.RunScriptWithParams(ctx, &v)
}

type RuntimeSetAsyncCallStackDepthParams struct {
	// Maximum depth of async call stacks. Setting to `0` will effectively disable collecting async call stacks (default).
	MaxDepth int `json:"maxDepth"`
}

// SetAsyncCallStackDepthWithParams - Enables or disables async call stacks tracking.
func (c *Runtime) SetAsyncCallStackDepthWithParams(ctx context.Context, v *RuntimeSetAsyncCallStackDepthParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.setAsyncCallStackDepth", Params: v})
}

// SetAsyncCallStackDepth - Enables or disables async call stacks tracking.
// maxDepth - Maximum depth of async call stacks. Setting to `0` will effectively disable collecting async call stacks (default).
func (c *Runtime) SetAsyncCallStackDepth(ctx context.Context, maxDepth int) (*gcdmessage.ChromeResponse, error) {
	var v RuntimeSetAsyncCallStackDepthParams
	v.MaxDepth = maxDepth
	return c.SetAsyncCallStackDepthWithParams(ctx, &v)
}

type RuntimeSetCustomObjectFormatterEnabledParams struct {
	//
	Enabled bool `json:"enabled"`
}

// SetCustomObjectFormatterEnabledWithParams -
func (c *Runtime) SetCustomObjectFormatterEnabledWithParams(ctx context.Context, v *RuntimeSetCustomObjectFormatterEnabledParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.setCustomObjectFormatterEnabled", Params: v})
}

// SetCustomObjectFormatterEnabled -
// enabled -
func (c *Runtime) SetCustomObjectFormatterEnabled(ctx context.Context, enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v RuntimeSetCustomObjectFormatterEnabledParams
	v.Enabled = enabled
	return c.SetCustomObjectFormatterEnabledWithParams(ctx, &v)
}

type RuntimeSetMaxCallStackSizeToCaptureParams struct {
	//
	Size int `json:"size"`
}

// SetMaxCallStackSizeToCaptureWithParams -
func (c *Runtime) SetMaxCallStackSizeToCaptureWithParams(ctx context.Context, v *RuntimeSetMaxCallStackSizeToCaptureParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.setMaxCallStackSizeToCapture", Params: v})
}

// SetMaxCallStackSizeToCapture -
// size -
func (c *Runtime) SetMaxCallStackSizeToCapture(ctx context.Context, size int) (*gcdmessage.ChromeResponse, error) {
	var v RuntimeSetMaxCallStackSizeToCaptureParams
	v.Size = size
	return c.SetMaxCallStackSizeToCaptureWithParams(ctx, &v)
}

// Terminate current or next JavaScript execution. Will cancel the termination when the outer-most script execution ends.
func (c *Runtime) TerminateExecution(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.terminateExecution"})
}

type RuntimeAddBindingParams struct {
	//
	Name string `json:"name"`
	// If specified, the binding would only be exposed to the specified execution context. If omitted and `executionContextName` is not set, the binding is exposed to all execution contexts of the target. This parameter is mutually exclusive with `executionContextName`. Deprecated in favor of `executionContextName` due to an unclear use case and bugs in implementation (crbug.com/1169639). `executionContextId` will be removed in the future.
	ExecutionContextId int `json:"executionContextId,omitempty"`
	// If specified, the binding is exposed to the executionContext with matching name, even for contexts created after the binding is added. See also `ExecutionContext.name` and `worldName` parameter to `Page.addScriptToEvaluateOnNewDocument`. This parameter is mutually exclusive with `executionContextId`.
	ExecutionContextName string `json:"executionContextName,omitempty"`
}

// AddBindingWithParams - If executionContextId is empty, adds binding with the given name on the global objects of all inspected contexts, including those created later, bindings survive reloads. Binding function takes exactly one argument, this argument should be string, in case of any other input, function throws an exception. Each binding function call produces Runtime.bindingCalled notification.
func (c *Runtime) AddBindingWithParams(ctx context.Context, v *RuntimeAddBindingParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.addBinding", Params: v})
}

// AddBinding - If executionContextId is empty, adds binding with the given name on the global objects of all inspected contexts, including those created later, bindings survive reloads. Binding function takes exactly one argument, this argument should be string, in case of any other input, function throws an exception. Each binding function call produces Runtime.bindingCalled notification.
// name -
// executionContextId - If specified, the binding would only be exposed to the specified execution context. If omitted and `executionContextName` is not set, the binding is exposed to all execution contexts of the target. This parameter is mutually exclusive with `executionContextName`. Deprecated in favor of `executionContextName` due to an unclear use case and bugs in implementation (crbug.com/1169639). `executionContextId` will be removed in the future.
// executionContextName - If specified, the binding is exposed to the executionContext with matching name, even for contexts created after the binding is added. See also `ExecutionContext.name` and `worldName` parameter to `Page.addScriptToEvaluateOnNewDocument`. This parameter is mutually exclusive with `executionContextId`.
func (c *Runtime) AddBinding(ctx context.Context, name string, executionContextId int, executionContextName string) (*gcdmessage.ChromeResponse, error) {
	var v RuntimeAddBindingParams
	v.Name = name
	v.ExecutionContextId = executionContextId
	v.ExecutionContextName = executionContextName
	return c.AddBindingWithParams(ctx, &v)
}

type RuntimeRemoveBindingParams struct {
	//
	Name string `json:"name"`
}

// RemoveBindingWithParams - This method does not remove binding function from global object but unsubscribes current runtime agent from Runtime.bindingCalled notifications.
func (c *Runtime) RemoveBindingWithParams(ctx context.Context, v *RuntimeRemoveBindingParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.removeBinding", Params: v})
}

// RemoveBinding - This method does not remove binding function from global object but unsubscribes current runtime agent from Runtime.bindingCalled notifications.
// name -
func (c *Runtime) RemoveBinding(ctx context.Context, name string) (*gcdmessage.ChromeResponse, error) {
	var v RuntimeRemoveBindingParams
	v.Name = name
	return c.RemoveBindingWithParams(ctx, &v)
}
