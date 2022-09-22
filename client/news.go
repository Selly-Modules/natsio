package client

import (
	"encoding/json"
	"errors"
	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
)

// News ...
type News struct{}

// GetNews ...
func GetNews() News {
	return News{}
}

// GetProductNoticesByInventory ...
func (n News) GetProductNoticesByInventory(p model.GetProductNoticesByInventoryRequest) (*model.GetProductNoticesByInventoryResponse, error) {
	msg, err := natsio.GetServer().Request(subject.News.GetProductNoticesByInventory, toBytes(p))
	if err != nil {
		return nil, err
	}
	var r struct {
		Data  *model.GetProductNoticesByInventoryResponse `json:"data"`
		Error string                                      `json:"error"`
	}
	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}
	if r.Error != "" {
		return nil, errors.New(r.Error)
	}
	return r.Data, nil
}
