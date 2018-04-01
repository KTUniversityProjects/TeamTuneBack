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
	userID := Database.checkCredentials(data)

	//Creating session
	Database.CreateSession(data, bson.ObjectId(userID)) //Logs in
}

//Check if correct username and password
func (r ServiceDatabase) checkCredentials(user users.User) (bson.ObjectId) {
	r.Dao.C("users")

	var login = users.User{}
	err := Database.Dao.Collection.Find(bson.M{"username": user.Username}).Select(bson.M{"_id" :1, "password" : 1}).One(&login)
	if err != nil {
		fmt.Println("Wrong username")
		core.ThrowResponse("wrong_credentials")
	}

	if success := users.CheckPasswordHash(user.Password, login.Password); !success {
		fmt.Println("Wrong password")
		core.ThrowResponse("wrong_credentials")
	}

	return login.Id
}

//Check if correct username and password
func (r ServiceDatabase) CreateSession(user users.User, userID bson.ObjectId) {
	r.Dao.C("sessions")
	i := bson.NewObjectId()

	//Create session object for insertion
	var session = structures.Session{
		SessionID:i,UserID:userID,
		Expires:time.Now().Add(time.Duration(24 * time.Hour))}

	//Database insert
	err := r.Dao.Collection.Insert(&session)
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.SetData(i)
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