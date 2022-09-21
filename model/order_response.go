package model

import "time"

// SupplierOrderList ...
type SupplierOrderList struct {
	List  []SupplierOrder `json:"list"`
	Total int64           `json:"total" example:"100"`
	Limit int64           `json:"limit" example:"20"`
}

// SupplierOrder ...
type SupplierOrder struct {
	ID        string                `json:"_id"`
	Code      string                `json:"code"`
	CreatedAt time.Time             `json:"createdAt"`
	Status    string                `json:"status"`
	Items     []SupplierOrderItem   `json:"items"`
	Delivery  SupplierOrderDelivery `json:"delivery"`
}

// SupplierOrderItem ...
type SupplierOrderItem struct {
	ID          string `json:"_id" example:"1231"`
	SupplierSKU string `json:"supplierSku" example:"SUPPLIER_SKU"`
	Quantity    int64  `json:"quantity" example:"2"`
}

// SupplierOrderDelivery ...
type SupplierOrderDelivery struct {
	Code    string `json:"code" example:"123187287"`
	Status  string `json:"status" enums:"waiting_to_confirm,waiting_to_pick,picking,picked,delay_pickup,pickup_failed,delivering,delay_delivery,delivered,cancelled,delivery_failed,waiting_to_return,returning,delay_return,compensation,returned"`
	TPLCode string `json:"tplCode" enums:"SLY,GHTK,GHN,SSC,SPY,VTP,SE,NTL,BEST"`
}
