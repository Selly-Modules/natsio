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
		Code    int      `json:"code"`
		Codes   []int    `json:"codes"`
		Slug    string   `json:"slug"`
		Slugs   []string `json:"slugs"`
		Keyword string   `json:"keyword"`
		Region  string   `json:"region"`
	}

	// DistrictRequestPayload ...
	DistrictRequestPayload struct {
		Codes []int `json:"codes"`
	}

	// DistrictRequestCondition ...
	DistrictRequestCondition struct {
		Code         int    `json:"code"`
		Codes        []int  `json:"codes"`
		ProvinceCode int    `json:"provinceCode"`
		Slug         string `json:"slug"`
		ProvinceSlug string `json:"provinceSlug"`
		Keyword      string `json:"keyword"`
	}

	// WardRequestPayload ...
	WardRequestPayload struct {
		Codes []int `json:"codes"`
	}

	// WardRequestCondition ...
	WardRequestCondition struct {
		Code         int    `json:"code"`
		Codes        []int  `json:"codes"`
		DistrictCode int    `json:"districtCode"`
		ProvinceCode int    `json:"provinceCode"`
		Slug         string `json:"slug"`
		DistrictSlug string `json:"districtSlug"`
		ProvinceSlug string `json:"provinceSlug"`
	}

	// ProvinceDistinctWithField ...
	ProvinceDistinctWithField struct {
		Conditions struct {
			Region string `json:"region"`
		} `json:"conditions"`
		Field string `json:"filed"`
	}
)
