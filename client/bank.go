package client

import (
	"encoding/json"
	"errors"

	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
)

// Bank ...
type Bank struct{}

// GetBank ...
func GetBank() Bank {
	return Bank{}
}

func (s Bank) GetBankById(bankID string) (*model.BankBrief, error) {
	msg, err := natsio.GetServer().Request(subject.Bank.GetBankById, toBytes(bankID))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.BankBrief `json:"data"`
		Error string           `json:"error"`
	}

	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}
	if r.Error != "" {
		return nil, errors.New(r.Error)
	}

	return r.Data, nil
}
