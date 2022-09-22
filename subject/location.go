package subject

import "fmt"

func getLocationValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Location, val)
}

var Location = struct {
	GetLocationByCode          string
	GetProvincesByCodes        string
	GetDistrictsByCodes        string
	GetWardsByCodes            string
	FindOneProvinceByCondition string
	FindOneDistrictByCondition string
	FindOneWardByCondition     string
	FindProvinceByCondition    string
	FindDistrictByCondition    string
	FindWardByCondition        string
	CountProvinceByCondition   string
	CountDistrictByCondition   string
	CountWardByCondition       string
}{
	GetLocationByCode:          getLocationValue("get_location_warehouse"),
	GetProvincesByCodes:        getLocationValue("get_provinces_by_codes"),
	GetDistrictsByCodes:        getLocationValue("get_districts_by_codes"),
	GetWardsByCodes:            getLocationValue("get_wards_by_codes"),
	FindOneProvinceByCondition: getLocationValue("find_one_province_by_condition"),
	FindOneDistrictByCondition: getLocationValue("find_one_district_by_condition"),
	FindOneWardByCondition:     getLocationValue("find_one_ward_by_condition"),
	FindProvinceByCondition:    getLocationValue("find_province_by_condition"),
	FindDistrictByCondition:    getLocationValue("find_district_by_condition"),
	FindWardByCondition:        getLocationValue("find_ward_by_condition"),
	CountProvinceByCondition:   getLocationValue("count_province_by_condition"),
	CountDistrictByCondition:   getLocationValue("count_district_by_condition"),
	CountWardByCondition:       getLocationValue("count_ward_by_condition"),
}
