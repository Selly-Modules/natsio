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

// UpdateIsClosedSupplier ...
func (w Warehouse) UpdateIsClosedSupplier(p model.UpdateSupplierIsClosedRequest) error {
	msg, err := natsio.GetServer().Request(subject.Warehouse.UpdateIsClosedSupplier, toBytes(p))
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

// AfterCreateWarehouse ...
func (w Warehouse) AfterCreateWarehouse(p model.WarehouseNatsResponse) error {
	msg, err := natsio.GetServer().Request(subject.Warehouse.AfterCreateWarehouse, toBytes(p))
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

// AfterUpdateWarehouse ...
func (w Warehouse) AfterUpdateWarehouse(p model.WarehouseNatsResponse) error {
	msg, err := natsio.GetServer().Request(subject.Warehouse.AfterUpdateWarehouse, toBytes(p))
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

// GetWarehouses ...
func (w Warehouse) GetWarehouses(p model.GetWarehousesRequest) (*model.GetWarehousesResponse, error) {
	msg, err := natsio.GetServer().Request(subject.Warehouse.GetWarehouses, toBytes(p))
	if err != nil {
		return nil, err
	}
	var r struct {
		Data  *model.GetWarehousesResponse `json:"data"`
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
