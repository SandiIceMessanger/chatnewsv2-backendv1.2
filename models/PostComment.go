package models

import "gopkg.in/mgo.v2/bson"

type (
	PostComment struct {
		Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		PostId     string        `json:"post_id" bson:"post_id"`
		Comment    string        `json:"comment" bson:"comment"`
		ParentId   string        `json:"parent_id" bson:"comment"`
		TimeStamps []TimeStamps
		Status     string `json:"status" bson:"status"`
	}
)
