package main

import (
	"../../../core"
	"../../projects"
	"gopkg.in/mgo.v2/bson"
)

var servicePort = "1333"

func do() {

	//Parses request data to
	var data projects.ProjectRequest
	core.DecodeRequest(&data)

	//Gets user
	user := core.Dao.CheckSession(data.Session)
	data.Project.Users = []projects.ProjectUser{
		{
			ID:user,
			Creator:true,
		},
	}

	//validates
	checkUser(data.Project)

	//Remove Boards
	removeBoards(data.Project.ID)

	//Remove Project
	removeProject(data.Project.ID)
}

//Checks right for deleting
func checkUser(project projects.Project) {

	core.Dao.C("projects")

	count, err := core.Dao.Collection.Find(bson.M{"_id": project.ID, "users": bson.M{"$elemMatch": project.Users[0]}}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}
	if count == 0{
		core.ThrowResponse("project_not_exists")
	}
}


//Removes Project from Database
func removeBoards(projectID bson.ObjectId) {

	core.Dao.C("boards")

	err := core.Dao.Collection.Remove(bson.M{"project":projectID})
	if err != nil {
		core.ThrowResponse("database_error")
	}
}


//Removes Project from Database
func removeProject(projectID bson.ObjectId) {

	//RemoveProject
	core.Dao.C("projects")

	err := core.Dao.Collection.Remove(bson.M{"_id":projectID})
	if err != nil {
		core.ThrowResponse("database_error")
	}
}


/*           Every Webservice             */
func main() {
	core.Initialize(do, servicePort)
}