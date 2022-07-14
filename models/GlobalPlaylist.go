package models

import "gopkg.in/mgo.v2/bson"

type (
	GlobalPlaylist struct {
		Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
		SlideShow   string        `json:"slide_show" bson:"slide_show"`
		SubHeader   string        `json:"sub_header" bson:"sub_header"`
		Trendings   string        `json:"trendings" bson:"trendings"`
		StickerNews string        `json:"sticker_news" bson:"sticker_news"`
		TimeStamps  []TimeStamps
		Status      string `json:"status" bson:"status"`
	}
)
