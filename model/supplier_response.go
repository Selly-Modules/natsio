package model

// ResponseSupplierInfo ...
type ResponseSupplierInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	BusinessType string `json:"businessType"`
}

// ResponseSupplierContract ...
type ResponseSupplierContract struct {
	ID       string `json:"id"`
	Supplier string `json:"supplier"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}

// SupplierBrief ...
type SupplierBrief struct {
	ID           string `json:"_id"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	BusinessType string `json:"businessType"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type SupplierAll struct {
	Suppliers []SupplierBrief `json:"suppliers"`
	Total     int64           `json:"total"`
}
