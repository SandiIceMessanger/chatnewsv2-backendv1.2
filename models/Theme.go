package models

import "gopkg.in/mgo.v2/bson"

type (
	Theme struct {
		Id               bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name             string        `json:"name" bson:"name"`
		StartDate        string        `json:"start_date" bson:"start_date"`
		EndDate          string        `json:"end_date" bson:"end_date"`
		Tags             string        `json:"tags" bson:"tags"`
		HeaderIcon       string        `json:"header_icon" bson:"header_icon"`
		HeaderIconMobile string        `json:"header_icon_mobile" bson:"header_icon_mobile"`
		ImageLeft        string        `json:"image_left" bson:"image_left"`
		ImageRight       string        `json:"image_right" bson:"image_right"`
		TimeStamps       []TimeStamps
		Status           string `json:"status" bson:"status"`
	}
)
