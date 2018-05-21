package main

import (
	"core"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

//Checks right for deleting
func checkUser(boardID bson.ObjectId, userID bson.ObjectId) {

	core.Dao.C("boards")
	var board core.Board

	err := core.Dao.Collection.Find(bson.M{"_id": boardID}).Select(bson.M{"project": 1}).One(&board)
	if err != nil {
		fmt.Println("getProjectID")
		fmt.Println(err)
		core.ThrowResponse("database_error")
	}

	core.Dao.C("projects")
	count, err := core.Dao.Collection.Find(bson.M{"_id": board.ProjectID, "users": bson.M{"$elemMatch": bson.M{"_id":userID}}}).Count()
	if err != nil {
		fmt.Println("FindUser")
		fmt.Println(err)
		core.ThrowResponse("database_error")
	}
	if count == 0 {
		fmt.Println("CheckProject")
		fmt.Println(err)
		core.ThrowResponse("no_permission")
	}
}

// Remove board from project
func RemoveBoardFromProject(boardID bson.ObjectId) {

	var board core.Board
	fmt.Println(boardID)
	core.Dao.C("boards")
	err := core.Dao.Collection.Find(bson.M{"_id" : boardID}).Select(bson.M{"project":1}).One(&board)
	if err != nil{
		fmt.Println(err)
		core.ThrowResponse("database_error")
	}

	fmt.Println(board.ProjectID)
	core.Dao.C("projects")
	err = core.Dao.Collection.UpdateId(board.ProjectID, bson.M{"$pull": bson.M{"boards": boardID}})
	if err != nil{
		fmt.Println(err)
		core.ThrowResponse("database_error")
	}
}

//Removes Board from Database
func removeBoard(boardID bson.ObjectId) {

	//Removes tasks
	core.Dao.C("tasks")
	_, err := core.Dao.Collection.RemoveAll(bson.M{"board": boardID})
	if err != nil {
		fmt.Println("tasks delete")
		fmt.Println(err)
		core.ThrowResponse("database_error")
	}

	RemoveBoardFromProject(boardID)
	//Removes board
	core.Dao.C("boards")
	err = core.Dao.Collection.Remove(bson.M{"_id": boardID})
	if err != nil {
		fmt.Println("RemoveBoards")
		fmt.Println(err)
		core.ThrowResponse("database_error")
	}
}

//gets board list by ProjectID
func getList(projectID bson.ObjectId) {
	core.Dao.C("boards")

	var results []core.Board
	fmt.Println(projectID)
	err := core.Dao.Collection.Find(bson.M{"project": projectID}).Select(bson.M{"_id": 1, "name": 1}).All(&results)
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.Dao.C("tasks")

	for id, element := range results {
		err := core.Dao.Collection.Find(bson.M{"board": element.ID}).Select(bson.M{"_id": 1, "name": 1}).All(&results[id].Tasks)
		if err != nil {
			fmt.Println(err)
			core.ThrowResponse("database_error")
		}
	}


	core.SetData(results)
}

//Check if user can see project
func CheckUser(projectID bson.ObjectId, userID bson.ObjectId) {
	core.Dao.C("projects")

	count, err := core.Dao.Collection.Find(bson.M{"_id": projectID, "users": bson.M{"$elemMatch": bson.M{"_id": userID}}}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}

	if count == 0 {
		core.ThrowResponse("database_error")
	}
}

//Checks board existance for project
func checkFieldsExistance(board core.Board) {
	core.Dao.C("projects")

	count, err := core.Dao.Collection.Find(bson.M{"name": board.Name, "project": board.ProjectID}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}
	if count > 0 {
		core.ThrowResponse("name_exists")
	}
}

//Validates BOard data
func validate(board core.Board, project core.Project) {

	if board.Name == "" {
		core.ThrowResponse("empty_fields")
	}

	//Checks for permissions
	if !project.Users[0].Creator {
		core.ThrowResponse("no_permission")
	}

	if board.Name == "" {
		core.ThrowResponse("system_mistake")
	}

	checkFieldsExistance(board)
}

//Gets project for board creation
func getProject(board core.Board, userID bson.ObjectId) core.Project {
	core.Dao.C("projects")

	var project = core.Project{}
	err := core.Dao.Collection.Find(bson.M{"_id": board.ProjectID, "users": bson.M{"$elemMatch": bson.M{"_id": userID}}}).One(&project)

	if err != nil || project.ID == "" {
		core.ThrowResponse("no_permission")
	}

	return project
}

//Adds Board to Database
func addBoard(board core.Board) {
	core.Dao.C("boards")

	board.ID = bson.NewObjectId()

	err := core.Dao.Collection.Insert(&board)
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.Dao.C("projects")
	err = core.Dao.Collection.Update(bson.M{"_id": board.ProjectID}, bson.M{"$push": bson.M{"boards": board.ID}})
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.SetData(board.ID)
	core.ThrowResponse("board_created")
}

func editBoard(data core.Board) {
	core.Dao.C("boards")

	if data.Name != "" {
		err := core.Dao.Collection.Update(bson.M{"_id": data.ID}, bson.M{"$set": bson.M{"name": data.Name}})
		if err != nil {
			fmt.Println(err)
			core.ThrowResponse("database_error")
		}
	}
	if data.Description != "" {
		err := core.Dao.Collection.Update(bson.M{"_id": data.ID}, bson.M{"$set": bson.M{"description": data.Description}})
		if err != nil {
			fmt.Println(err)
			core.ThrowResponse("database_error")
		}
	}
}
//Removes Project from Database
func removeTasks(boardID bson.ObjectId) {

	core.Dao.C("tasks")

	_, err := core.Dao.Collection.RemoveAll(bson.M{"board": boardID})
	if err != nil {
		fmt.Println("RemoveTasks")
		fmt.Println(err)
		core.ThrowResponse("database_error")
	}
}
