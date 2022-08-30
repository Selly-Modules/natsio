package model

import "encoding/json"

// CommunicationHttpResponse ...
type CommunicationHttpResponse struct {
	Response  *HttpResponse `json:"response"`
	Error     bool          `json:"error"`
	Message   string        `json:"message"`
	RequestID string        `json:"requestId"`
}

// ParseResponseData ...
func (r *CommunicationHttpResponse) ParseResponseData(result interface{}) error {
	if r.Response == nil {
		return nil
	}
	return json.Unmarshal([]byte(r.Response.Body), result)
}

// HttpResponse ...
type HttpResponse struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statusCode"`
}
