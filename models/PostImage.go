package models

import "gopkg.in/mgo.v2/bson"

type (
	PostImage struct {
		Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Thumbnail  string        `json:"thumbnail" bson:"thumbnail"`
		Small      string        `json:"small" bson:"small"`
		Medium     string        `json:"medium" bson:"medium"`
		Large      string        `json:"large" bson:"large"`
		TimeStamps []TimeStamps
		Status     string `json:"status" bson:"status"`
	}
)
