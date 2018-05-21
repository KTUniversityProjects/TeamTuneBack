package main

import "gopkg.in/mgo.v2/bson"
import "core"

type TaskCreationRequest struct {
	Task core.Task 				`json:"task,omitempty"`
	Session core.Session 		`json:"session,omitempty"`
	Board core.Board 			`json:"board,omitempty"`
}

type TaskRequest struct {
	Task core.Task 				`json:"task,omitempty"`
	Session core.Session 		`json:"session,omitempty"`
}

//listo grazinimui
type TaskListRequest struct{
	BoardID bson.ObjectId		`json:"board"`
	Session core.Session		`json:"session"`
}
