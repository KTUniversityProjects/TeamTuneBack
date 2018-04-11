package main

import (
	"../../core"
	"../../core/structures"
)

var servicePort = "1337"

func main() {
	//core.AddRouting("DELETE", deleteHandler)
	core.AddRouting("PUT", createHandler)
	core.AddRouting("POST", getHandler)
	core.Initialize(servicePort)
}

func getHandler() {

	//Parses request data to
	var data structures.BoardListRequest
	core.DecodeRequest(&data)

	//Session check
	UserID := core.Dao.CheckSession(data.Session)

	//Checks user
	CheckUser(data.ProjectID, UserID)

	//Gets project list
	getList(data.ProjectID)

}

func createHandler() {

	//Parses request data to
	var data structures.BoardCreation
	core.DecodeRequest(&data)

	//Gets user
	user := core.Dao.CheckSession(data.Session)

	//gets project
	project := getProject(data.Board, user)

	//validates
	validate(data.Board, project)

	//Adds project to database
	addBoard(data.Board)
}
