package main

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
	Board core.Board		`json:"board"`
	Session core.Session		`json:"session"`
}
