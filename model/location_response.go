package model

type ResponseLocationAddress struct {
	Province LocationProvince `json:"province"`
	District LocationDistrict `json:"district"`
	Ward     LocationWard     `json:"ward"`
}

// LocationProvince ...
type LocationProvince struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code int    `json:"code"`
}

// LocationDistrict ...
type LocationDistrict struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code int    `json:"code"`
}

// LocationWard ...
type LocationWard struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code int    `json:"code"`
}
