package servicebus

// Request to send to servicebus call
type Request struct {
	Jsonrpc        string  `json:"jsonrpc,omitempty"`
	Method         string  `json:"method,omitempty"`
	Params         *Params `json:"params,omitempty"`
	ID             int32   `json:"id,omitempty"`
	Asyncmessageid int32   `json:"asyncmessageid,omitempty"`
	Traceid        int32   `json:"traceid,omitempty"`
}

// Params are part of the request sent to servicebus call
type Params struct {
	DelayMilliSeconds int32 `json:"delayMilliseconds,omitempty"`
}

// Response from servicebus
type Response struct {
	Jsonrpc string  `json:"jsonrpc,omitempty"`
	ID      int32   `json:"id,omitempty"`
	Result  *Result `json:"result,omitempty"`
	Error   *Error  `json:"error,omitempty"`
}

// Result part of the response
type Result struct {
	PingResponse string `json:"PingResponse,omitempty"`
}

// Error part of the response
type Error struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}
