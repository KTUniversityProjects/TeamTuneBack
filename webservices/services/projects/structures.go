package projects

import (
	"../../core/structures"
	"gopkg.in/mgo.v2/bson"
)

type Project struct {
	Id bson.ObjectId   `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	User bson.ObjectId `json:"user,omitempty" bson:"user,omitempty"`
}

type ProjectCreation struct{
	Project Project    `json:"project,omitempty"`
	Session structures.Session       `json:"session,omitempty"`
}