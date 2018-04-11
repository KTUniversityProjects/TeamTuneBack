package main

import (
	"gopkg.in/mgo.v2/bson"
	"../../core"
	"../../core/structures"
	"fmt"
)

//gets board list by ProjectID
func getList(projectID bson.ObjectId) {
	core.Dao.C("boards")

	var results []structures.Board
	fmt.Println(projectID)
	err := core.Dao.Collection.Find(bson.M{"project":projectID}).Select(bson.M{"_id": 1, "name":1}).All(&results)
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.SetData(results)
}



//Check if user can see project
func CheckUser(projectID bson.ObjectId, userID bson.ObjectId) {
	core.Dao.C("projects")

	count,err := core.Dao.Collection.Find(bson.M{"_id":projectID,"users":bson.M{"$elemMatch":bson.M{"_id":userID}}}).Count()
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

	if board.Name == ""{
		core.ThrowResponse("empty_fields")
	}


	//Checks for permissions
	if !project.Users[0].Creator {
		core.ThrowResponse("no_permission")
	}

	if board.Name == ""{
		core.ThrowResponse("system_mistake")
	}

	checkFieldsExistance(board)
}

//Gets project for board creation
func getProject(board structures.Board, user bson.ObjectId)  structures.Project {
	core.Dao.C("projects")

	var project = structures.Project{}
	err := core.Dao.Collection.Find(bson.M{"_id": board.ProjectID, "users": bson.M{"$elemMatch": bson.M{"_id" : user}}}).One(&project)

	if err != nil || project.ID == ""{
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
	err = core.Dao.Collection.Update(bson.M{"_id": board.ProjectID}, bson.M{"$push": bson.M{"boards" : board.ID}})
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.SetData(board.ID)
	core.ThrowResponse("board_created")
}
