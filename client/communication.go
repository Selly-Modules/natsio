package client

import (
	"encoding/json"

	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
)

// Communication ...
type Communication struct{}

// GetCommunication ...
func GetCommunication() Communication {
	return Communication{}
}

// RequestHttp ...
func (c Communication) RequestHttp(p model.CommunicationRequestHttp) (r *model.CommunicationHttpResponse, err error) {
	msg, err := natsio.GetServer().Request(subject.CommunicationRequestHTTP, toBytes(p))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(msg.Data, &r)
	return r, err
}
