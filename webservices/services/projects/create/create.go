package main

import (
	"../../../core"
	"../../projects"
	"gopkg.in/mgo.v2/bson"
)

var servicePort = "1336"

func do() {

	//Parses request data to
	var data projects.ProjectRequest
	core.DecodeRequest(&data)

	//Gets user
	user := core.Dao.CheckSession(data.Session)

	//sets user as creator
	data.Project.Users = []projects.ProjectUser{
		{
			ID:      user,
			Creator: true,
		},
	}

	//validates
	validate(data.Project)

	//Adds project to database
	addProject(data.Project)
}


func  checkFieldsExistance(project projects.Project) {
	core.Dao.C("projects")

	count, err := core.Dao.Collection.Find(bson.M{"name": project.Name, "users": bson.M{"$elemMatch": project.Users[0]}}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}
	if count > 0 {
		core.ThrowResponse("name_exists")
	}
}

func validate(project projects.Project) {

	if project.Name == ""{
		core.ThrowResponse("empty_fields")
	}

	checkFieldsExistance(project)
}


//Adds Project to Database
func addProject(project projects.Project) {
	core.Dao.C("projects")

	project.ID = bson.NewObjectId()

	err := core.Dao.Collection.Insert(&project)
	if err != nil {
		core.ThrowResponse("database_error")
	}
	core.SetData(project.ID)
	core.ThrowResponse("project_created")
}


/*           Every Webservice             */
func main() {
	core.Initialize(do, servicePort)
}