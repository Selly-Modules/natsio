package model

// CommunicationRequestHttp ...
type CommunicationRequestHttp struct {
	ResponseImmediately bool        `json:"responseImmediately"`
	Authentication      string      `json:"authentication"`
	Payload             HttpRequest `json:"payload"`
}

// HttpRequest ...
type HttpRequest struct {
	URL    string            `json:"url"`
	Method string            `json:"method"`
	Data   string            `json:"data"`
	Header map[string]string `json:"header"`
	Query  map[string]string `json:"query"`
}
