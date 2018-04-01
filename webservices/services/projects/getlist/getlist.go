package main

import (
	"../../../core"
	"../../projects"
	"gopkg.in/mgo.v2/bson"
)


var servicePort = "1337"

func do() {

	//Parses request data to
	var data projects.ProjectRequest
	core.DecodeRequest(&data)

	//Gets user
	UserID := core.Dao.CheckSession(data.Session)

	//Gets all projects
	getList(UserID)
}


//Check if correct username and password
func getList(userID bson.ObjectId)  {
	core.Dao.C("projects")

	var results []projects.Project

	err := core.Dao.Collection.Find(bson.M{"users": bson.M{"$elemMatch":bson.M{"_id":userID}}}).Select(bson.M{"_id": 1, "name":1}).All(&results)
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.SetData(results)
	core.ThrowResponse("list_retrieved")
}


/*           Every Webservice             */
func main() {
	core.Initialize(do, servicePort)
}