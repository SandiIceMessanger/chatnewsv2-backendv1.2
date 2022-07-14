package models

import "gopkg.in/mgo.v2/bson"

type (
	Category struct {
		Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name       string        `json:"name" bson:"name"`
		TimeStamps []TimeStamps
		Age        int `json:"age" bson:"age"`
	}
)
