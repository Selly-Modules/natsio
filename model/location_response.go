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
}

// LocationDistrict ...
type LocationDistrict struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// LocationWard ...
type LocationWard struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
