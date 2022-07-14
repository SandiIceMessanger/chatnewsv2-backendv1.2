package models

import "gopkg.in/mgo.v2/bson"

type (
	Page struct {
		Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name       string        `json:"name" bson:"name"`
		Logo       string        `json:"logo" bson:"logo"`
		Image      string        `json:"image" bson:"image"`
		Color      string        `json:"color" bson:"color"`
		Content    string        `json:"content" bson:"content"`
		SlugUrl    string        `json:"slug_url" bson:"slug_url"`
		TimeStamps []TimeStamps
		Status     string `json:"status" bson:"status"`
	}
)
