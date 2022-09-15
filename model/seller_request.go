package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// GetSellerByIDRequest ...
type GetSellerByIDRequest struct {
	SellerID primitive.ObjectID `json:"sellerId"`
}
