package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	ErrorMessage string
	ReturnCode int
}

func main() {
	http.HandleFunc("/", Project)
	http.ListenAndServe("localhost:1339", nil)
}

func Project(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()                     // Parses the request body

	UserId := r.Form.Get("UserID") // x will be "" if parameter is not set

	ErrorMsg := "Galima kurti projekta"
	ErrorCode := 0

	if(UserId != "1337" ) {
		ErrorMsg = "Vartotojas negali kurti projekto"
		ErrorCode = 1
	}

	p := Payload{ErrorMsg, ErrorCode}


	json, err := json.MarshalIndent(p, "", "  ")
	if err != nil{
		panic(nil)
	}
	fmt.Fprintf(w, string(json))

}