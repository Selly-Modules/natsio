package model

import "go.mongodb.org/mongo-driver/mongo/options"

type FindWithCondition struct {
	Conditions interface{}            `json:"conditions"`
	Opts       []*options.FindOptions `json:"opts"`
}

type FindOneCondition struct {
	Conditions interface{} `json:"conditions"`
}

type DistinctWithField struct {
	Conditions interface{} `json:"conditions"`
	Filed      string      `json:"filed"`
}
