package main

import (
	"net/http"
	"../../core"
	"gopkg.in/mgo.v2/bson"
	"../../users"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}

//Adds User to Database
func (r ServiceDatabase) addUser(user users.UserStruct) bool {
	r.Dao.C("users")

	user.Password = users.EncryptPassword(user.Password)
	err := r.Dao.Collection.Insert(&user)
	if err != nil {
		core.SetResponse("database_error")
		return false
	}
	core.SetResponse("user_created")
	return true
}

//Checks if User and Email does not exists in Database
func (r ServiceDatabase) checkFieldsExistance(user users.UserStruct) bool {
	r.Dao.C("users")

	count, err := Database.Dao.Collection.Find(bson.M{"username": user.Username}).Count()
	if err != nil {
		core.SetResponse("database_error")
		return false
	}
	if count > 0 {
		core.SetResponse("username_exists")
		return false
	}

	count, err = Database.Dao.Collection.Find(bson.M{"email": user.Email}).Count()
	if err != nil {
		core.SetResponse("database_error")
		return false
	}
	if count > 0 {
		core.SetResponse("email_exists")
		return false
	}
	return true
}

//Checks if User and Email does not exists in Database
func (r ServiceDatabase) validate(user users.UserStruct) bool {

	if user.Username == "" || user.Password == "" || user.Email == "" {
		core.SetResponse("empty_fields")
		return false
	}

	if  user.Password2 != user.Password {
		core.SetResponse("password_match")
		return false
	}

	return Database.checkFieldsExistance(user)
}

var Database = ServiceDatabase{&core.Dao}

//Connects to database and listens to port
func main() {
	Database.Dao.Connect("localhost:27017", "teamtune")
	http.HandleFunc("/", do)
	http.ListenAndServe("localhost:1339", nil)
}

func do(w http.ResponseWriter, r *http.Request) {
	core.CORS(w)

	//Parses request data to
	var data users.UserStruct
	if core.DecodeRequest(&data, r){
		//Validates sent params
		if Database.validate(data) {
			Database.addUser(data) //Adds user to database
		}
	}

	//Prints R
	core.PrintReponse(w)
}
