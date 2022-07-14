package models

import "gopkg.in/mgo.v2/bson"

type (
	User struct {
		Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		RoleId     int           `json:"role_id" bson:"role_id"`
		Username   string        `json:"username" bson:"username"`
		Password   string        `json:"password" bson:"password"`
		Name       string        `json:"name" bson:"name"`
		Email      string        `json:"email" bson:"email"`
		TimeStamps []TimeStamps
		Status     string `json:"status" bson:"status"`
	}
)
