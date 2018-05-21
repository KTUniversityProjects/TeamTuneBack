package main

func usersTest()(string, string, string, string){
	registerTest()
	return loginTest()
}

func registerTest(){
	usersRequest("{\"username\":\"test\",\"password\":\"test\"}",
		[]int{220}, "PUT")
	usersRequest("{\"username\":\"test\",\"password\":\"test\",\"email\":\"test@test.lt\",\"password2\":\"test\"}",
		[]int{0, 260, 261}, "PUT")
	usersRequest("{\"username\":\"test2\",\"password\":\"test2\",\"email\":\"test2@test2.lt\",\"password2\":\"test2\"}",
		[]int{0, 260, 261}, "PUT")

	usersRequest("{\"username\":\"test2\",\"password\":\"test2\",\"email\":\"test2@ccc.lt\",\"password2\":\"test2\"}",
		[]int{260}, "PUT")
	usersRequest("{\"username\":\"dsadsadsa\",\"password\":\"test2\",\"email\":\"test2@test2.lt\",\"password2\":\"test2\"}",
		[]int{261}, "PUT")
	usersRequest("{\"username\":\"dsadsadsa\",\"password\":\"tesxt2\",\"email\":\"dsadsadsa@test2.lt\",\"password2\":\"test2\"}",
		[]int{262}, "PUT")
}

func loginTest()(string, string, string, string){
	data := usersRequest("{\"username\":\"test\",\"password\":\"test\"}",
		[]int{0}, "POST").(map[string]interface{})
	data2 := usersRequest("{\"username\":\"test2\",\"password\":\"test2\"}",
		[]int{0}, "POST").(map[string]interface{})
	usersRequest("{\"username\":\"threat\",\"password\":\"slaptazodis2\"}",
		[]int{263}, "POST")

	usersRequest("{\"username\":\"....\"}",
		[]int{263}, "POST")
	usersRequest("{\"username\":\"....\",\"password\":\"....\"}",
		[]int{263}, "POST")

	return data["id"].(string), data["user"].(string), data2["id"].(string), data2["user"].(string)
}

func usersRequest(data string, expectedResponse []int, reqType string)(interface{}) {
	return makeRequest(data, expectedResponse, "1339", reqType)
}