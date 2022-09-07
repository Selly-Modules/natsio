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

// LocationProvinceResponse ...
type LocationProvinceResponse struct {
	Province []LocationProvince `json:"province"`
}

// LocationDistrictResponse ...
type LocationDistrictResponse struct {
	District []LocationDistrict `json:"district"`
}

// LocationWardResponse ...
type LocationWardResponse struct {
	Ward []LocationWard `json:"ward"`
}
