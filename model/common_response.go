package model

import "encoding/json"

// CommonResponseData ...
type CommonResponseData struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// ParseData ...
func (c CommonResponseData) ParseData(resultPointer interface{}) error {
	b, err := json.Marshal(c.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, resultPointer)
}
