package client

import (
	"encoding/json"
	"errors"

	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
)

// BankBranch ...
type BankBranch struct{}

// GetBankBranch ...
func GetBankBranch() BankBranch {
	return BankBranch{}
}

func (s Bank) GetBankBranchById(bankBranchID string) (*model.BankBranchBrief, error) {
	msg, err := natsio.GetServer().Request(subject.Bank.GetBankBranchById, toBytes(bankBranchID))
	if err != nil {
		return nil, err
	}

	var r struct {
		Data  *model.BankBranchBrief `json:"data"`
		Error string                 `json:"error"`
	}

	if err = json.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}
	if r.Error != "" {
		return nil, errors.New(r.Error)
	}

	return r.Data, nil
}
