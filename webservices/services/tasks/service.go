package main

import (
"../../core"
"../../core/structures"
)

var servicePort = "1341"

func main() {
	core.AddRouting("POST", addHandler)
	core.AddRouting("GET", getHandler)
	//core.AddRouting("DELETE", deleteHandler)
	core.Initialize(servicePort)
}

func deleteHandler(){


}

func getHandler() {

	//Parses request data to
	var data structures.Board
	core.DecodeRequest(&data)

	//Gets all projects
	getList(data.ID)
}

func addHandler() {

	var data structures.Task
	//Parses request data
	core.DecodeRequest(&data)

	//Adds user to database
	addTask(data)
}
