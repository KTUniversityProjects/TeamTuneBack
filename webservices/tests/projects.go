package main

func createTest(sessionID string, userID string)(string){

	data := projectsRequest("{\"session\": {\"id\": \"" + sessionID + "\",\"user\": \"" + userID + "\"},\"project\": {\"name\": \"Projektas2\",\"description\": \"Čia yra projektas 1 sukūriau jį pasitestavimui\"}}",
		[]int{0}, "PUT")

	return data.(string)
}

func getAndDeleteTest(sessionID string, userID string,sessionID2 string, userID2 string){

	projectsRequest("{\"session\": {\"id\": \"" + sessionID + "\",\"user\": \"" + userID + "\"},\"project\": {\"name\": \"Projektas2\",\"description\": \"Čia yra projektas 1 sukūriau jį pasitestavimui\"}}",
		[]int{0, 221}, "PUT")

	projectsRequest("{\"session\": {\"id\": \"" + sessionID + "\",\"user\": \"" + userID + "\"},\"project\": {\"name\": \"Projektas2\",\"description\": \"Čia yra projektas 1 sukūriau jį pasitestavimui\"}}",
		[]int{0, 221}, "PUT")

	data := projectsRequest("{\"session\": {\"id\": \"" + sessionID + "\",\"user\": \"" + userID + "\"}}",
		[]int{0}, "POST").([]interface{})

	data2 := projectsRequest("{\"session\": {\"id\": \"" + sessionID2 + "\",\"user\": \"" + userID2 + "\"}}",
		[]int{0}, "POST").([]interface{})

	for i:=0; i < len(data); i++{
		projectsRequest("{\"session\": {\"id\": \"" + sessionID + "\",\"user\": \"" + userID + "\"},\"project\": {\"id\":\"" + data[i].(map[string]interface{})["id"].(string) + "\"}}",
			[]int{0}, "DELETE")
	}

	for i:=0; i < len(data2); i++{
		projectsRequest("{\"session\": {\"id\": \"" + sessionID + "\",\"user\": \"" + userID + "\"},\"project\": {\"id\":\"" + data2[i].(map[string]interface{})["id"].(string) + "\"}}",
			[]int{222}, "DELETE")
	}

	for i:=0; i < len(data2); i++{
		projectsRequest("{\"session\": {\"id\": \"" + sessionID2 + "\",\"user\": \"" + userID2 + "\"},\"project\": {\"id\":\"" + data2[i].(map[string]interface{})["id"].(string) + "\"}}",
			[]int{0}, "DELETE")
	}
}

func projectsTest(sessionID string, userID string,sessionID2 string, userID2 string)(string, string){

	getAndDeleteTest(sessionID, userID, sessionID2, userID2)

	projectID1 := createTest(sessionID, userID)
	projectID2 := createTest(sessionID2, userID2)
	return projectID1, projectID2
}

func projectsRequest(data string, expectedResponse []int, reqType string)(interface{}) {
	return makeRequest(data, expectedResponse, "http://192.168.99.100:1338", reqType)
}