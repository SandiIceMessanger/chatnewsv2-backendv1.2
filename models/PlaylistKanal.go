package models

import "gopkg.in/mgo.v2/bson"

type (
	PlaylistKanal struct {
		Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		CategoryId string        `json:"category_id" bson:"category_id"`
		Headlines  string        `json:"headlines" bson:"headlines"`
		Trendings  string        `json:"trendings" bson:"trendings"`
		TimeStamps []TimeStamps
		Status     string `json:"status" bson:"status"`
	}
)
