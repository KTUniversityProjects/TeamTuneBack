package core

import "os"

//Main project config
var Config = ConfigStruct{

	//Mongo Database
	DatabaseHost: "mongo",
	DatabasePort: "27017",
	DatabaseName: "teamtune"}

var ConfigDev = ConfigStruct{

	//Mongo Database
	DatabaseHost: "mongo",
	DatabasePort: "27017",
	DatabaseName: "teamtune"}


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