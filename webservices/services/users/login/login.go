package main

import (
	"net/http"
	"../../../core"
	"gopkg.in/mgo.v2/bson"
	"../../users"
	_ "fmt"
	"golang.org/x/crypto/bcrypt"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}

//Check if correct username and password
func (r ServiceDatabase) checkCredentials(user users.UserStruct) string {
	r.Dao.C("users")

	user.Password = users.EncryptPassword(user.Password)

	var login = users.UserStruct{}
	err := Database.Dao.Collection.Find(bson.M{"username": user.Username}).One(&login)
	if err != nil {
		core.SetResponse("database_error")
		return ""
	}

	return bson.ObjectId(login.Id).Hex() //
	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(user.Password)); err != nil {
		core.SetResponse("wrong_credentials")
		return ""
	}

	return bson.ObjectId(login.Id).Hex()
}

//Check if correct username and password
func (r ServiceDatabase) GetUser(user users.UserStruct) bool {
	r.Dao.C("sessions")
	err := r.Dao.Collection.Insert(bson.M{"user": user.Username})
	if err != nil {
		core.SetResponse("database_error")
		return false
	}
	core.SetResponse("logged_in")
	return true
}

//Check if correct username and password
func (r ServiceDatabase) CreateSession(user users.UserStruct, userID string) bool {
	r.Dao.C("sessions")
	i := bson.NewObjectId()


	err := r.Dao.Collection.Insert(bson.M{"_id": i, "user": userID})
	if err != nil {
		core.SetResponse("database_error")
		return false
	}

	core.P = core.Response{ResponseCode: 0, ResponseMsg: i.Hex()}

	return true
}

var Database = ServiceDatabase{&core.Dao}

//Connects to database and listens to port
func main() {
	Database.Dao.Connect("localhost:27017", "teamtune")
	http.HandleFunc("/", do)
	http.ListenAndServe("localhost:1338", nil)
}

func do(w http.ResponseWriter, r *http.Request) {
	core.CORS(w)

	//Parses request data to
	var data users.UserStruct
	if core.DecodeRequest(&data, r){
		//Checks Username and Email
		if userID := Database.checkCredentials(data); userID != "" {
			Database.CreateSession(data, userID) //Logs in
		}
	}

	//Prints R
	core.PrintReponse(w)
}
