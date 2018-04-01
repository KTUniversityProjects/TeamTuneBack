package main

import (
	"../../../core"
	"gopkg.in/mgo.v2/bson"
	"../../users"
)

var servicePort = "1339"

func do() {

	var data users.User
	//Parses request data
	core.DecodeRequest(&data)

	//Validates register data
	validate(data)

	//Adds user to database
	addUser(data)
}


//Adds User to Database
func addUser(user users.User) {
	core.Dao.C("users")

	var err error
	user.Password,err = users.EncryptPassword(user.Password)
	if err != nil {
		core.ThrowResponse("encryption_error")
	}

	err = core.Dao.Collection.Insert(bson.M{"username":user.Username, "password":user.Password, "email":user.Email})
	if err != nil {
		core.ThrowResponse("database_error")
	}
}

//Checks if User and Email does not exists in Database
func checkFieldsExistance(user users.User) {
	core.Dao.C("users")

	count, err := core.Dao.Collection.Find(bson.M{"username": user.Username}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}
	if count > 0 {
		core.ThrowResponse("username_exists")
	}

	count, err = core.Dao.Collection.Find(bson.M{"email": user.Email}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}
	if count > 0 {
		core.ThrowResponse("email_exists")
	}
}

//Validates form fields
func validate(user users.User) {

	if user.Username == "" || user.Password == "" || user.Email == "" {
		core.ThrowResponse("empty_fields")
	}

	if  user.Password2 != user.Password {
		core.ThrowResponse("password_match")
	}

	checkFieldsExistance(user)
}


/*           Every Webservice             */
func main() {
	core.Initialize(do, servicePort)
}