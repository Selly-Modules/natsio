package model

// ResponseSupplierInfo ...
type ResponseSupplierInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ResponseSupplierContract ...
type ResponseSupplierContract struct {
	ID       string `json:"id"`
	Supplier string `json:"supplier"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}
