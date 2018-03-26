package structures

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct{
	SessionID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID bson.ObjectId `json:"user" bson:"user,omitempty"`
	Expires time.Time `json:"expires" bson:"expires,omitempty"`
}

