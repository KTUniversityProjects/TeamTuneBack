package main

import (
	"gopkg.in/mgo.v2/bson"
	"core"
	"fmt"
)

//Adds task to Database
func addTask(task TaskCreationRequest) {
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

func editBoard(data core.Task) {
	core.Dao.C("tasks")

	if data.Name != "" {
		err := core.Dao.Collection.Update(bson.M{"_id": data.ID}, bson.M{"$set": bson.M{"name": data.Name}})
		if err != nil {
			fmt.Println(err)
			core.ThrowResponse("database_error")
		}
	}
}

func getList(boardID bson.ObjectId) {
	core.Dao.C("tasks")

	var results []core.Task
	err := core.Dao.Collection.Find(bson.M{"board": boardID}).Select(bson.M{"_id": 1, "name": 1}).All(&results)
	if err != nil {
		core.ThrowResponse("database_error")
	}
	core.SetData(results)
}

//Validates project Data
func validate(task core.Task) {

}

//Deletes Task from database
func deleteTask(taskID bson.ObjectId) {
	core.Dao.C("tasks")

	err := core.Dao.Collection.Remove(bson.M{"_id":taskID})
	if err != nil {
		core.ThrowResponse("database_error")
	}
}

//checks if board contains user
func checkUser(boardID bson.ObjectId, userID bson.ObjectId) {

	var board = core.Board{}

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

//checks if project contains user
func checkTaskUser(taskID bson.ObjectId, userID bson.ObjectId) {

	var task = core.Task{}

	core.Dao.C("tasks")
	err := core.Dao.Collection.Find(bson.M{"_id": taskID}).Select(bson.M{"board": 1}).One(&task) //gauni project id

	if err != nil || task.BoardID == "" {
		fmt.Println(err)
		core.ThrowResponse("no_permission")
	}

	//redirects to already defined function which does remaining check
	checkUser(task.BoardID, userID)
}
