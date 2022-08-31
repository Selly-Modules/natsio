package model

// SupplierBrief ...
type SupplierBrief struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type SupplierAll struct {
	Suppliers []SupplierBrief `json:"suppliers"`
	Total     int64           `json:"total"`
}
