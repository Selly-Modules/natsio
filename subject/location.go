package subject

import "fmt"

func getLocationValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Location, val)
}

var Location = struct {
	GetLocationByCode         string
	GetProvincesByCodes       string
	GetDistrictsByCodes       string
	GetWardsByCodes           string
	GetProvinceByCondition    string
	GetProvincesByCondition   string
	GetDistrictByCondition    string
	GetDistrictsByCondition   string
	GetWardByCondition        string
	GetWardsByCondition       string
	CountProvinceByCondition  string
	CountDistrictByCondition  string
	CountWardByCondition      string
	ProvinceDistinctWithField string
}{
	GetLocationByCode:         getLocationValue("get_location_warehouse"),
	GetProvincesByCodes:       getLocationValue("get_provinces_by_codes"),
	GetDistrictsByCodes:       getLocationValue("get_districts_by_codes"),
	GetWardsByCodes:           getLocationValue("get_wards_by_codes"),
	GetProvinceByCondition:    getLocationValue("get_province_by_condition"),
	GetProvincesByCondition:   getLocationValue("get_provinces_by_condition"),
	GetDistrictByCondition:    getLocationValue("get_district_by_condition"),
	GetDistrictsByCondition:   getLocationValue("get_districts_byCondition"),
	GetWardByCondition:        getLocationValue("get_ward_by_condition"),
	GetWardsByCondition:       getLocationValue("get_wards_by_condition"),
	CountProvinceByCondition:  getLocationValue("count_province_by_condition"),
	CountDistrictByCondition:  getLocationValue("count_district_by_condition"),
	CountWardByCondition:      getLocationValue("count_ward_by_condition"),
	ProvinceDistinctWithField: getLocationValue("province_distinct_with_field"),
}
