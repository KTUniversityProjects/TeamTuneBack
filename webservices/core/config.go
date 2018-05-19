package core

import "os"

//Main project config
var Config = ConfigStruct{

	//Mongo Database
	DatabaseHost: "localhost",
	DatabasePort: "27017",
	DatabaseName: "teamtune",

	//Server
	Host: "localhost"}

var ConfigDev = ConfigStruct{

	//Mongo Database
	DatabaseHost: "localhost",
	DatabasePort: "27017",
	DatabaseName: "teamtune",

	//Server
	Host: "localhost"}


func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

type ConfigStruct struct {
	DatabaseHost string
	DatabasePort string
	DatabaseName string
	Host string
}