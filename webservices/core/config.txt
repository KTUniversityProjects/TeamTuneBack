package core

//Main project config
var Config = ConfigStruct{

	//Mongo Database
	DatabaseHost: "localhost",
	DatabasePort: "27017",
	DatabaseName: "teamtune",

	//Server
	Host: "localhost"}

//Mongo database collections
var MGOCollections = map[string]string{
	"users":    "users",
	"sessions": "sessions",
	"boards":   "boards",
	"projects": "projects"}


//EOF

type ConfigStruct struct {
	DatabaseHost string
	DatabasePort string
	DatabaseName string
	Host string
}