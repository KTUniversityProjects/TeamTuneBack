package main

import (
	"net/http"
	"../core"
	"gopkg.in/mgo.v2/bson"
	"../structures/users"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}

func (r ServiceDatabase) addUser(user users.RegisterStructure) bool {
	r.Dao.C("users")

	err := r.Dao.Collection.Insert(&user)
	if err != nil {
		core.SetReponse("database_error")
		return false
	}
	core.SetReponse("user_created")
	return true
}

func (r ServiceDatabase) checkFieldsExistance(user users.RegisterStructure) bool {
	r.Dao.C("users")

	count, err := Database.Dao.Collection.Find(bson.M{"Username": user.Username}).Count()
	if err != nil {
		core.SetReponse("database_error")
		return false
	}
	if count > 0 {
		core.SetReponse("username_exists")
		return false
	}

	count, err = Database.Dao.Collection.Find(bson.M{"Email": user.Email}).Count()
	if err != nil {
		core.SetReponse("database_error")
		return false
	}
	if count > 0 {
		core.SetReponse("email_exists")
		return false
	}
	return true
}

var Database = ServiceDatabase{&core.Dao}

func main() {
	Database.Dao.Start("localhost:27017", "teamtune")
	http.HandleFunc("/", do)
	http.ListenAndServe("localhost:1339", nil)
}

func do(w http.ResponseWriter, r *http.Request) {
	core.CORS(w)

	//Parses request data to
	var data users.RegisterStructure
	if core.DecodeRequest(&data, r){
		//Checks Username and Email
		if Database.checkFieldsExistance(data) {
			Database.addUser(data) //Adds user to database
		}
	}

	//Prints R
	core.PrintReponse(w)
}
