package models

import "gopkg.in/mgo.v2/bson"

type (
	Event struct {
		Id             bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name           string        `json:"name" bson:"name"`
		StartDate      string        `json:"start_date" bson:"start_date"`
		EndDate        string        `json:"end_date" bson:"end_date"`
		Tags           string        `json:"tags" bson:"tags"`
		Endpoint       string        `json:"endpoint" bson:"endpoint"`
		EndpointDetail string        `json:"endpoint_detail" bson:"endpoint_detail"`
		TimeStamps     []TimeStamps
		Status         string `json:"status" bson:"status"`
	}
)
