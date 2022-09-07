package subject

import "fmt"

func getLocationValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Location, val)
}

var Location = struct {
	GetLocationByCode   string
	GetProvincesByCodes string
	GetDistrictsByCodes string
	GetWardsByCodes     string
}{
	GetLocationByCode:   getLocationValue("get_location_warehouse"),
	GetProvincesByCodes: getLocationValue("get_provinces_by_codes"),
	GetDistrictsByCodes: getLocationValue("get_districts_by_codes"),
	GetWardsByCodes:     getLocationValue("get_wards_by_codes"),
}
