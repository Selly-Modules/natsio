package subject

import "fmt"

func getOrderValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Order, val)
}

var Order = struct {
	UpdateORStatus string
	CancelDelivery string
}{
	UpdateORStatus: getOrderValue("update_outbound_request_status"),
	CancelDelivery: getOrderValue("cancel_delivery"),
}
