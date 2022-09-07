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
	GetProvincesByCodes: getLocationValue("get_provinces_warehouse"),
	GetDistrictsByCodes: getLocationValue("get_districts_warehouse"),
	GetWardsByCodes:     getLocationValue("get_wards_warehouse"),
}
