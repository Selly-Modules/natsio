package subject

const warehousePrefix = "warehouse_"

const (
	WarehouseCreateOutboundRequest         = warehousePrefix + "create_outbound_request"
	WarehouseUpdateOutboundRequestLogistic = warehousePrefix + "update_outbound_request_logistic_info"
	WarehouseCancelOutboundRequest         = warehousePrefix + "cancel_outbound_request"
	WarehouseGetConfiguration              = warehousePrefix + "get_configuration"
)
