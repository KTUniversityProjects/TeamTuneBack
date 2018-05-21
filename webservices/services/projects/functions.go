package main

import (
	"gopkg.in/mgo.v2/bson"
	"core"
	"fmt"
)

//Chech Project existance for User
func checkFieldsExistance(project core.Project) {
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
func validate(project core.Project) {

	if project.Name == "" {
		core.ThrowResponse("empty_fields")
	}

	checkFieldsExistance(project)
}

//Adds Project to Database
func addProject(project core.Project) {
	core.Dao.C("projects")

	project.ID = bson.NewObjectId()

	err := core.Dao.Collection.Insert(&project)
	if err != nil {
		core.ThrowResponse("database_error")
	}
	core.SetData(project.ID)
}

//Edits Project in Database
func editProject(data core.ProjectEdit) {
	core.Dao.C("projects")

	fmt.Println("name")
	fmt.Println(data.Name)
	if data.Name != "" {
		err := core.Dao.Collection.Update(bson.M{"_id": data.ID}, bson.M{"$set": bson.M{"name": data.Name}})
		if err != nil {
			fmt.Println(err)
			core.ThrowResponse("database_error")
		}
	}
	if data.Description != "" {
		err := core.Dao.Collection.Update(bson.M{"_id": data.ID}, bson.M{"$set": bson.M{"description": data.Description}})
		if err != nil {
			fmt.Println(err)
			core.ThrowResponse("database_error")
		}
	}
	if len(data.Users) != 0{
		var users []core.ProjectUser	//is projekto
		var users2 []bson.ObjectId		//is gautos uzklausos

		fmt.Println("test 1")
		err := core.Dao.Collection.Find(bson.M{"_id":data.ID}).Select(bson.M{"users":1}).All(&users) //gauna useriu id is projekto
		fmt.Println("test 2")
		if err != nil {
			fmt.Println(err)
			core.ThrowResponse("database_error")
		}

		//err = core.Dao.Collection.Find(bson.M{"_id":data.ID}).Select(bson.M{"users":1}).All(&users2) //gauna useriu id pagal atsiustus email
		//if err != nil {
		//	fmt.Println(err)
		//	core.ThrowResponse("database_error")
		//}

		fmt.Println(users)

		for i:=0;i < len(users2); i++{
			if stringInSlice(users2[i], users){
				fmt.Println("reikia pridet")
				fmt.Println(data.Users[i])
				addUser(data.Users[i])
			}
		}


	}

}
func addUser(email string){
	//core.Dao.C("users")
	//err := core.Dao.Collection.Find(bson.M{"email":email}.Select(bson.M{"users":1}).One(user)
	//if err != nil {
	//	fmt.Println(err)
	//	core.ThrowResponse("database_error")
	//}

}
func stringInSlice(a bson.ObjectId, list []core.ProjectUser) bool {
	for _, b := range list {
		if b.ID == a {
			return false
		}
	}
	return true
}

//Gets projects list by userID or One parcitular project of project.ID is not nil
func getList(userID bson.ObjectId, project core.Project) {
	core.Dao.C("projects")

	//picks one project
	if project.ID != "" {

		var result []core.Project
		err := core.Dao.Collection.Find(bson.M{"_id": project.ID, "users": bson.M{"$elemMatch": bson.M{"_id": userID}}}).One(&result)
		if err != nil {
			core.ThrowResponse("database_error")
		}

		core.SetData(result)
		//Gets all list
	} else {
		var results []core.Project
		err := core.Dao.Collection.Find(bson.M{"users": bson.M{"$elemMatch": bson.M{"_id": userID}}}).Select(bson.M{"_id": 1, "name": 1}).All(&results)
		if err != nil {
			core.ThrowResponse("database_error")
		}

		core.SetData(results)
	}
}

//Checks right for deleting
func checkUser(project core.Project) {

	core.Dao.C("projects")

	count, err := core.Dao.Collection.Find(bson.M{"_id": project.ID, "users": bson.M{"$elemMatch": project.Users[0]}}).Count()
	if err != nil {
		fmt.Println("FindUser")
		fmt.Println(err)
		core.ThrowResponse("database_error")
	}
	if count == 0 {
		fmt.Println("CheckProject")
		fmt.Println(err)
		core.ThrowResponse("no_permission")
	}
}

//Removes Project from Database
func removeBoards(projectID bson.ObjectId) {

	core.Dao.C("boards")

	var boards []core.Board
	err := core.Dao.Collection.Find(bson.M{"project": projectID}).Select(bson.M{"_id": 1}).All(&boards)
	if err != nil {
		fmt.Println("GetBoards")
		fmt.Println(err)
		core.ThrowResponse("database_error")
	}

	core.Dao.C("tasks")
	for _, element := range boards {
		_, err = core.Dao.Collection.RemoveAll(bson.M{"board": element.ID})
		if err != nil {
			fmt.Println("tasks delete")
			fmt.Println(err)
			core.ThrowResponse("database_error")
		}
		fmt.Println(element.Tasks)
	}

	core.Dao.C("boards")
	_, err = core.Dao.Collection.RemoveAll(bson.M{"project": projectID})
	if err != nil {
		fmt.Println("RemoveBoards")
		fmt.Println(err)
		core.ThrowResponse("database_error")
	}
}

//Removes Project from Database
func removeProject(projectID bson.ObjectId) {

	//RemoveProject
	core.Dao.C("projects")

	err := core.Dao.Collection.Remove(bson.M{"_id": projectID})
	if err != nil {
		fmt.Println("RemoveProject")
		fmt.Println(err)
		core.ThrowResponse("database_error")
	}
}
