package main

import (
	"net/http"
	"../../../core"
	"../../projects"
	"../../boards"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}
var Database = ServiceDatabase{&core.Dao}



//Checks if User and Email does not exists in Database
func (r ServiceDatabase) checkFieldsExistance(board boards.Board) bool {
	r.Dao.C("projects")

	count, err := Database.Dao.Collection.Find(bson.M{"name": board.Name, "project": board.ProjectID}).Count()
	if err != nil {
		core.SetResponse("database_error")
		return false
	}
	if count > 0 {
		core.SetResponse("name_exists")
		return false
	}
	return true
}

//Checks if User and Email does not exists in Database
func (r ServiceDatabase) validate(board boards.Board, project projects.Project)  bool {

	if board.Name == ""{
		core.SetResponse("empty_fields")
		return false
	}


	//Checks for permissions
	if !project.Users[0].Creator {
		core.SetResponse("no_permission")
		return false
	}

	if board.Name == ""{
		core.SetResponse("system_mistake")
		return false
	}

	return Database.checkFieldsExistance(board)
}

//Checks if User and Email does not exists in Database
func (r ServiceDatabase) getProject(board boards.Board, user bson.ObjectId)  projects.Project {
	Database.Dao.C("projects")

	var project = projects.Project{}
	err := Database.Dao.Collection.Find(bson.M{"_id": board.ProjectID, "users": bson.M{"$elemMatch": bson.M{"_id" : user}}}).One(&project)

	if err != nil || project.ID == ""{
		core.SetResponse("project_not_exists")
		return project
	}

	return project
}



//Adds Board to Database
func (r ServiceDatabase) addBoard(board boards.Board) bool {
	r.Dao.C("boards")

	board.ID = bson.NewObjectId()

	err := r.Dao.Collection.Insert(&board)
	if err != nil {
		core.SetResponse("database_error")
		return false
	}

	r.Dao.C("projects")
	err = Database.Dao.Collection.Update(bson.M{"_id": board.ProjectID}, bson.M{"$push": bson.M{"boards" : board.ID}})
	if err != nil {
		core.SetResponse("database_error")
		return false
	}

	core.SetResponse("board_created")
	core.SetData(board.ID)
	return true
}


//Connects to database and listens to port
func main() {
	Database.Dao.Connect(core.Config.DatabaseHost + ":" + core.Config.DatabasePort, core.Config.DatabaseName)
	http.HandleFunc("/", do)
	http.ListenAndServe(core.Config.Host + ":1335", nil)
}

func do(w http.ResponseWriter, r *http.Request) {
	core.CORS(w)

	//Parses request data to
	var data boards.BoardCreation
	if core.DecodeRequest(&data, r){

		success,user := Database.Dao.CheckSession(data.Session)
		if success {
			project := Database.getProject(data.Board, user)
			if Database.validate(data.Board, project) {
				Database.addBoard(data.Board) //Adds project to database
			}
		}
	}

	//Prints R
	core.PrintReponse(w)
}
