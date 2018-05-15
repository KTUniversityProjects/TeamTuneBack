package main

import (
	"gopkg.in/mgo.v2/bson"
	"../../core"
	"../../core/structures"
	"fmt"
)

//Checks right for deleting
func checkUser(boardID bson.ObjectId, userID bson.ObjectId) {

	core.Dao.C("boards")
	var board structures.Board

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
		core.ThrowResponse("project_not_exists")
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

	var results []structures.Board
	fmt.Println(projectID)
	err := core.Dao.Collection.Find(bson.M{"project": projectID}).Select(bson.M{"_id": 1, "name": 1}).All(&results)
	if err != nil {
		core.ThrowResponse("database_error")
	}
	core.Dao.C("tasks")

	for id, element := range results {
		fmt.Println(results[id].ID)
		err := core.Dao.Collection.Find(bson.M{"board": element.ID}).Select(bson.M{"_id": 1, "name": 1}).All(&results[id].Tasks)
		if err != nil {
			fmt.Println(err)
			core.ThrowResponse("database_error")
		}
		fmt.Println("taskai")
		fmt.Println(results[id].Tasks)
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
func checkFieldsExistance(board structures.Board) {
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
func validate(board structures.Board, project structures.Project) {

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
func getProject(board structures.Board, userID bson.ObjectId) structures.Project {
	core.Dao.C("projects")

	var project = structures.Project{}
	err := core.Dao.Collection.Find(bson.M{"_id": board.ProjectID, "users": bson.M{"$elemMatch": bson.M{"_id": userID}}}).One(&project)

	if err != nil || project.ID == "" {
		core.ThrowResponse("project_not_exists")
	}

	return project
}

//Adds Board to Database
func addBoard(board structures.Board) {
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
