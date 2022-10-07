package subject

import "fmt"

func getStaffValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Staff, val)
}

// Staff ...
var Staff = struct {
	GetListStaffInfoByIDs string
}{
	GetListStaffInfoByIDs: getStaffValue("get_list_staff_info_by_ids"),
}
