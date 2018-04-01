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
	Database.validate(data)

	//Adds user to database
	Database.addUser(data)
}


//Adds User to Database
func (r ServiceDatabase) addUser(user users.User) bool {
	r.Dao.C("users")

	var err error
	user.Password,err = users.EncryptPassword(user.Password)
	if err != nil {
		core.ThrowResponse("encryption_error")
	}

	err = r.Dao.Collection.Insert(bson.M{"username":user.Username, "password":user.Password, "email":user.Email})
	if err != nil {
		core.ThrowResponse("database_error")
	}
}

//Checks if User and Email does not exists in Database
func (r ServiceDatabase) checkFieldsExistance(user users.User) bool {
	r.Dao.C("users")

	count, err := Database.Dao.Collection.Find(bson.M{"username": user.Username}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
		return false
	}
	if count > 0 {
		core.ThrowResponse("username_exists")
		return false
	}

	count, err = Database.Dao.Collection.Find(bson.M{"email": user.Email}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
		return false
	}
	if count > 0 {
		core.ThrowResponse("email_exists")
		return false
	}
	return true
}

//Checks if User and Email does not exists in Database
func (r ServiceDatabase) validate(user users.User) bool {

	if user.Username == "" || user.Password == "" || user.Email == "" {
		core.ThrowResponse("empty_fields")
		return false
	}

	if  user.Password2 != user.Password {
		core.ThrowResponse("password_match")
		return false
	}

	return Database.checkFieldsExistance(user)
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