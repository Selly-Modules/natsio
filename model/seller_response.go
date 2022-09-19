package model

// ResponseSellerInfo ...
type ResponseSellerInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// ResponseListSellerInfo ...
type ResponseListSellerInfo struct {
	Sellers []ResponseSellerInfo `json:"sellers"`
}
