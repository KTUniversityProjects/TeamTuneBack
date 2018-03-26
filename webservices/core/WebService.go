package core

import (
	"net/http"
	"encoding/json"
	"fmt"
)

var P = Response{ResponseCode: 201, ResponseMsg: "No Response Returned"}

//Adds CORS header to response Writer
func CORS(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

//Decodes response to ,,item"
func DecodeRequest(item interface{}, r *http.Request) bool {

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&item)
	if err != nil {
		SetResponse("decode_failure")
		return false
	}

	defer r.Body.Close()
	return true
}

//Prints generated Response
func PrintReponse(w http.ResponseWriter) {
	result, err := json.MarshalIndent(P, "", "  ")
	if err != nil {
		SetResponse("parse_error")
	}
	fmt.Fprintf(w, string(result))
}

//Sets Response by ID (From Errors.go file)
func SetResponse(ID string) {
	if len(Responses) == 0 {
		loadResponses()
	}
	if _, ok := Responses[ID]; ok {
		P = Responses[ID]
	} else {
		panic("WRONG ERROR ID  - " + ID)
	}
}

//Sets Response by ID (From Errors.go file)
func SetData(data interface{}) {
	P.ReturnData = data
}

//Find out if this might be possible
func ThrowResponse(ErrorID string, w http.ResponseWriter) {
	SetResponse(ErrorID)
	PrintReponse(w)
}
