package main

import (
"../../core"
"../../core/structures"
	_"fmt"
)

var servicePort = "1341"

func main() {
	core.AddRouting("PUT", addHandler)
	core.AddRouting("POST", getHandler)
	//core.AddRouting("DELETE", deleteHandler)
	core.Initialize(servicePort)
}

func deleteHandler(){


}

func getHandler() {

	//Parses request data to
	var data structures.TaskListRequest
	core.DecodeRequest(&data)
	//Gets all projects
	getList(data.BoardID)
}

func addHandler() {

	var data structures.TaskCreationRequest
	//Parses request data
	core.DecodeRequest(&data)

	//Gets user
	user := core.Dao.CheckSession(data.Session)

	//check if user able to edit board
	checkUser(data.Board.ID, user)

	validate(data.Task)
	//Adds user to database
	data.Task.BoardID = data.Board.ID
	addTask(data)
}
