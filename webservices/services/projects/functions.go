package main

import "gopkg.in/mgo.v2/bson"
import "../../core"
import "../../core/structures"

//Chech Project existance for User
func  checkFieldsExistance(project structures.Project) {
	core.Dao.C("projects")

	count, err := core.Dao.Collection.Find(bson.M{"name": project.Name, "users": bson.M{"$elemMatch": project.Users[0]}}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}
	if count > 0 {
		core.ThrowResponse("name_exists")
	}
}

//Validates project Data
func validate(project structures.Project) {

	if project.Name == ""{
		core.ThrowResponse("empty_fields")
	}

	checkFieldsExistance(project)
}


//Adds Project to Database
func addProject(project structures.Project) {
	core.Dao.C("projects")

	project.ID = bson.NewObjectId()

	err := core.Dao.Collection.Insert(&project)
	if err != nil {
		core.ThrowResponse("database_error")
	}
	core.SetData(project.ID)
}

//Gets projects list by userID
func getList(userID bson.ObjectId)  {
	core.Dao.C("projects")

	var results []structures.Project

	err := core.Dao.Collection.Find(bson.M{"users": bson.M{"$elemMatch":bson.M{"_id":userID}}}).Select(bson.M{"_id": 1, "name":1}).All(&results)
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.SetData(results)
	core.ThrowResponse("list_retrieved")
}


//Checks right for deleting
func checkUser(project structures.Project) {

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

