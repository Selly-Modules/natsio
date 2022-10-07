package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// GetListStaffInfoByIDsRequest ...
type GetListStaffInfoByIDsRequest struct {
	StaffIDs []primitive.ObjectID `json:"staffIds"`
}
