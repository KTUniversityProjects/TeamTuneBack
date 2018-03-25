package core

import (
	"net/http"
	"encoding/json"
	"fmt"
)

//Interface for Json body decoding
type Decodable interface{}

var p = Response{201, "No Response Returned"}

//Adds CORS header to response Writer
func CORS(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

//Decodes response to ,,item"
func DecodeRequest(item Decodable, r *http.Request) bool {

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&item)
	if err != nil {
		SetReponse("decode_failure")
		return false
	}

	defer r.Body.Close()
	return true
}

//Prints generated Response
func PrintReponse(w http.ResponseWriter) {
	json, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		SetReponse("parse_error")
	}
	fmt.Fprintf(w, string(json))
}

//Sets Response by ID (From Errors.go file)
func SetReponse(ID string) {
	if len(Responses) == 0 {
		loadReponses()
	}
	if _, ok := Responses[ID]; ok {
		p = Responses[ID]
	} else {
		panic("WRONG ERROR ID  - " + ID)
	}
}

/* Try to find way to Close Response on PrintResponse(w http.ResponseWriter) method and then use this method for error handling

func ThrowResponse(ErrorID string, w http.ResponseWriter){
	SetReponse(ErrorID)
	PrintReponse(w)
}
*/