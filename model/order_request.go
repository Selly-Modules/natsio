package model

// OrderUpdateORStatus ...
type OrderUpdateORStatus struct {
	OrderCode string `json:"orderCode"`
	ORCode    string `json:"orCode"`
	Status    string `json:"status"`
	Reason    string `json:"reason"`
}
