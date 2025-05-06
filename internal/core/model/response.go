package model

type Response struct {
	StatusCode bool   `json:"status_code"`
	RequestID  string `json:"request_id"`
	Message    string `json:"message"`
	Payload    any    `json:"payload,omitempty"`
	Error      *Error `json:"error,omitempty"`
}

type Error struct {
	Message   string `json:"message"`
	Code      string `json:"code"`
	OtherInfo any    `json:"otherInfo"`
}
