package structures

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct{
	SessionID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID bson.ObjectId `json:"user,omitempty" bson:"user,omitempty"`
	Expires time.Time `json:"expires,omitempty" bson:"expires,omitempty"`
}

type SessionRequest struct{
	Session Session `json:"session"`
}

