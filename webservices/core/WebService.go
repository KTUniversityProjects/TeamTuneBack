package core

import (
	"net/http"
	"encoding/json"
	"fmt"
)

var P Response

//Adds CORS header to response Writer
func CORS(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

var dofunc func()
var currentRequest *http.Request

func Initialize(function func(), port string){
	loadResponses()
	dofunc = function
	Dao.Connect(Config.DatabaseHost + ":" + Config.DatabasePort, Config.DatabaseName)
	http.HandleFunc("/", requestFunc)
	http.ListenAndServe(Config.Host + ":" + port, nil)
}

func requestFunc(w http.ResponseWriter, r *http.Request){
	CORS(w)
	currentRequest = r
	P = Response{ResponseCode: 0, ResponseMsg: "Success"}
	defer PrintReponse(w)
	dofunc()
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
