package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetSupplierRequest ...
type GetSupplierRequest struct {
	ID primitive.ObjectID `json:"_id"`
}
