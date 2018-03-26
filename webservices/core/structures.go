package core

import "gopkg.in/mgo.v2/bson"

type Session struct{
	SessionID bson.ObjectId       `json:"session,omitempty" bson:"_id,omitempty"`
	UserID string       `json:"omitempty" bson:"user,omitempty"`
}