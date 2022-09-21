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

// GetListSellerInfoByIDs ...
func (s Seller) GetListSellerInfoByIDs(p model.GetListSellerByIDsRequest) (*model.ResponseListSellerInfo, error) {
	msg, err := natsio.GetServer().Request(subject.Seller.GetListSellerInfoByIDs, toBytes(p))

	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.ResponseListSellerInfo `json:"data"`
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

// GetListSellerInfoSupportChatByIDs ...
func (s Seller) GetListSellerInfoSupportChatByIDs(p model.GetListSellerSupportChatByIDsRequest) (*model.ResponseListSellerInfoSupportChat, error) {
	msg, err := natsio.GetServer().Request(subject.Seller.GetListSellerInfoSupportChatByIDs, toBytes(p))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.ResponseListSellerInfoSupportChat `json:"data"`
		Error string                                   `json:"error"`
	}

	if err := json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)

	}
	return r.Data, nil
}
