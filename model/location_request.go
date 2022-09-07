package model

// LocationRequestPayload ...
type LocationRequestPayload struct {
	Province int `json:"province"`
	District int `json:"district"`
	Ward     int `json:"ward"`
}

// ProvinceRequestPayload ...
type ProvinceRequestPayload struct {
	Codes []int `json:"codes"`
}

// DistrictRequestPayload ...
type DistrictRequestPayload struct {
	Codes []int `json:"codes"`
}

// WardRequestPayload ...
type WardRequestPayload struct {
	Codes []int `json:"codes"`
}
