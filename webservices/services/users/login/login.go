package main

import (
	"../../../core"
	"../../../core/structures"
	"gopkg.in/mgo.v2/bson"
	"../../users"
	"time"
	"fmt"
)

var servicePort = "1338"

func do() {

	//Parses request data to
	var data users.User
	core.DecodeRequest(&data)

	//Credentials check
	userID := checkCredentials(data)

	//Creating session
	CreateSession(data, bson.ObjectId(userID)) //Logs in
}

//Check if correct username and password
func  checkCredentials(user users.User) (bson.ObjectId) {
	core.Dao.C("users")

	var login = users.User{}
	err := core.Dao.Collection.Find(bson.M{"username": user.Username}).Select(bson.M{"_id" :1, "password" : 1}).One(&login)
	if err != nil {
		fmt.Println("Wrong username")
		core.ThrowResponse("wrong_credentials")
	}

	if !users.CheckPasswordHash(user.Password, login.Password) {
		fmt.Println("Wrong password")
		core.ThrowResponse("wrong_credentials")
	}

	return login.Id
}

//Creates session / returns SessionID
func CreateSession(user users.User, userID bson.ObjectId) {
	core.Dao.C("sessions")
	sessionID := bson.NewObjectId()

	//Create session object for insertion
	var session = structures.Session{
		SessionID:sessionID,UserID:userID,
		Expires:time.Now().Add(time.Duration(24 * time.Hour))}

	//Database insert
	err := core.Dao.Collection.Insert(&session)
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.SetData(sessionID)
}


/*           Every Webservice             */
func main() {
	core.Initialize(do, servicePort)
}