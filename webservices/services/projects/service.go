package main

import (
	"core"
	"fmt"
)

func main() {
	core.AddRouting("DELETE", deleteHandler)
	core.AddRouting("PUT", createHandler)
	core.AddRouting("POST", getHandler)
	core.AddRouting("PATCH", patchHandler)
	core.Initialize()
}

func createHandler() {
	//Parses request data to
	var data ProjectRequest
	core.DecodeRequest(&data)
	fmt.Println(data)
	//Gets user
	user := core.Dao.CheckSession(data.Session)

	//sets user as creator
	data.Project.Users = []core.ProjectUser{
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
	var data ProjectRequest
	core.DecodeRequest(&data)

	//Gets userID
	UserID := core.Dao.CheckSession(data.Session)

	//Gets all projects
	getList(UserID, data.Project)
}

func deleteHandler() {

	//Parses request data to
	var data ProjectRequest
	core.DecodeRequest(&data)

	//Gets user
	user := core.Dao.CheckSession(data.Session)
	data.Project.Users = []core.ProjectUser{
		{
			ID:user,
			Creator:true,
		},
	}

	//validates
	checkUser(data.Project.ID, user)

	//Remove Boards
	removeBoards(data.Project.ID)

	//Remove Project
	removeProject(data.Project.ID)
}

func patchHandler() {
	//Parses request data to
	var data ProjectEditRequest
	core.DecodeRequest(&data)

	//Gets user
	user := core.Dao.CheckSession(data.Session)

	//checkFieldsExistanceEdit(data.Project, user) //nieko nebevykdo po validate

	//validates
	checkUser(data.Project.ID, user)

	editProject(data.Project)
}