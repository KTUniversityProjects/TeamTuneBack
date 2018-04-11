package main

import (
	"../../core"
	"../../core/structures"
)

var servicePort = "1339"

func main() {
	core.AddRouting("POST", loginHandler)
	core.AddRouting("PUT", registerHandler)
	core.AddRouting("DELETE", deleteHandler)
	core.Initialize(servicePort)
}

func deleteHandler(){

	// gaudai structures.Session, reiškia kai darysi requestą turėsi siųsti datą tokiu formatu - {session:{id:"vaCiaSessijosID"}}
	var data structures.Session
	core.DecodeRequest(&data)

	//Va čia gausi user id
	userID := core.Dao.CheckSession(data)
	deleteUser(userID) //User ID jau grįžta kaip bson.ObjectId, tai nereikia čia, užtenka userID parašyt
}

func registerHandler() {

	var data structures.User
	//Parses request data
	core.DecodeRequest(&data)

	//Validates register data
	validate(data)

	//Adds user to database
	addUser(data)
}

func loginHandler() {

	//Parses request data to
	var data structures.User
	core.DecodeRequest(&data)

	//Credentials check
	userID := checkCredentials(data)

	//Creating session
	createSession(userID) //Logs in
}