package structures

import "gopkg.in/mgo.v2/bson"

type Task struct {
	ID bson.ObjectId        `json:"id,omitempty" bson:"_id,omitempty"`
	Name string  			`json:"name,omitempty" bson:"_name,omitempty"`
	BoardID bson.ObjectId   `json:"board,omitempty" bson:"_board,omitempty"`
}