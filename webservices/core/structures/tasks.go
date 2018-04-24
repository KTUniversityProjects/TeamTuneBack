package structures

import "gopkg.in/mgo.v2/bson"

type Task struct {
	ID bson.ObjectId        `json:"id,omitempty" bson:"_id,omitempty"`
	Name string  			`json:"name,omitempty" bson:"_name,omitempty"`
	BoardID bson.ObjectId   `json:"boardID,omitempty" bson:"_board,omitempty"`
}

type TaskCreationRequest struct {
	Task Task 				`json:"task,omitempty"`
	Session Session 		`json:"session,omitempty"`
	Board Board 			`json:"board,omitempty"`
}

//listo grazinimui
type TaskListRequest struct{
	BoardID bson.ObjectId         `json:"board"`
	Session Session    			  `json:"session"`
}