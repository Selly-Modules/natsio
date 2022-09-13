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

// GetLocationByCode ...
func (l Location) GetLocationByCode(payload model.LocationRequestPayload) (*model.ResponseLocationAddress, error) {
	msg, err := natsio.GetServer().Request(subject.Location.GetLocationByCode, toBytes(payload))
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

// GetProvincesByCodes ... ...
func (l Location) GetProvincesByCodes(p model.ProvinceRequestPayload) (*model.LocationProvinceResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.GetProvincesByCodes, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.LocationProvinceResponse `json:"data"'`
		Error string                          `json:"error"`
	}

	if err := json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}

// GetDistrictsByCodes ...
func (l Location) GetDistrictsByCodes(p model.DistrictRequestPayload) (*model.LocationDistrictResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.GetDistrictsByCodes, toBytes(p))
	if err != nil {
		return nil, err
	}
	var r struct {
		Data  *model.LocationDistrictResponse `json:"data"`
		Error string                          `json:"error"`
	}

	if err := json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}

	return r.Data, nil
}

// GetWardsByCodes ...
func (l Location) GetWardsByCodes(p model.WardRequestPayload) (*model.LocationWardResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.GetWardsByCodes, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.LocationWardResponse `json:"data"`
		Error string                      `json:"error"`
	}

	if err := json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}

	return r.Data, nil
}
