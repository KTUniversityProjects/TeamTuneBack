package core

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func AddRouting(requestType string, function func()){
	dofunc[requestType] = function
}

func Initialize(port string){
	AddRouting("OPTIONS", func(){})
	loadResponses()
	if Exists("developer") {
		Dao.Connect(ConfigDev.DatabaseHost + ":" + ConfigDev.DatabasePort, ConfigDev.DatabaseName)
	} else {
		Dao.Connect(Config.DatabaseHost + ":" + Config.DatabasePort, Config.DatabaseName)
	}
	http.HandleFunc("/", requestFunc)
	http.ListenAndServe(Config.Host + ":" + port, nil)
}

var P Response

//Adds CORS header to response Writer
func CORS(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
}

var dofunc = make(map[string]func())
var currentRequest *http.Request

func requestFunc(w http.ResponseWriter, r *http.Request){
	CORS(w)
	currentRequest = r
	P = Response{ResponseCode: 0, ResponseMsg: "Success"}
	defer PrintReponse(w)

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
