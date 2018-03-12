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
	http.HandleFunc("/", Login)
	http.ListenAndServe("localhost:1338", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()                     // Parses the request body

	username := r.Form.Get("login") // x will be "" if parameter is not set
	password := r.Form.Get("password") // x will be "" if parameter is not set



	ErrorMsg := "Prisijungta Sėkmingai"
	ErrorCode := 0

	if(username != "vilius" || password != "slaptazodis") {
		ErrorMsg = "Blogi prisijungimo duomenys"
		ErrorCode = 1
	}

	p := Payload{ErrorMsg, ErrorCode}


	json, err := json.MarshalIndent(p, "", "  ")
	if err == nil{
		panic(nil)
	}
	fmt.Fprintf(w, string(json))
}