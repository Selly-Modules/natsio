package subject

import "fmt"

func getLocationValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Location, val)
}

var Location = struct {
	GetLocationByCode string
}{
	GetLocationByCode: getLocationValue("get_location_warehouse"),
}
