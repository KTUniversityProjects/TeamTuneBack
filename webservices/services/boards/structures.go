package boards

import (
	"../../core/structures"
	"gopkg.in/mgo.v2/bson"
)

type Board struct {
	Id bson.ObjectId        `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`
	Description string      `json:"description,omitempty" bson:"description,omitempty"`
	ProjectID bson.ObjectId `json:"project,omitempty" bson:"project"`
}

type BoardCreation struct{
	Board Board    `json:"board,omitempty"`
	Session structures.Session       `json:"session,omitempty"`
}