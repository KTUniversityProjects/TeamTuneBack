package main

import (
	"net/http"
	"../../../core"
	"gopkg.in/mgo.v2/bson"
	"../../users"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}

//Adds User to Database
func (r ServiceDatabase) addUser(user users.User) bool {
	r.Dao.C("users")

	var err error
	user.Password,err = users.EncryptPassword(user.Password)
	if err != nil {
		core.SetResponse("encryption_error")
		return false
	}

	err = r.Dao.Collection.Insert(bson.M{"username":user.Username, "password":user.Password, "email":user.Email})
	if err != nil {
		core.SetResponse("database_error")
		return false
	}
	core.SetResponse("user_created")
	return true
}

//Checks if User and Email does not exists in Database
func (r ServiceDatabase) checkFieldsExistance(user users.User) bool {
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
func (r ServiceDatabase) validate(user users.User) bool {

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
	core.Initialize(do, "1339")
}

func do(w http.ResponseWriter, r *http.Request) {
	core.CORS(w)

	//Parses request data to
	var data users.User
	if core.DecodeRequest(&data, r){
		//Validates sent params
		if Database.validate(data) {
			Database.addUser(data) //Adds user to database
		}
	}

	//Prints R
	core.PrintReponse(w)
}
