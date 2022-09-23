package jssubject

import (
	"fmt"
)

// getSellyValue ...
func getSellyValue(val string) string {
	return fmt.Sprintf("%s.%s.%s", root, prefixes.Selly, val)
}

// Selly ...
var Selly = struct {
	PushNotification      string
	UpdateSellerStatistic string
}{
	PushNotification:      getSellyValue("push_notifications"),
	UpdateSellerStatistic: getSellyValue("update_seller_statistic"),
}
