package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetSupplierRequest ...
type GetSupplierRequest struct {
	ListID []primitive.ObjectID `json:"listID"`
}

type GetSupplierContractRequest struct {
	SupplierID primitive.ObjectID `json:"supplierID"`
}

// SupplierRequestPayload ...
type SupplierRequestPayload struct {
	Limit          int
	Page           int
	Keyword        string
	Status         string
	PIC            string
	ContractStatus string
}

type CreateSupplierWarehousePayload struct {
	Supplier     string `json:"supplier"`
	Warehouse    string `json:"warehouse"`
	ProvinceCode int    `json:"provinceCode"`
	DistrictCode int    `json:"districtCode"`
	WardCode     int    `json:"wardCode"`
}

type UpdateSupplierWarehousePayload struct {
	Supplier     string `json:"supplier"`
	Warehouse    string `json:"warehouse"`
	ProvinceCode int    `json:"provinceCode"`
	DistrictCode int    `json:"districtCode"`
	WardCode     int    `json:"wardCode"`
}
