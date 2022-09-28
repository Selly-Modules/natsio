package client

import (
	"encoding/json"
	"errors"

	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
)

// Order ...
type Order struct{}

// GetOrder ...
func GetOrder() Order {
	return Order{}
}

// UpdateORStatus ...
func (o Order) UpdateORStatus(p model.OrderUpdateORStatus) error {
	msg, err := natsio.GetServer().Request(subject.Order.UpdateORStatus, toBytes(p))
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

// CancelDelivery ...
func (o Order) CancelDelivery(p model.OrderCancelDelivery) error {
	msg, err := natsio.GetServer().Request(subject.Order.CancelDelivery, toBytes(p))
	if err != nil {
		return err
	}
	var r model.CommonResponseData
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return err
	}
	if r.Error != "" {
		return errors.New(r.Error)
	}
	return nil
}

// ChangeDeliveryStatus ...
func (o Order) ChangeDeliveryStatus(p model.OrderChangeDeliveryStatus) error {
	msg, err := natsio.GetServer().Request(subject.Order.ChangeDeliveryStatus, toBytes(p))
	if err != nil {
		return err
	}
	var r model.CommonResponseData
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return err
	}
	if r.Error != "" {
		return errors.New(r.Error)
	}
	return nil
}

// UpdateLogisticInfoFailed ...
func (o Order) UpdateLogisticInfoFailed(p model.OrderUpdateLogisticInfoFailed) error {
	msg, err := natsio.GetServer().Request(subject.Order.UpdateLogisticInfoFailed, toBytes(p))
	if err != nil {
		return err
	}
	var r model.CommonResponseData
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return err
	}
	if r.Error != "" {
		return errors.New(r.Error)
	}
	return nil
}

// ORNotUpdateStatus ...
func (o Order) ORNotUpdateStatus(p model.OrderORsNotUpdateStatus) error {
	msg, err := natsio.GetServer().Request(subject.Order.ORNotUpdateStatus, toBytes(p))
	if err != nil {
		return err
	}
	var r model.CommonResponseData
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return err
	}
	if r.Error != "" {
		return errors.New(r.Error)
	}
	return nil
}

// GetSupplierOrders ...
func (o Order) GetSupplierOrders(p model.OrderSupplierQuery) (*model.SupplierOrderList, error) {
	msg, err := natsio.GetServer().Request(subject.Order.GetSupplierOrders, toBytes(p))
	if err != nil {
		return nil, err
	}
	var (
		r struct {
			Data  model.SupplierOrderList `json:"data"`
			Error string                  `json:"error"`
		}
	)
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}
	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return &r.Data, nil
}
