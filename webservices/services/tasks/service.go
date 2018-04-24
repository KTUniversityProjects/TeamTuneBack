package main

import (
"../../core"
"../../core/structures"
	_"fmt"
	"fmt"
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
	fmt.Println("Testas0")
	//Gets user
	user := core.Dao.CheckSession(data.Session)
    fmt.Println("Testas1")
	//gets project
	checkProject(data, user)
	fmt.Println("Testas2")
	validate(data.Task)
	//Adds user to database
	fmt.Println("Testas3")
	addTask(data)
}
