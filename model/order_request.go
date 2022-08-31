package model

// OrderUpdateORStatus ...
type OrderUpdateORStatus struct {
	OrderCode      string `json:"orderCode"`
	ORCode         string `json:"orCode"`
	Status         string `json:"status"`
	DeliveryStatus string `json:"deliveryStatus"`
	Reason         string `json:"reason"`
}

// OrderCancelDelivery ...
type OrderCancelDelivery struct {
	OrderID string `json:"orderId"`
}

// OrderChangeDeliveryStatus ...
type OrderChangeDeliveryStatus struct {
	OrderID        string   `json:"orderId"`
	DeliveryStatus string   `json:"deliveryStatus"`
	ActionBy       ActionBy `json:"actionBy"`
}
