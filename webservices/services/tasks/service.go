package main

import (
	"core"
	_"fmt"
)

func main() {
	core.AddRouting("PUT", addHandler)
	core.AddRouting("POST", getHandler)
	core.AddRouting("DELETE", deleteHandler)
	core.Initialize()
}

func deleteHandler(){
	var data TaskRequest
	//Parses request data
	core.DecodeRequest(&data)

	//Gets user
	user := core.Dao.CheckSession(data.Session)

	//check if user able to delete task
	checkUser(data.Task.ID, user)

	deleteTask(data.Task.ID)
}

func getHandler() {
	//Parses request data to
	var data TaskListRequest
	core.DecodeRequest(&data)
	//Gets all projects
	getList(data.BoardID)
}

func addHandler() {
	var data TaskCreationRequest
	//Parses request data
	core.DecodeRequest(&data)

	//Sets name for task
	data.Task.Name = defaultName

	//Gets user
	user := core.Dao.CheckSession(data.Session)

	//check if user able to edit board
	checkUser(data.Board.ID, user)

	validate(data.Task)
	//Adds user to database
	data.Task.BoardID = data.Board.ID
	addTask(data)
}
