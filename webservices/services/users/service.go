package main

import (
	"core"
)


func main() {
	core.AddRouting("POST", loginHandler)
	core.AddRouting("PUT", registerHandler)
	core.AddRouting("DELETE", deleteHandler)
	core.Initialize()
}

func deleteHandler(){

	var data core.SessionRequest
	core.DecodeRequest(&data)
	userID := core.Dao.CheckSession(data.Session)
	deleteUser(userID)
	deleteSessions(userID)
}

func registerHandler() {

	var data core.User
	//Parses request data
	core.DecodeRequest(&data)

	//Validates register data
	validate(data)

	//Adds user to database
	addUser(data)
}

func loginHandler() {

	//Parses request data to
	var data core.User
	core.DecodeRequest(&data)

	//Credentials check
	userID := checkCredentials(data)

	//Creating session
	createSession(userID) //Logs in
}