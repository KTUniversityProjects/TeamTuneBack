package core

type Response struct {
	ResponseCode int         `json:"code"`
	ResponseMsg  string      `json:"message,omitempty"`
	ReturnData       interface{} `json:"data,omitempty"`
}

var Responses = make(map[string]Response)

//Error responses used in WEB services
func loadResponses() {
	//System
	Responses["database_error"] = Response{ResponseCode: 300, ResponseMsg: "Failed to make database query"}
	Responses["decode_failure"] = Response{ResponseCode: 200, ResponseMsg: "Failed to decode Request data"}
	Responses["no_response"] = Response{ResponseCode: 500, ResponseMsg: "No Response Return"}

	//Request
	Responses["empty_fields"] = Response{ResponseCode: 402, ResponseMsg: "No Empty Fields"}

	//Session
	Responses["wrong_session"] = Response{ResponseCode: 300, ResponseMsg: "Wrong session ID"}

	//Projects module
	Responses["project_exists"] = Response{ResponseCode: 401, ResponseMsg: "Name Already Exists"}

	Responses["project_created"] = Response{ResponseCode: 0, ResponseMsg: "Project Created"}

	//Users module
	Responses["username_exists"] = Response{ResponseCode: 400, ResponseMsg: "Username Already Exists"}
	Responses["email_exists"] = Response{ResponseCode: 401, ResponseMsg: "Email Already Exists"}
	Responses["password_match"] = Response{ResponseCode: 403, ResponseMsg: "Passwords do not match"}
	Responses["wrong_credentials"] = Response{ResponseCode: 403, ResponseMsg: "Wrong username or Password"}

	Responses["user_created"] = Response{ResponseCode: 0, ResponseMsg: "User Created"}
	Responses["logged_in"] = Response{ResponseCode: 0, ResponseMsg: "Logged In"}
}
