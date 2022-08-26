package subject

import "fmt"

func getWarehouseValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Warehouse, val)
}

var Warehouse = struct {
	CreateOutboundRequest         string
	UpdateOutboundRequestLogistic string
	CancelOutboundRequest         string
	GetConfiguration              string
}{
	CreateOutboundRequest:         getWarehouseValue("create_outbound_request"),
	UpdateOutboundRequestLogistic: getWarehouseValue("update_outbound_request_logistic_info"),
	CancelOutboundRequest:         getWarehouseValue("cancel_outbound_request"),
	GetConfiguration:              getWarehouseValue("get_configuration"),
}
