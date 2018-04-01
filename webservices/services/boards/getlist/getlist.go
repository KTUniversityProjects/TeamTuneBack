package main

import (
	"net/http"
	"../../../core"
	"../../boards"
	"gopkg.in/mgo.v2/bson"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}
var Database = ServiceDatabase{&core.Dao}

//Check if correct username and password
func (r ServiceDatabase) getList(projectID bson.ObjectId) bool {
	r.Dao.C("boards")

	var results []boards.Board

	err := Database.Dao.Collection.Find(bson.M{"project":projectID}).Select(bson.M{"_id": 1, "name":1}).All(&results)
	if err != nil {
		core.SetResponse("database_error")
		return false
	}

	core.SetResponse("list_retrieved")
	core.SetData(results)
	return true
}



//Check if correct username and password
func (r ServiceDatabase) CheckUser(projectID bson.ObjectId, userID bson.ObjectId) bool {
	r.Dao.C("projects")

	count,err := Database.Dao.Collection.Find(bson.M{"_id":projectID,"users":bson.M{"$elemMatch":bson.M{"_id":userID}}}).Count()
	if err != nil {
		core.SetResponse("database_error")
		return false
	}

	if count == 0 {
		core.SetResponse("database_error")
		return false
	}

	return true
}


//Connects to database and listens to port
func main() {
	core.Initialize(do, "1334")
}

func do(w http.ResponseWriter, r *http.Request) {
	core.CORS(w)

	//Parses request data to
	var data boards.BoardListRequest
	if core.DecodeRequest(&data, r){

		success,UserID := Database.Dao.CheckSession(data.Session)
		if success {
			if Database.CheckUser(data.ProjectID, UserID){
				Database.getList(UserID) //Adds project to database
			}
		}
	}

	//Prints R
	core.PrintReponse(w)
}
