package core

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Decodable interface{

}

type Payload struct {
	ErrorMessage string
	ReturnCode   int
}

var p = Response{201, "No Response Returned"}

func CORS(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func DecodeRequest(t Decodable, r *http.Request) bool {

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		SetReponse("decode_failure")
		return false
	}

	defer r.Body.Close()
	return true
}

func PrintReponse(w http.ResponseWriter) {
	json, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		SetReponse("parse_error")
	}
	fmt.Fprintf(w, string(json))
}

func SetReponse(ErrorID string) {
	if len(Responses) == 0 {
		loadReponses()
	}
	if _, ok := Responses[ErrorID]; ok {
		p = Responses[ErrorID]
	} else {
		panic("WRONG ERROR ID  - " + ErrorID)
	}
}

/* Try to find way to Close Response on PrintResponse(w http.ResponseWriter) method and then use this method for error handling

func ThrowResponse(ErrorID string, w http.ResponseWriter){
	SetReponse(ErrorID)
	PrintReponse(w)
}
*/