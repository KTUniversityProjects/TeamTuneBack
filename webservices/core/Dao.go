package core

import  (
	"gopkg.in/mgo.v2"
	"fmt"
	"github.com/StulIK/TeamTune/webservices/core/structures"
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
func (r *MongoDatabase) CheckSession(session structures.Session) (bson.ObjectId){
	r.C("sessions")

	//finds session
	err := r.Collection.Find(bson.M{"_id":session.SessionID, "user":session.UserID, "expires" : bson.M{"$gt": time.Now()}}).One(&session)
	if err != nil {
		fmt.Println("CheckSession")
		fmt.Println(err)
		ThrowResponse("wrong_session")
	}

	//updates current session expiration time
	err = r.Collection.Update(bson.M{"_id":session.SessionID}, bson.M{"$set": bson.M{"expires" : time.Now().Add(time.Duration(24*time.Hour))}})
	if err != nil {
		fmt.Println("UpdateSession")
		fmt.Println(err)
		ThrowResponse("database_error")
	}

	//deletes old unaccessible sessions
	r.Collection.RemoveAll(bson.M{"user":session.UserID, "expires" : bson.M{"$lt": time.Now()}})

	return session.UserID
}

