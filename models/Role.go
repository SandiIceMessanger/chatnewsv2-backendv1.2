package models

import "gopkg.in/mgo.v2/bson"

type (
	Role struct {
		Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name       string        `json:"name" bson:"name"`
		TimeStamps []TimeStamps
		Status     string `json:"status" bson:"status"`
	}
)
