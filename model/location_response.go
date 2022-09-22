package model

import "time"

type (
	// ResponseLocationAddress ...
	ResponseLocationAddress struct {
		Province LocationProvince `json:"province"`
		District LocationDistrict `json:"district"`
		Ward     LocationWard     `json:"ward"`
	}

	// LocationProvince ...
	LocationProvince struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Code           int    `json:"code"`
		RegionCode     string `json:"regionCode"`
		MainRegionCode string `json:"mainRegionCode"`
	}

	// LocationDistrict ...
	LocationDistrict struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Code int    `json:"code"`
	}

	// LocationWard ...
	LocationWard struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Code int    `json:"code"`
	}

	// LocationProvinceResponse ...
	LocationProvinceResponse struct {
		Provinces []LocationProvince `json:"provinces"`
	}

	// LocationDistrictResponse ...
	LocationDistrictResponse struct {
		Districts []LocationDistrict `json:"districts"`
	}

	// LocationWardResponse ...
	LocationWardResponse struct {
		Wards []LocationWard `json:"wards"`
	}

	// LocationProvinceDetailResponse ...
	LocationProvinceDetailResponse struct {
		ID             string    `json:"_id"`
		Name           string    `json:"name"`
		SearchString   string    `json:"searchString"`
		Slug           string    `json:"slug"`
		Code           int       `json:"code"`
		CountryCode    string    `json:"countryCode"`
		RegionCode     string    `json:"regionCode"`
		MainRegionCode string    `json:"mainRegionCode"`
		TotalDistricts int       `json:"totalDistricts"`
		TotalWards     int       `json:"totalWards"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
	}

	// LocationDistrictDetailResponse ...
	LocationDistrictDetailResponse struct {
		ID           string    `json:"_id"`
		Name         string    `json:"name"`
		SearchString string    `json:"searchString"`
		Slug         string    `json:"slug"`
		Code         int       `json:"code"`
		ProvinceCode int       `json:"provinceCode"`
		Area         int       `json:"area"`
		TotalWards   int       `json:"totalWards"`
		CreatedAt    time.Time `json:"createdAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
	}

	// LocationWardDetailResponse ...
	LocationWardDetailResponse struct {
		ID           string    `json:"_id"`
		Name         string    `json:"name"`
		SearchString string    `json:"searchString"`
		Slug         string    `json:"slug"`
		Code         int       `json:"code"`
		DistrictCode int       `json:"districtCode"`
		ProvinceCode int       `json:"provinceCode"`
		CreatedAt    time.Time `json:"createdAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
	}
)
