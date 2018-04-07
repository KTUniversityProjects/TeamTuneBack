package main

import (
	"../../../core"
	"gopkg.in/mgo.v2/bson"
	"../../projects"
	"fmt"
)

var servicePort = "1340"

func do() {

	//Parses request data to
	var data projects.ProjectRequest
	core.DecodeRequest(&data)

	//Creating session
	//CreateSession(data, bson.ObjectId(userID)) //Logs in

	userID := core.Dao.CheckSession(data.Session)
	deletesUser(bson.ObjectId((userID)))
}

func deletesUser(userID bson.ObjectId) {

	err := core.Dao.Collection.Remove(bson.M{"user":bson.ObjectId(userID)})// nezinojau ka tikslai i sita rasyt reikia

	if(err != nil){
		core.ThrowResponse("database_error")
	}
}

/*           Every Webservice             */
func main(){
	core.Initialize(do, servicePort)
}