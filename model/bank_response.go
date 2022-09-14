package model

// MultiLang ...
type MultiLang struct {
	En string `json:"en"`
	Vi string `json:"vi"`
}

// BankBrief ...
type BankBrief struct {
	ID                       string    `json:"_id"`
	Name                     MultiLang `json:"name"`
	ShortName                string    `json:"shortName"`
	Active                   bool      `json:"active"`
	CreatedAt                string    `json:"createdAt"`
	UpdatedAt                string    `json:"updatedAt"`
	BenBankName              string    `json:"benBankName"`
	BankCode                 int       `json:"bankCode"`
	IsBranchRequired         bool      `json:"isBranchRequired"`
	SearchString             string    `json:"searchString"`
	BeneficiaryForVietinbank string    `json:"beneficiaryForVietinbank"`
	CreatedBy                string    `json:"createdBy,omitempty"`
}
