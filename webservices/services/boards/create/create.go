package main

import (
	"../../../core"
	"../../projects"
	"../../boards"
	"gopkg.in/mgo.v2/bson"
)


var servicePort = "1335"

func do() {

	//Parses request data to
	var data boards.BoardCreation
	core.DecodeRequest(&data)

	//Gets user
	user := core.Dao.CheckSession(data.Session)

	//gets project
	project := getProject(data.Board, user)

	//validates
	validate(data.Board, project)

	//Adds project to database
	addBoard(data.Board)
}


//Checks if User and Email does not exists in Database
func checkFieldsExistance(board boards.Board) {
	core.Dao.C("projects")

	count, err := core.Dao.Collection.Find(bson.M{"name": board.Name, "project": board.ProjectID}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}
	if count > 0 {
		core.ThrowResponse("name_exists")
	}
}

//Checks if User and Email does not exists in Database
func validate(board boards.Board, project projects.Project) {

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

//Checks if User and Email does not exists in Database
func getProject(board boards.Board, user bson.ObjectId)  projects.Project {
	core.Dao.C("projects")

	var project = projects.Project{}
	err := core.Dao.Collection.Find(bson.M{"_id": board.ProjectID, "users": bson.M{"$elemMatch": bson.M{"_id" : user}}}).One(&project)

	if err != nil || project.ID == ""{
		core.ThrowResponse("project_not_exists")
	}

	return project
}



//Adds Board to Database
func addBoard(board boards.Board) {
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

/*           Every Webservice             */
func main() {
	core.Initialize(do, servicePort)
}