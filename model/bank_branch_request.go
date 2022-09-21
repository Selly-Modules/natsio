package model

type BankBranchRequest struct {
	BankID   string `json:"bankId"`
	BranchID string `json:"branchId"`
}
