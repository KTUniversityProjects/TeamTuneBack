package core

import  (
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"fmt"
)
var Dao = MongoDatabase{}

type MongoDatabase struct {
	Session *mgo.Session
	Database *mgo.Database
	Collection *mgo.Collection
	Error error
}

func (r *MongoDatabase) D(databaseName string){
	r.Database = r.Session.DB(databaseName)
}

func (r *MongoDatabase) C(collection string){
	r.Collection = r.Database.C(collection)
}

func (r *MongoDatabase) Start(host string, databaseName string){

	r.Session, r.Error = mgo.Dial(host)
	if r.Error != nil{
		fmt.Print(r.Error)
	}

	r.D(databaseName)
}

