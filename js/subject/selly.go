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
	PushNotification               string
	UpdateSellerAffiliateStatistic string
}{
	PushNotification:               getSellyValue("push_notifications"),
	UpdateSellerAffiliateStatistic: getSellyValue("update_seller_affiliate_statistic"),
}
