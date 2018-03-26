package main

import (
	"net/http"
	"../../../core"
	"../../../core/structures"
	"gopkg.in/mgo.v2/bson"
	"../../users"
	_ "fmt"
	"golang.org/x/crypto/bcrypt"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}

//Check if correct username and password
func (r ServiceDatabase) checkCredentials(user users.User) string {
	r.Dao.C("users")

	user.Password = users.EncryptPassword(user.Password)

	var login = users.User{}
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
func (r ServiceDatabase) CreateSession(user users.User, userID string) bool {
	r.Dao.C("sessions")
	i := bson.NewObjectId()

	var session = structures.Session{SessionID:i,UserID:userID}
	err := r.Dao.Collection.Insert(&session)
	if err != nil {
		core.SetResponse("database_error")
		return false
	}

	core.SetResponse("logged_in")
	core.SetData(i.Hex())

	return true
}

var Database = ServiceDatabase{&core.Dao}

//Connects to database and listens to port
func main() {
	Database.Dao.Connect(core.Config.DatabaseHost + ":" + core.Config.DatabasePort, core.Config.DatabaseName)
	http.HandleFunc("/", do)
	http.ListenAndServe(core.Config.Host + ":1338", nil)
}

func do(w http.ResponseWriter, r *http.Request) {
	core.CORS(w)

	//Parses request data to
	var data users.User
	if core.DecodeRequest(&data, r){
		//Checks Username and Email
		if userID := Database.checkCredentials(data); userID != "" {
			Database.CreateSession(data, userID) //Logs in
		}
	}

	//Prints R
	core.PrintReponse(w)
}
