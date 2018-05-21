package core

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Task struct {
	ID bson.ObjectId        `json:"id,omitempty" bson:"_id,omitempty"`
	Name string  			`json:"name,omitempty" bson:"name,omitempty"`
	BoardID bson.ObjectId   `json:"boardID,omitempty" bson:"board,omitempty"`
}

type Project struct {
	ID bson.ObjectId    `json:"id,omitempty" bson:"_id,omitempty"`
	Name string         `json:"name,omitempty" bson:"name,omitempty"`
	Description string  `json:"description,omitempty" bson:"description,omitempty"`
	Users []ProjectUser `json:"users,omitempty" bson:"users,omitempty"`
}

type User struct {
	Id bson.ObjectId        `json:"id,omitempty" bson:"_id,omitempty"`
	Username string  `json:"username,omitempty" bson:"username,omitempty"`
	Password string  `json:"password,omitempty" bson:"password,omitempty"`
	Password2 string `json:"password2,omitempty"`
	Email    string  `json:"email,omitempty" bson:"email,omitempty"`
}


type Board struct {
	ID bson.ObjectId        `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`
	Description string      `json:"description,omitempty" bson:"description,omitempty"`
	ProjectID bson.ObjectId `json:"project,omitempty" bson:"project,"`
	Tasks []Task   `json:"tasks,omitempty" bson:"_tasks,omitempty"`
}



type ProjectUser struct {
	ID bson.ObjectId     `json:"user,omitempty" bson:"_id,omitempty"`
	Permissions []int    `json:"permissions,omitempty" bson:"permissions,omitempty"`
	Creator bool         `json:"creator,omitempty" bson:"creator,omitempty"`
}

type Session struct{
	SessionID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID bson.ObjectId `json:"user,omitempty" bson:"user,omitempty"`
	Expires time.Time `json:"expires,omitempty" bson:"expires,omitempty"`
}

type SessionRequest struct{
	Session Session `json:"session"`
}