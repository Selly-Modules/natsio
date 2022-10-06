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
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}

// GetProvincesByCodes ...
func (l Location) GetProvincesByCodes(p model.ProvinceRequestPayload) (*model.LocationProvinceResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.GetProvincesByCodes, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.LocationProvinceResponse `json:"data"`
		Error string                          `json:"error"`
	}

	if err = json.Unmarshal(msg.Data, &r); err != nil {
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

	if err = json.Unmarshal(msg.Data, &r); err != nil {
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

	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}

	return r.Data, nil
}

// GetProvinceByCondition ...
func (l Location) GetProvinceByCondition(p model.ProvinceRequestCondition) (*model.LocationProvinceDetailResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.GetProvinceByCondition, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.LocationProvinceDetailResponse `json:"data"`
		Error string                                `json:"error"`
	}

	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}

// GetProvincesByCondition ...
func (l Location) GetProvincesByCondition(p model.ProvinceRequestCondition) ([]*model.LocationProvinceDetailResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.GetProvincesByCondition, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  []*model.LocationProvinceDetailResponse `json:"data"`
		Error string                                  `json:"error"`
	}

	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}

// GetDistrictByCondition ...
func (l Location) GetDistrictByCondition(p model.DistrictRequestCondition) (*model.LocationDistrictDetailResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.GetDistrictByCondition, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.LocationDistrictDetailResponse `json:"data"`
		Error string                                `json:"error"`
	}

	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}

// GetDistrictsByCondition ...
func (l Location) GetDistrictsByCondition(p model.DistrictRequestCondition) ([]*model.LocationDistrictDetailResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.GetDistrictsByCondition, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  []*model.LocationDistrictDetailResponse `json:"data"`
		Error string                                  `json:"error"`
	}

	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}

// GetWardByCondition ...
func (l Location) GetWardByCondition(p model.WardRequestCondition) (*model.LocationWardDetailResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.GetWardByCondition, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.LocationWardDetailResponse `json:"data"`
		Error string                            `json:"error"`
	}

	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}

// GetWardsByCondition ...
func (l Location) GetWardsByCondition(p model.WardRequestCondition) ([]*model.LocationWardDetailResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.GetWardsByCondition, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  []*model.LocationWardDetailResponse `json:"data"`
		Error string                              `json:"error"`
	}

	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}

// CountProvinceByCondition ...
func (l Location) CountProvinceByCondition(p model.ProvinceRequestCondition) (int64, error) {
	msg, err := natsio.GetServer().Request(subject.Location.CountProvinceByCondition, toBytes(p))
	if err != nil {
		return 0, err
	}
	var r struct {
		Data  int64  `json:"data"`
		Error string `json:"error"`
	}
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return 0, err
	}
	if r.Error != "" {
		return 0, errors.New(r.Error)
	}
	return r.Data, nil
}

// CountDistrictByCondition ...
func (l Location) CountDistrictByCondition(p model.DistrictRequestCondition) (int64, error) {
	msg, err := natsio.GetServer().Request(subject.Location.CountDistrictByCondition, toBytes(p))
	if err != nil {
		return 0, err
	}
	var r struct {
		Data  int64  `json:"data"`
		Error string `json:"error"`
	}
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return 0, err
	}
	if r.Error != "" {
		return 0, errors.New(r.Error)
	}
	return r.Data, nil
}

// CountWardByCondition ...
func (l Location) CountWardByCondition(p model.WardRequestCondition) (int64, error) {
	msg, err := natsio.GetServer().Request(subject.Location.CountWardByCondition, toBytes(p))
	if err != nil {
		return 0, err
	}
	var r struct {
		Data  int64  `json:"data"`
		Error string `json:"error"`
	}
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return 0, err
	}
	if r.Error != "" {
		return 0, errors.New(r.Error)
	}
	return r.Data, nil
}
