package main

import "gopkg.in/mgo.v2/bson"
import "../../core"
import (
	"../../core/structures"
	"fmt"
)

//Adds Project to Database
func addTask(task structures.Task) {
	core.Dao.C("tasks")

	task.ID = bson.NewObjectId()

	err := core.Dao.Collection.Insert(&task)
	if err != nil {
		fmt.Println(err)
		core.ThrowResponse("database_error")
	}
	core.Dao.C("boards")
	err = core.Dao.Collection.Update(bson.M{"_id": task.BoardID }, bson.M{"$push": bson.M{"tasks" : task.ID}})
	if err != nil {
		core.ThrowResponse("database_error")
	}
	core.SetData(task.ID)
}

func getList(BoardID bson.ObjectId) {
	core.Dao.C("tasks")

	var results []structures.Task
	err := core.Dao.Collection.Find(bson.M{"board":BoardID}).Select(bson.M{"_id": 1, "name":1}).All(&results)
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.SetData(results)
}



