package main

import (
	"net/http"
	"../../../core"
	"../../../core/structures"
	"gopkg.in/mgo.v2/bson"
	"../../users"
	"time"
	"fmt"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}

//Check if correct username and password
func (r ServiceDatabase) checkCredentials(user users.User) (bool, bson.ObjectId) {
	r.Dao.C("users")

	var login = users.User{}
	err := Database.Dao.Collection.Find(bson.M{"username": user.Username}).Select(bson.M{"_id" :1, "password" : 1}).One(&login)
	if err != nil {
		fmt.Println("Wrong username")
		core.SetResponse("wrong_credentials")
		return false, login.Id
	}

	if success := users.CheckPasswordHash(user.Password, login.Password); !success {
		fmt.Println("Wrong password")
		core.SetResponse("wrong_credentials")
		return false, login.Id
	}

	return true, login.Id
}

//Check if correct username and password
func (r ServiceDatabase) CreateSession(user users.User, userID bson.ObjectId) bool {
	r.Dao.C("sessions")
	i := bson.NewObjectId()

	//Create session object for insertion
	var session = structures.Session{
		SessionID:i,UserID:userID,
		Expires:time.Now().Add(time.Duration(24 * time.Hour))}

	//Database insert
	err := r.Dao.Collection.Insert(&session)
	if err != nil {
		core.SetResponse("database_error")
		return false
	}

	core.SetResponse("logged_in")
	core.SetData(i)

	return true
}

var Database = ServiceDatabase{&core.Dao}

//Connects to database and listens to port
func main() {
	core.Initialize(do, "1338")
}

func do(w http.ResponseWriter, r *http.Request) {
	core.CORS(w)

	//Parses request data to
	var data users.User
	if core.DecodeRequest(&data, r){
		//Checks Username and Email
		if success, userID := Database.checkCredentials(data); success {
			Database.CreateSession(data, bson.ObjectId(userID)) //Logs in
		}
	}

	//Prints R
	core.PrintReponse(w)
}
