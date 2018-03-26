package structures

import "gopkg.in/mgo.v2/bson"

type Session struct{
	SessionID bson.ObjectId       `json:"id,omitempty" bson:"_id,omitempty"`
	UserID bson.ObjectId       `json:"omitempty" bson:"user,omitempty"`
}