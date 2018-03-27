package main

import (
	"net/http"
	"../../../core"
	"../../projects"
	"gopkg.in/mgo.v2/bson"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}
var Database = ServiceDatabase{&core.Dao}



//Checks if User and Email does not exists in Database
func (r ServiceDatabase) checkFieldsExistance(project projects.Project) bool {
	r.Dao.C("projects")

	count, err := Database.Dao.Collection.Find(bson.M{"name": project.Name, "user": project.Users[0]}).Count()
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
func (r ServiceDatabase) validate(project projects.Project) bool {

	if project.Name == ""{
		core.SetResponse("empty_fields")
		return false
	}

	return Database.checkFieldsExistance(project)
}


//Adds User to Database
func (r ServiceDatabase) addProject(project projects.Project) bool {
	r.Dao.C("projects")

	project.ID = bson.NewObjectId()

	err := r.Dao.Collection.Insert(&project)
	if err != nil {
		core.SetResponse("database_error")
		return false
	}
	core.SetResponse("project_created")
	core.SetData(project.ID)
	return true
}


//Connects to database and listens to port
func main() {
	Database.Dao.Connect(core.Config.DatabaseHost + ":" + core.Config.DatabasePort, core.Config.DatabaseName)
	http.HandleFunc("/", do)
	http.ListenAndServe(core.Config.Host + ":1336", nil)
}

func do(w http.ResponseWriter, r *http.Request) {
	core.CORS(w)

	//Parses request data to
	var data projects.ProjectCreation
	if core.DecodeRequest(&data, r){

		success,user := Database.Dao.CheckSession(data.Session)
		if success {
			data.Project.Users = []projects.ProjectUser{
				{
					ID:user,
					Creator:true,
				},
			}
			if Database.validate(data.Project) {
				Database.addProject(data.Project) //Adds project to database
			}
		}
	}

	//Prints R
	core.PrintReponse(w)
}
