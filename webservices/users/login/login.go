package main

import (
	"net/http"
	"../../core"
	"gopkg.in/mgo.v2/bson"
	"../../users"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}

//Check if correct username and password
func (r ServiceDatabase) checkCredentials(user users.LoginStructure) bool {
	r.Dao.C("users")

	user.Password = users.EncryptPassword(user.Password)

	var login = users.LoginStructure{}
	err := Database.Dao.Collection.Find(bson.M{"username": user.Username}).One(&login)
	if err != nil {
		fmt.Print(err)
		core.SetReponse("database_error")
		return false
	}

	fmt.Println(len(users.EncryptPassword(user.Password)))
	fmt.Println(len(login.Password))
	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(user.Password)); err != nil {
		fmt.Print(err)
		core.SetReponse("wrong_credentials")
		return false
	}

	return true
}

//Check if correct username and password
func (r ServiceDatabase) Login(user users.LoginStructure) bool {
	r.Dao.C("users")
	core.SetReponse("logged_in")
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
	var data users.LoginStructure
	if core.DecodeRequest(&data, r){
		//Checks Username and Email
		if Database.checkCredentials(data) {
			Database.Login(data) //Logs in
		}
	}

	//Prints R
	core.PrintReponse(w)
}
