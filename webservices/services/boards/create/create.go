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
	user := Database.Dao.CheckSession(data.Session)

	//gets project
	project := Database.getProject(data.Board, user)

	//validates
	Database.validate(data.Board, project)

	//Adds project to database
	Database.addBoard(data.Board)
}


//Checks if User and Email does not exists in Database
func (r ServiceDatabase) checkFieldsExistance(board boards.Board) {
	r.Dao.C("projects")

	count, err := Database.Dao.Collection.Find(bson.M{"name": board.Name, "project": board.ProjectID}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}
	if count > 0 {
		core.ThrowResponse("name_exists")
	}
}

//Checks if User and Email does not exists in Database
func (r ServiceDatabase) validate(board boards.Board, project projects.Project) {

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

	Database.checkFieldsExistance(board)
}

//Checks if User and Email does not exists in Database
func (r ServiceDatabase) getProject(board boards.Board, user bson.ObjectId)  projects.Project {
	Database.Dao.C("projects")

	var project = projects.Project{}
	err := Database.Dao.Collection.Find(bson.M{"_id": board.ProjectID, "users": bson.M{"$elemMatch": bson.M{"_id" : user}}}).One(&project)

	if err != nil || project.ID == ""{
		core.ThrowResponse("project_not_exists")
	}

	return project
}



//Adds Board to Database
func (r ServiceDatabase) addBoard(board boards.Board) {
	r.Dao.C("boards")

	board.ID = bson.NewObjectId()

	err := r.Dao.Collection.Insert(&board)
	if err != nil {
		core.ThrowResponse("database_error")
	}

	r.Dao.C("projects")
	err = Database.Dao.Collection.Update(bson.M{"_id": board.ProjectID}, bson.M{"$push": bson.M{"boards" : board.ID}})
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.SetData(board.ID)
	core.ThrowResponse("board_created")
}

/*           Every Webservice             */
type ServiceDatabase struct {
	Dao *core.MongoDatabase
}

var Database = ServiceDatabase{&core.Dao}

//Connects to database and listens to port
func main() {
	core.Initialize(do, servicePort)
}