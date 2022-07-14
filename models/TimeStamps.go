package models

import "gopkg.in/mgo.v2/bson"

type (
	TimeStamps struct {
		Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		CreatedBy string        `json:"created_by" bson:"created_by"`
		CreatedAt string        `json:"created_at" bson:"created_at"`
		UpdatedBy string        `json:"update_by" bson:"update_by"`
		DeletedBy string        `json:"deleted_by" bson:"deleted_by"`
	}
)
