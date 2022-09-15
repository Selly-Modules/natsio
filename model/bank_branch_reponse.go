package model

// BankBranchBrief ...
type BankBranchBrief struct {
	ID       string `json:"_id"`
	City     string `json:"city"`
	BankCode string `json:"bankCode"`
	Bank     string `json:"bank"`
	Active   bool   `json:"active"`
	Name     string `json:"name"`
}
