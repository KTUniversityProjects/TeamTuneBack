package users

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id bson.ObjectId        `json:"id,omitempty" bson:"_id,omitempty"`
	Username string  `json:"username,omitempty" bson:"username,omitempty"`
	Password string  `json:"password,omitempty" bson:"password,omitempty"`
	Password2 string `json:"password2,omitempty"`
	Email    string  `json:"email,omitempty" bson:"email,omitempty"`
}