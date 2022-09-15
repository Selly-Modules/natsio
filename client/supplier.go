package client

import (
	"encoding/json"
	"errors"

	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
)

// Supplier ...
type Supplier struct{}

// GetSupplier ...
func GetSupplier() Supplier {
	return Supplier{}
}

func (s Supplier) GetListSupplierInfo(p model.GetSupplierRequest) ([]*model.ResponseSupplierInfo, error) {
	msg, err := natsio.GetServer().Request(subject.Supplier.GetListSupplierInfo, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  []*model.ResponseSupplierInfo `json:"data"`
		Error string                        `json:"error"`
	}

	if err := json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}
	if r.Error != "" {
		return nil, errors.New(r.Error)
	}

	return r.Data, nil
}

func (s Supplier) GetSupplierContractBySupplierID(p model.GetSupplierContractRequest) (*model.ResponseSupplierContract, error) {
	msg, err := natsio.GetServer().Request(subject.Supplier.GetSupplierContractBySupplierID, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.ResponseSupplierContract `json:"data"`
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

func (s Supplier) FindAll(supplierID model.SupplierRequestPayload) (*model.SupplierAll, error) {
	msg, err := natsio.GetServer().Request(subject.Supplier.FindAll, toBytes(supplierID))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.SupplierAll `json:"data"`
		Error string             `json:"error"`
	}

	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}
	if r.Error != "" {
		return nil, errors.New(r.Error)
	}

	return r.Data, nil
}

// CreateWarehouseIntoServiceSupplier ...
func (s Supplier) CreateWarehouseIntoServiceSupplier(p model.CreateSupplierWarehousePayload) error {
	msg, err := natsio.GetServer().Request(subject.Warehouse.CreateWarehouseIntoServiceSupplier, toBytes(p))
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

// UpdateWarehouseIntoServiceSupplier ...
func (s Supplier) UpdateWarehouseIntoServiceSupplier(p model.UpdateSupplierWarehousePayload) error {
	msg, err := natsio.GetServer().Request(subject.Warehouse.UpdateWarehouseIntoServiceSupplier, toBytes(p))
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
