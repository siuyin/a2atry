package jsonrpc

import "encoding/json"

const (
	MessageSend                       = "message/send"
	MessageStream                     = "message/stream"
	TasksGet                          = "tasks/get"
	TasksCancel                       = "tasks/cancel"
	TasksPushNotificationConfigSet    = "tasks/pushNotificationConfig/set"
	TasksPushNotificationConfigGet    = "tasks/pushNotificationConfig/get"
	TasksPushNotificationConfigList   = "tasks/pushNotificationConfig/list"
	TasksPushNotificationConfigDelete = "tasks/pushNotificationConfig/delete"
	TasksResubscribe                  = "tasks/resubscribe"
	AgentGetAuthenticatedExtendedCard = "agent/getAuthenticatedExtendedCard"
)

// MessageIdentifier represents the base interface for identifying JSON-RPC messages
type MessageIdentifier struct {
	// ID is the request identifier. Can be a string, number, or null.
	// Responses must have the same ID as the request they relate to.
	// Notifications (requests without an expected response) should omit the ID or use null.
	ID interface{} `json:"id,omitempty"`
}

// Message represents the base interface for all JSON-RPC messages
type Message struct {
	MessageIdentifier
	// JSONRPC specifies the JSON-RPC version. Must be "2.0"
	JSONRPC string `json:"jsonrpc,omitempty"`
}

// Request represents a JSON-RPC request object base structure
type Request struct {
	Message
	// Method is the name of the method to be invoked
	Method string `json:"method"`
	// Params are the parameters for the method
	Params json.RawMessage `json:"params,omitempty"`
}

// Error represents a JSON-RPC error object
type Error struct {
	// Code is a number indicating the error type that occurred
	Code int `json:"code"`
	// Message is a string providing a short description of the error
	Message string `json:"message"`
	// Data is optional additional data about the error
	Data interface{} `json:"data,omitempty"`
}

// Response represents a JSON-RPC response object
type Response struct {
	Message
	// Result is the result of the method invocation. Required on success.
	// Should be null or omitted if an error occurred.
	Result interface{} `json:"result,omitempty"`
	// Error is an error object if an error occurred during the request.
	// Required on failure. Should be null or omitted if the request was successful.
	Error *Error `json:"error,omitempty"`
}
