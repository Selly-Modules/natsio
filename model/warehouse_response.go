package model

// OutboundRequestResponse ...
type OutboundRequestResponse struct {
	// System code
	OrderCode    string `json:"orderCode"`
	TrackingCode string `json:"trackingCode"`
	ID           string `json:"id"` // OR id

	// Partner response
	ORCode    string `json:"orCode"`
	RequestID string `json:"requestId"`
	Status    string `json:"status"`
	Reason    string `json:"reason"`
}

// WarehouseConfiguration ...
type WarehouseConfiguration struct {
	Warehouse               string            `json:"warehouse"`
	DoesSupportSellyExpress bool              `json:"doesSupportSellyExpress"`
	Supplier                WarehouseSupplier `json:"supplier"`
	Order                   WarehouseOrder    `json:"order"`
	Partner                 WarehousePartner  `json:"partner"`
	Delivery                WarehouseDelivery `json:"delivery"`
}

// WarehouseSupplier ...
type WarehouseSupplier struct {
	CanAutoSendMail       bool   `json:"canAutoSendMail"`
	InvoiceDeliveryMethod string `json:"invoiceDeliveryMethod"`
}

// WarehouseOrder ...
type WarehouseOrder struct {
	MinimumValue             float64                `json:"minimumValue"`
	PaymentMethod            WarehousePaymentMethod `json:"paymentMethod"`
	IsLimitNumberOfPurchases bool                   `json:"isLimitNumberOfPurchases"`
	LimitNumberOfPurchases   int64                  `json:"limitNumberOfPurchases"`
}

// WarehousePaymentMethod ...
type WarehousePaymentMethod struct {
	Cod          bool `json:"cod"`
	BankTransfer bool `json:"bankTransfer"`
}

// WarehouseDelivery ...
type WarehouseDelivery struct {
	DeliveryMethods      []string `json:"deliveryMethods"`
	PriorityServiceCodes []string `json:"priorityDeliveryServiceCodes"`
	EnabledSources       []int    `json:"enabledDeliverySources"`
	Types                []string `json:"type"`
}

// WarehousePartner ...
type WarehousePartner struct {
	IdentityCode   string `json:"identityCode"`
	Code           string `json:"code"`
	Enabled        bool   `json:"enabled"`
	Authentication string `json:"authentication"`
}
