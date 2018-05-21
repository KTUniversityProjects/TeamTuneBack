package core

import (
	"net/http"
	"encoding/json"
	"fmt"
	"os"
	"log"
)

func AddRouting(requestType string, function func()){
	dofunc[requestType] = function
}

var P Response
var connected = false

func Initialize(){
	AddRouting("OPTIONS", func(){})
	loadResponses()

	if Exists("C:/developer") {
		fmt.Println(ConfigDev.DatabaseHost)
		fmt.Println(ConfigDev.DatabasePort)
		fmt.Println(ConfigDev.DatabaseName)
		connected = Dao.Connect(ConfigDev.DatabaseHost + ":" + ConfigDev.DatabasePort, ConfigDev.DatabaseName)
	} else {
		connected = Dao.Connect(Config.DatabaseHost + ":" + Config.DatabasePort, Config.DatabaseName)
	}

	http.HandleFunc("/", requestFunc)
	http.ListenAndServe(":777", nil)
}


//Adds CORS header to response Writer
func CORS(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
}

func writeLog(error string) {
	f, err := os.OpenFile("logFile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.SetOutput(f)
	log.Println(error)
	f.Close()
}

var dofunc = make(map[string]func())
var currentRequest *http.Request

func requestFunc(w http.ResponseWriter, r *http.Request){
	CORS(w)
	currentRequest = r
	P = Response{ResponseCode: 0, ResponseMsg: "Success"}
	defer PrintReponse(w)

	if !connected{
		ThrowResponse("database_connection")
	}

	if val, ok := dofunc[r.Method]; ok {
		val()
	} else {
		ThrowResponse("routing_mistake")
	}
}

//Decodes response to ,,item"
func DecodeRequest(item interface{}) {
	
	decoder := json.NewDecoder(currentRequest.Body)
	err := decoder.Decode(&item)
	if err != nil {
		ThrowResponse("decode_failure")
	}

	defer currentRequest.Body.Close()
}

//Prints generated Response
func PrintReponse(w http.ResponseWriter) {
	result, err := json.MarshalIndent(P, "", "  ")
	if err != nil {
		ThrowResponse("parse_error")
		return
	}
	fmt.Fprintf(w, string(result))
	recover()
}

//Sets Response by ID (From Errors.go file)
func SetData(data interface{}) {
	P.ReturnData = data
	panic(nil)
}

//Find out if this might be possible
func ThrowResponse(ID string) {
	if _, ok := Responses[ID]; ok {
		P = Responses[ID]
		panic(nil)
	} else {
		panic("WRONG ERROR ID  - " + ID)
	}
}
