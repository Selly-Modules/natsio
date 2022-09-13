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
