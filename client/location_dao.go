package client

import (
	"encoding/json"
	"errors"
	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
)

// FindOneProvinceByCondition ...
func (l Location) FindOneProvinceByCondition(p model.FindOneCondition) (*model.LocationProvinceDetailResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.FindOneProvinceByCondition, bsonToBytes(p))
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

// FindOneDistrictByCondition ...
func (l Location) FindOneDistrictByCondition(p model.FindOneCondition) (*model.LocationDistrictDetailResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.FindOneDistrictByCondition, bsonToBytes(p))
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

// FindOneWardByCondition ...
func (l Location) FindOneWardByCondition(p model.FindOneCondition) (*model.LocationWardDetailResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.FindOneWardByCondition, bsonToBytes(p))
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

// FindProvinceByCondition ...
func (l Location) FindProvinceByCondition(p model.FindWithCondition) ([]*model.LocationProvinceDetailResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.FindProvinceByCondition, bsonToBytes(p))
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

// FindDistrictByCondition ...
func (l Location) FindDistrictByCondition(p model.FindWithCondition) ([]*model.LocationDistrictDetailResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.FindDistrictByCondition, bsonToBytes(p))
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

// FindWardByCondition ...
func (l Location) FindWardByCondition(p model.FindWithCondition) ([]*model.LocationWardDetailResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Location.FindWardByCondition, bsonToBytes(p))
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
func (l Location) CountProvinceByCondition(p model.FindWithCondition) (int64, error) {
	msg, err := natsio.GetServer().Request(subject.Location.CountProvinceByCondition, bsonToBytes(p))
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
func (l Location) CountDistrictByCondition(p model.FindWithCondition) (int64, error) {
	msg, err := natsio.GetServer().Request(subject.Location.CountDistrictByCondition, bsonToBytes(p))
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
func (l Location) CountWardByCondition(p model.FindWithCondition) (int64, error) {
	msg, err := natsio.GetServer().Request(subject.Location.CountWardByCondition, bsonToBytes(p))
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
