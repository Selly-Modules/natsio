package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// GetSellerByIDRequest ...
type GetSellerByIDRequest struct {
	SellerID primitive.ObjectID `json:"sellerId"`
}

// GetListSellerByIDsRequest ...
type GetListSellerByIDsRequest struct {
	SellerIDs []primitive.ObjectID `json:"sellerIds"`
}

// GetListSellerSupportChatByIDsRequest ...
type GetListSellerSupportChatByIDsRequest struct {
	SellerIDs []primitive.ObjectID `json:"sellerIds"`
}
