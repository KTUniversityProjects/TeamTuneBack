package core

import  (
	"gopkg.in/mgo.v2"
	"fmt"
	"./structures"
)

//Database Instance
var Dao = MongoDatabase{}

//Structure for Mongo Database
type MongoDatabase struct {
	Session *mgo.Session
	Database *mgo.Database
	Collection *mgo.Collection
	Error error
}

//Method for selecting Database
func (r *MongoDatabase) D(databaseName string){
	r.Database = r.Session.DB(databaseName)
}

//Method for selecting Collection
func (r *MongoDatabase) C(collection string){
	r.Collection = r.Database.C(MGOCollections[collection])
}

//Method for connection
func (r *MongoDatabase) Connect(host string, databaseName string){

	r.Session, r.Error = mgo.Dial(host)
	if r.Error != nil{
		fmt.Print(r.Error)
	}

	r.D(databaseName)
}

//Method for connection
func (r *MongoDatabase) CheckSession(session structures.Session) (bool, string){
	r.C("sessions")

	err := r.Collection.Find(session).One(&session)
	if err != nil {
		SetResponse("wrong_session")
		return false, session.UserID
	}

	return true, session.UserID
}

