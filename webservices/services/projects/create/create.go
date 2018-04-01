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
	user := Database.Dao.CheckSession(data.Session)

	//sets user as creator
	data.Project.Users = []projects.ProjectUser{
		{
			ID:      user,
			Creator: true,
		},
	}

	//validates
	Database.validate(data.Project)

	//Adds project to database
	Database.addProject(data.Project)
}


func (r ServiceDatabase) checkFieldsExistance(project projects.Project) {
	r.Dao.C("projects")

	count, err := Database.Dao.Collection.Find(bson.M{"name": project.Name, "users": bson.M{"$elemMatch": project.Users[0]}}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}
	if count > 0 {
		core.ThrowResponse("name_exists")
	}
}

func (r ServiceDatabase) validate(project projects.Project) {

	if project.Name == ""{
		core.ThrowResponse("empty_fields")
	}

	Database.checkFieldsExistance(project)
}


//Adds Project to Database
func (r ServiceDatabase) addProject(project projects.Project) {
	r.Dao.C("projects")

	project.ID = bson.NewObjectId()

	err := r.Dao.Collection.Insert(&project)
	if err != nil {
		core.ThrowResponse("database_error")
	}
	core.SetData(project.ID)
	core.ThrowResponse("project_created")
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