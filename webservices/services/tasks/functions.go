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
	err = core.Dao.Collection.Update(bson.M{"_id": task.Board.ID}, bson.M{"$push": bson.M{"tasks": task.Task.ID}})
	if err != nil {
		core.ThrowResponse("database_error")
	}
	core.SetData(task.Task.ID)
}

func getList(boardID bson.ObjectId) {
	core.Dao.C("tasks")

	var results []structures.Task
	err := core.Dao.Collection.Find(bson.M{"board": boardID}).Select(bson.M{"_id": 1, "name": 1}).All(&results)
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
func checkUser(boardID bson.ObjectId, userID bson.ObjectId) {

	var board = structures.Board{}

	core.Dao.C("boards")
	err := core.Dao.Collection.Find(bson.M{"_id": boardID}).Select(bson.M{"project": 1}).One(&board) //gauni project id

	if err != nil || board.ProjectID == "" {
		fmt.Println(err)
		core.ThrowResponse("no_permission")
	}

	core.Dao.C("projects")
	count, err := core.Dao.Collection.Find(bson.M{"_id": board.ProjectID, "users": bson.M{"$elemMatch": bson.M{"_id":userID}}}).Count()
	if err != nil{
		fmt.Println(err)
		core.ThrowResponse("database_error")
		fmt.Println(err)
	}
	if count == 0 {
		core.ThrowResponse("no_permission")
	}
}
