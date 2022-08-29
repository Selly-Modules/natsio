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
