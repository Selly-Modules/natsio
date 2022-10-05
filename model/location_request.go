package model

// LocationRequestPayload ...
type (
	LocationRequestPayload struct {
		Province int `json:"province"`
		District int `json:"district"`
		Ward     int `json:"ward"`
	}

	// ProvinceRequestPayload ...
	ProvinceRequestPayload struct {
		Codes []int `json:"codes"`
	}

	// ProvinceRequestCondition ...
	ProvinceRequestCondition struct {
		Code  int   `json:"code"`
		Codes []int `json:"codes"`
	}

	// DistrictRequestPayload ...
	DistrictRequestPayload struct {
		Codes []int `json:"codes"`
	}

	// DistrictRequestCondition ...
	DistrictRequestCondition struct {
		Code         int `json:"code"`
		Codes        int `json:"codes"`
		ProvinceCode int `json:"provinceCode"`
	}

	// WardRequestPayload ...
	WardRequestPayload struct {
		Codes []int `json:"codes"`
	}

	// WardRequestCondition ...
	WardRequestCondition struct {
		Code         int   `json:"code"`
		Codes        []int `json:"codes"`
		DistrictCode int   `json:"districtCode"`
		ProvinceCode int   `json:"provinceCode"`
	}
)
