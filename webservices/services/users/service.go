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

	var data structures.SessionRequest
	core.DecodeRequest(&data)
	userID := core.Dao.CheckSession(data.Session)
	deleteUser(userID)
	deleteSessions(userID)
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