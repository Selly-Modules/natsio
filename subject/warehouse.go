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
	SyncORStatus                  string
	WebhookTNC                    string
	WebhookGlobalCare             string
	WebhookOnPoint                string
	FindOne                       string
	FindByCondition               string
	Distinct                      string
	Count                         string
	AfterUpdateWarehouse          string
	AfterCreateWarehouse          string
}{
	AfterCreateWarehouse:          getWarehouseValue("after_create_warehouse"),
	AfterUpdateWarehouse:          getWarehouseValue("after_update_warehouse"),
	CreateOutboundRequest:         getWarehouseValue("create_outbound_request"),
	UpdateOutboundRequestLogistic: getWarehouseValue("update_outbound_request_logistic_info"),
	CancelOutboundRequest:         getWarehouseValue("cancel_outbound_request"),
	GetConfiguration:              getWarehouseValue("get_configuration"),
	SyncORStatus:                  getWarehouseValue("sync_or_status"),
	WebhookTNC:                    getWarehouseValue("webhook_tnc"),
	WebhookGlobalCare:             getWarehouseValue("webhook_global_care"),
	WebhookOnPoint:                getWarehouseValue("webhook_on_point"),
	FindOne:                       getWarehouseValue("find_one"),
	FindByCondition:               getWarehouseValue("find_all_by_condition"),
	Distinct:                      getWarehouseValue("distinct"),
	Count:                         getWarehouseValue("count"),
}
