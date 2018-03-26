package core

import  (
	"gopkg.in/mgo.v2"
	"fmt"
	"./structures"
	"gopkg.in/mgo.v2/bson"
	"time"
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
func (r *MongoDatabase) CheckSession(session structures.Session) (bool, bson.ObjectId){
	r.C("sessions")


	err := r.Collection.Find(bson.M{"_id":session.SessionID, "expires" : bson.M{"$gt": time.Now()}}).One(&session)
	if err != nil {
		SetResponse("wrong_session")
		return false, session.UserID
	}

	err = r.Collection.Update(bson.M{"_id":session.SessionID}, bson.M{"$set": bson.M{"expires" : time.Now().Add(time.Duration(24*time.Hour))}})
	if err != nil {
		SetResponse("database_error")
		return true, session.UserID
	}

	return true, session.UserID
}

