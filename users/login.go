package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Payload struct {
	Response map[string]string
	ReturnCode int
}

func Login(w http.ResponseWriter, r *http.Request)(bool){

	return true
}

func main() {
	http.HandleFunc("/login", ServRest)
	http.ListenAndServe("localhost:1338", nil)
	fmt.Scanln()
	os.Exit(0)
}

func ServRest(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()                     // Parses the request body
	username := r.Form.Get("login") // x will be "" if parameter is not set
	password := r.Form.Get("password") // x will be "" if parameter is not set
	fmt.Println(username)
	fmt.Println(password)

	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(response))
}

func getJsonResponse()([]byte, error){
	data := make(map[string]string)
	data["username"] = "threatx"

	p := Payload{data, 0}

	return json.MarshalIndent(p, "", "  ")
}