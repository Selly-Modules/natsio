package model

// GetSupplierRequest ...
type GetSupplierRequest struct {
	Limit          int
	Page           int
	Keyword        string
	Status         string
	PIC            string
	ContractStatus string
}
