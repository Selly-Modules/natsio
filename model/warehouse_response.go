package model

import "time"

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
	Other                   WarehouseOther    `json:"other"`
}

// WarehouseOther ...
type WarehouseOther struct {
	DoesSupportSellyExpress bool `json:"doesSupportSellyExpress"`
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
	PriorityServiceCodes []string `json:"priorityServiceCodes"`
	EnabledSources       []int    `json:"enabledSources"`
	Types                []string `json:"types"`
}

// WarehousePartner ...
type WarehousePartner struct {
	IdentityCode   string `json:"identityCode"`
	Code           string `json:"code"`
	Enabled        bool   `json:"enabled"`
	Authentication string `json:"authentication"`
}

// SyncORStatusResponse ...
type SyncORStatusResponse struct {
	ORCode         string `json:"orCode"`
	OrderCode      string `json:"orderCode"`
	Status         string `json:"status"`
	DeliveryStatus string `json:"deliveryStatus"`
}

// ResponseWarehouseContact ...
type ResponseWarehouseContact struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

// ResponseWarehouseLocation ...
type ResponseWarehouseLocation struct {
	Province            CommonLocation `json:"province"`
	District            CommonLocation `json:"district"`
	Ward                CommonLocation `json:"ward"`
	Address             string         `json:"address"`
	LocationCoordinates ResponseLatLng `json:"locationCoordinates"`
}

type CommonLocation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code int    `json:"code"`
}

// ResponseLatLng ...
type ResponseLatLng struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// WarehouseNatsResponse ...
type WarehouseNatsResponse struct {
	ID             string                    `json:"_id"`
	Name           string                    `json:"name"`
	SearchString   string                    `json:"searchString"`
	Slug           string                    `json:"slug"`
	Status         string                    `json:"status"`
	Supplier       string                    `json:"supplier"`
	Contact        ResponseWarehouseContact  `json:"contact"`
	Location       ResponseWarehouseLocation `json:"location"`
	Configurations WarehouseConfiguration    `json:"configurations"`
	CreatedAt      time.Time                 `json:"createdAt"`
	UpdatedAt      time.Time                 `json:"updatedAt"`
}
