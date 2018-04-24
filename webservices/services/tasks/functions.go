package main

import "gopkg.in/mgo.v2/bson"
import "../../core"
import (
	"../../core/structures"
	"fmt"
)

//Adds task to Database
func addTask(task structures.TaskCreationRequest) {
	core.Dao.C("tasks")

	task.Task.ID = bson.NewObjectId()

	err := core.Dao.Collection.Insert(&task.Task)
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.Dao.C("boards")
	err = core.Dao.Collection.Update(bson.M{"_id": task.Board.ID }, bson.M{"$push": bson.M{"tasks" : task.Task.ID}})
	if err != nil {
		core.ThrowResponse("database_error")
	}
	core.SetData(task.Task.ID)
}

func getList(boardID bson.ObjectId) {
	core.Dao.C("tasks")

	var results []structures.Task
	err := core.Dao.Collection.Find(bson.M{"board":boardID}).Select(bson.M{"_id": 1, "name":1}).All(&results)
	if err != nil {
		core.ThrowResponse("database_error")
	}
	core.SetData(results)
}

//Validates project Data
func validate(task structures.Task) {

	if task.Name == "" {
		core.ThrowResponse("empty_fields")
	}

	checkFieldsExistence(task)
}

//Chech Project existance for User
func checkFieldsExistence(project structures.Task) {
	core.Dao.C("projects")

	count, err := core.Dao.Collection.Find(bson.M{"name": project.Name}).Count()
	fmt.Println(count)
	if err != nil {
		core.ThrowResponse("database_error")
	}
	if count > 0 {
		core.ThrowResponse("name_exists")
	}
}

//checks if project contains user
func checkProject(task structures.TaskCreationRequest, user bson.ObjectId) {
	core.Dao.C("boards")
    fmt.Printf("projet %s \nuser %s\n",task.Board.ProjectID,user)
	var project = structures.Project{}
	var board = structures.Board{}
	err := core.Dao.Collection.Find(bson.M{"_id": task.Board.ID}).One(&board) //gauni project id
    fmt.Println(board.ProjectID)
	if err != nil || board.ProjectID == ""{
		fmt.Println(err)
		core.ThrowResponse("project_not_exists")
	}
	core.Dao.C("projects")
	err = core.Dao.Collection.Find(bson.M{"_id": board.ProjectID, "users": bson.M{"$elemMatch": bson.M{"_id" : user}}}).One(&project) //patikrini
	if err != nil || project.ID == ""{
		fmt.Println(err)
		core.ThrowResponse("project_not_exists")
	}
}



