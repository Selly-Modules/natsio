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

func (s Supplier) GetSupplierInfo(supplierID string) (*model.ResponseSupplierInfo, error) {
	msg, err := natsio.GetServer().Request(subject.Supplier.GetSupplierInfo, toBytes(supplierID))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.ResponseSupplierInfo `json:"data"`
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
