package main

import (
	"net/http"
	"encoding/json"
	"core"
	"fmt"
	"bytes"
)

func main() {
	sessionID1,userID1,sessionID2,userID2 := usersTest()
	fmt.Println("SESSIONID1 - " + sessionID1)
	fmt.Println("USERID1 - " + userID1)
	fmt.Println("SESSIONID2 - " + sessionID2)
	fmt.Println("USERID2 - " + userID2)
	projectID1, projectID2 := projectsTest(sessionID1, userID1,sessionID2,userID2)
	fmt.Println("PROJECTID1 - " + projectID1)
	fmt.Println("PROJECTID2 - " + projectID2)
}

func makeRequest(data string, expectedResponse []int, url string, reqType string)(interface{}){

	var jsonStr = []byte(data)
	req, err := http.NewRequest(reqType, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	var item core.Response

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&item)
	if err != nil {
		fmt.Println(err)
		testError(data, item, expectedResponse)
	}

	success := false
	for j := 0; j < len(expectedResponse); j++ {
		if !success && item.ResponseCode == expectedResponse[j]{
			success = true
		}
	}

	if !success{
		testError(data, item, expectedResponse)
	}

	return item.ReturnData
}

func testError(request string, response core.Response, expectedResponse []int) {
	fmt.Println("Request")
	fmt.Println(request)
	fmt.Print("Expected Response - ")
	fmt.Println(expectedResponse)
	fmt.Print("Response - ")
	fmt.Println(response.ResponseCode)
	fmt.Println(response.ResponseMsg)
	panic("")
}



