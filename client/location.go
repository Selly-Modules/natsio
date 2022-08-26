package client

import (
	"encoding/json"
	"errors"

	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
)

// Location ...
type Location struct{}

// GetLocation ...
func GetLocation() Location {
	return Location{}
}

func (l Location) GetLocationByCode(payload model.LocationRequestPayload) (*model.ResponseLocationAddress, error) {
	msg, err := natsio.GetServer().Request(subject.GetLocationWarehouse, toBytes(payload))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.ResponseLocationAddress `json:"data"`
		Error string                         `json:"error"`
	}
	if err := json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}
