package main

import (
	"../../../core"
	"../../boards"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

var servicePort = "1334"

func do() {

	//Parses request data to
	var data boards.BoardListRequest
	core.DecodeRequest(&data)

	UserID := core.Dao.CheckSession(data.Session)

	//Checks user
	CheckUser(data.ProjectID, UserID)

	//Gets project list
	getList(data.ProjectID)

}

//Check if correct username and password
func getList(projectID bson.ObjectId) {
	core.Dao.C("boards")

	var results []boards.Board
	fmt.Println(projectID)
	err := core.Dao.Collection.Find(bson.M{"project":projectID}).Select(bson.M{"_id": 1, "name":1}).All(&results)
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.SetData(results)
}



//Check if correct username and password
func  CheckUser(projectID bson.ObjectId, userID bson.ObjectId) {
	core.Dao.C("projects")

	count,err := core.Dao.Collection.Find(bson.M{"_id":projectID,"users":bson.M{"$elemMatch":bson.M{"_id":userID}}}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}

	if count == 0 {
		core.ThrowResponse("database_error")
	}
}



/*           Every Webservice             */
func main() {
	core.Initialize(do, servicePort)
}