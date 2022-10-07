package client

import (
	"encoding/json"
	"errors"

	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
)

type Staff struct{}

func (s Staff) GetListStaffInfoByIDs(p model.GetListStaffInfoByIDsRequest) (*model.ResponseListStaffInfo, error) {
	msg, err := natsio.GetServer().Request(subject.Staff.GetListStaffInfoByIDs, toBytes(p))

	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.ResponseListStaffInfo `json:"data"`
		Error string                       `json:"error"`
	}

	if err := json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}

	return r.Data, nil
}
