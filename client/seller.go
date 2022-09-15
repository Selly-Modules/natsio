package client

import (
	"encoding/json"
	"errors"
	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
)

// Seller ...
type Seller struct{}

// GetSeller ...
func GetSeller() Seller {
	return Seller{}
}

// GetSellerInfoByID ...
func (s Seller) GetSellerInfoByID(p model.GetSellerByIDRequest) (*model.ResponseSellerInfo, error) {
	msg, err := natsio.GetServer().Request(subject.Seller.GetSellerInfoByID, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.ResponseSellerInfo `json:"data"`
		Error string                    `json:"error"`
	}

	if err := json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}

	return r.Data, nil
}
