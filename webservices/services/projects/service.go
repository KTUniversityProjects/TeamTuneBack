package main

import (
	"../../core"
	"../../core/structures"
)

var servicePort = "1338"

func main() {
	core.AddRouting("DELETE", deleteHandler)
	core.AddRouting("PUT", createHandler)
	core.AddRouting("POST", getHandler)
	core.Initialize(servicePort)
}

func createHandler() {
	//Parses request data to
	var data structures.ProjectRequest
	core.DecodeRequest(&data)

	//Gets user
	user := core.Dao.CheckSession(data.Session)

	//sets user as creator
	data.Project.Users = []structures.ProjectUser{
		{
			ID:      user,
			Creator: true,
		},
	}

	//validates
	validate(data.Project)

	//Adds project to database
	addProject(data.Project)
}

func getHandler() {

	//Parses request data to
	var data structures.ProjectRequest
	core.DecodeRequest(&data)

	//Gets userID
	UserID := core.Dao.CheckSession(data.Session)

	//Gets all projects
	getList(UserID)
}

func deleteHandler() {

	//Parses request data to
	var data structures.ProjectRequest
	core.DecodeRequest(&data)

	//Gets user
	user := core.Dao.CheckSession(data.Session)
	data.Project.Users = []structures.ProjectUser{
		{
			ID:user,
			Creator:true,
		},
	}

	//validates
	checkUser(data.Project)

	//Remove Boards
	removeBoards(data.Project.ID)

	//Remove Project
	removeProject(data.Project.ID)
}