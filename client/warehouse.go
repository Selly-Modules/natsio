package client

import (
	"encoding/json"
	"errors"

	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
)

// Warehouse ...
type Warehouse struct{}

// GetWarehouse ...
func GetWarehouse() Warehouse {
	return Warehouse{}
}

// DistinctWithField ...
func (w Warehouse) DistinctWithField(p model.FindWithCondition) ([]interface{}, error) {
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
func (w Warehouse) FindOneByCondition(p model.FindWithCondition) (*model.WarehouseNatsResponse, error) {
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

// CreateOutboundRequest ...
func (w Warehouse) CreateOutboundRequest(p model.OutboundRequestPayload) (*model.OutboundRequestResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Warehouse.CreateOutboundRequest, toBytes(p))
	if err != nil {
		return nil, err
	}
	var r struct {
		Data  *model.OutboundRequestResponse `json:"data"`
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

// UpdateOutboundRequestLogisticInfo ...
func (w Warehouse) UpdateOutboundRequestLogisticInfo(p model.UpdateOutboundRequestLogisticInfoPayload) error {
	msg, err := natsio.GetServer().Request(subject.Warehouse.UpdateOutboundRequestLogistic, toBytes(p))
	if err != nil {
		return err
	}
	var r struct {
		Error string `json:"error"`
	}
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return err
	}
	if r.Error != "" {
		return errors.New(r.Error)
	}
	return nil
}

// CancelOutboundRequest ...
func (w Warehouse) CancelOutboundRequest(p model.CancelOutboundRequest) error {
	msg, err := natsio.GetServer().Request(subject.Warehouse.CancelOutboundRequest, toBytes(p))
	if err != nil {
		return err
	}
	var r struct {
		Error string `json:"error"`
	}
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return err
	}
	if r.Error != "" {
		return errors.New(r.Error)
	}
	return nil
}

// GetConfigByWarehouseID ...
func (w Warehouse) GetConfigByWarehouseID(warehouseID string) (*model.WarehouseConfiguration, error) {
	msg, err := natsio.GetServer().Request(subject.Warehouse.GetConfiguration, toBytes(warehouseID))
	if err != nil {
		return nil, err
	}
	var r struct {
		Data  *model.WarehouseConfiguration `json:"data"`
		Error string                        `json:"error"`
	}
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}
	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}
