package models

import "gopkg.in/mgo.v2/bson"

type (
	PostTag struct {
		Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		PostId     string        `json:"post_id" bson:"post_id"`
		Tag        string        `json:"tag" bson:"tag"`
		TimeStamps []TimeStamps
		Status     string `json:"status" bson:"status"`
	}
)
