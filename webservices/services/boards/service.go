package main

import (
	"../../core"
	"../../core/structures"
)

var servicePort = "1337"

func main() {
	core.AddRouting("DELETE", deleteHandler)
	core.AddRouting("PUT", createHandler)
	core.AddRouting("POST", getHandler)
	core.Initialize(servicePort)
}

func deleteHandler() {

	//Parses request data to
	var data structures.BoardRequest
	core.DecodeRequest(&data)

	//Gets user
	user := core.Dao.CheckSession(data.Session)

	//validates
	checkUser(data.Board.ID, user)

	//Remove Boards
	removeTasks(data.Board.ID)

	//Remove Project
	removeBoard(data.Board.ID)
}

func getHandler() {

	//Parses request data to
	var data structures.BoardListRequest
	core.DecodeRequest(&data)

	//Session check
	UserID := core.Dao.CheckSession(data.Session)

	//Checks user
	CheckUser(data.Project.ID, UserID)

	//Gets project list
	getList(data.Project.ID)

}

func createHandler() {

	//Parses request data to
	var data structures.BoardRequest
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
