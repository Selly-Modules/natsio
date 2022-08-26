package model

// OutboundRequestPayload ...
type OutboundRequestPayload struct {
	OrderID      string                `json:"orderId"`
	OrderCode    string                `json:"orderCode"`
	TrackingCode string                `json:"trackingCode"`
	WarehouseID  string                `json:"warehouseId"`
	SupplierID   string                `json:"supplierId"`
	Note         string                `json:"note"`
	CODAmount    float64               `json:"codAmount"`
	TPLCode      string                `json:"tplCode"`
	Customer     CustomerInfo          `json:"customer"`
	Items        []OutboundRequestItem `json:"items"`
	Insurance    *InsuranceOpts
}

// InsuranceOpts ...
type InsuranceOpts struct {
	VehicleTypeID    string `json:"vehicleTypeId"`
	VehicleTypeName  string `json:"vehicleTypeName"`
	InsuranceTypeID  string `json:"insuranceTypeId"`
	YearsOfInsurance int    `json:"yearsOfInsurance"`
	License          string `json:"license"`
	Chassis          string `json:"chassis"`
	Engine           string `json:"engine"`
	BeginDate        string `json:"beginDate"`
}

// OutboundRequestItem ...
type OutboundRequestItem struct {
	SupplierSKU string `json:"supplierSKU"`
	Quantity    int64  `json:"quantity"`
	UnitCode    string `json:"unitCode"`
}

// CustomerInfo ...
type CustomerInfo struct {
	Name        string        `json:"name"`
	PhoneNumber string        `json:"phoneNumber"`
	Email       string        `json:"email"`
	Address     AddressDetail `json:"address"`
}

// AddressDetail ...
type AddressDetail struct {
	Address      string `json:"address"`
	FullAddress  string `json:"fullAddress"`
	ProvinceCode int    `json:"provinceCode"`
	DistrictCode int    `json:"districtCode"`
	WardCode     int    `json:"wardCode"`
}

// UpdateOutboundRequestLogisticInfoPayload ...
type UpdateOutboundRequestLogisticInfoPayload struct {
	ShippingLabel string `json:"shippingLabel"`
	TrackingCode  string `json:"trackingCode"`
	ORCode        string `json:"orCode"`
}

// CancelOutboundRequest ...
type CancelOutboundRequest struct {
	ORCode string `json:"orCode"`
	Note   string `json:"note"`
}
