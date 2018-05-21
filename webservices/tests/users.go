package main

import "fmt"

func usersTest()(string, string, string, string){
	registerTest()
	return loginTest()
}

func registerTest(){
	usersRequest("{\"username\":\"threat\",\"password\":\"slaptazodis\"}",
		[]int{220}, "PUT")
	usersRequest("{\"username\":\"threatx\",\"password\":\"slaptazodis\",\"email\":\"viliusx@rinkodara.lt\",\"password2\":\"slaptazodis\"}",
		[]int{0, 260, 261}, "PUT")
	usersRequest("{\"username\":\"threat\",\"password\":\"slaptazodis\",\"email\":\"viliusx2@rinkodara.lt\",\"password2\":\"slaptazodis\"}",
		[]int{0, 260, 261}, "PUT")

	usersRequest("{\"username\":\"threatxz\",\"password\":\"slaptazodis\",\"email\":\"viliusx@rinkodara.lt\",\"password2\":\"slaptazodis\"}",
		[]int{261}, "PUT")
	usersRequest("{\"username\":\"threatx\",\"password\":\"slaptazodis\",\"email\":\"viliusxz@rinkodara.lt\",\"password2\":\"slaptazodis\"}",
		[]int{260}, "PUT")

	fmt.Println("Register test passed")
}

func loginTest()(string, string, string, string){
	data := usersRequest("{\"username\":\"threatx\",\"password\":\"slaptazodis\"}",
		[]int{0}, "POST").(map[string]interface{})
	data2 := usersRequest("{\"username\":\"threat\",\"password\":\"slaptazodis\"}",
		[]int{0}, "POST").(map[string]interface{})
	usersRequest("{\"username\":\"threat\",\"password\":\"slaptazodis2\"}",
		[]int{263}, "POST")

	usersRequest("{\"username\":\"....\"}",
		[]int{263}, "POST")
	usersRequest("{\"username\":\"....\",\"password\":\"....\"}",
		[]int{263}, "POST")

	fmt.Println("Login test passed")
	return data["id"].(string), data["user"].(string), data2["id"].(string), data2["user"].(string)
}

func usersRequest(data string, expectedResponse []int, reqType string)(interface{}) {
	return makeRequest(data, expectedResponse, "1339", reqType)
}