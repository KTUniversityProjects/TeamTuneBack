package main

import (
	"net/http"
	"../../../core"
	_ "gopkg.in/mgo.v2/bson"
	"../../projects"
	"../../../core/structures"
	_ "fmt"
	_ "golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"os/user"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}
var Database = ServiceDatabase{&core.Dao}

//Check if correct username and password
func (r ServiceDatabase) getList(session core.Session) string {
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


//Connects to database and listens to port
func main() {
	Database.Dao.Connect(core.Config.DatabaseHost + ":" + core.Config.DatabasePort, core.Config.DatabaseName)
	http.HandleFunc("/", do)
	http.ListenAndServe(core.Config.Host + ":1337", nil)
}

func do(w http.ResponseWriter, r *http.Request) {
	core.CORS(w)

	//Parses request data to
	var data core.Session
	if core.DecodeRequest(&data, r){

		var success = false
		success,data.Project.User = Database.Dao.CheckSession(data)
		if success {
			Database.getList(data.Project) //Adds project to database
		}
	}

	//Prints R
	core.PrintReponse(w)
}
