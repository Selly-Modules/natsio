package client

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
)

// DistinctWithField ...
func (w Warehouse) DistinctWithField(p model.DistinctWithField) ([]interface{}, error) {
	msg, err := natsio.GetServer().Request(subject.Warehouse.Distinct, bsonToBytes(p))
	if err != nil {
		return nil, err
	}
	var r struct {
		Data  []interface{} `json:"data"`
		Error string        `json:"error"`
	}
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}
	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}

// FindOneByCondition ...
func (w Warehouse) FindOneByCondition(p model.FindOneCondition) (*model.WarehouseNatsResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Warehouse.FindOne, bsonToBytes(p))
	if err != nil {
		return nil, err
	}
	var r struct {
		Data  *model.WarehouseNatsResponse `json:"data"`
		Error string                       `json:"error"`
	}
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}
	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}

// CountByCondition ...
func (w Warehouse) CountByCondition(p model.FindWithCondition) (int64, error) {
	msg, err := natsio.GetServer().Request(subject.Warehouse.Count, bsonToBytes(p))
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

// FindByCondition ...
func (w Warehouse) FindByCondition(p model.FindWithCondition) ([]*model.WarehouseNatsResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Warehouse.FindByCondition, bsonToBytes(p))
	log.Println("find_warehouse: ", string(msg.Data), err)
	if err != nil {
		return nil, err
	}
	var r struct {
		Data  []*model.WarehouseNatsResponse `json:"data"`
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
