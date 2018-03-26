package core

type Response struct {
	ResponseCode int         `json:"code"`
	ResponseMsg  string      `json:"message,omitempty"`
	ReturnData   interface{} `json:"data,omitempty"`
}

var Responses = make(map[string]Response)

//Error responses used in WEB services
func loadResponses() {
	//System
	Responses["database_error"] = Response{ResponseCode: 200, ResponseMsg: "Failed to make database query"}
	Responses["decode_failure"] = Response{ResponseCode: 201, ResponseMsg: "Failed to decode Request data"}
	Responses["no_response"] = Response{ResponseCode: 202, ResponseMsg: "No Response Return"}

	//Request
	Responses["empty_fields"] = Response{ResponseCode: 220, ResponseMsg: "No Empty Fields"}

	//Session
	Responses["wrong_session"] = Response{ResponseCode: 240, ResponseMsg: "Wrong session ID"}

	//Projects module
	Responses["project_exists"] = Response{ResponseCode: 260, ResponseMsg: "Name Already Exists"}

	Responses["project_created"] = Response{ResponseCode: 0, ResponseMsg: "Project Created"}
	Responses["list_retrieved"] = Response{ResponseCode: 0, ResponseMsg: "List Retrieved"}

	//Users module
	Responses["username_exists"] = Response{ResponseCode: 280, ResponseMsg: "Username Already Exists"}
	Responses["email_exists"] = Response{ResponseCode: 281, ResponseMsg: "Email Already Exists"}
	Responses["password_match"] = Response{ResponseCode: 282, ResponseMsg: "Passwords do not match"}
	Responses["wrong_credentials"] = Response{ResponseCode: 283, ResponseMsg: "Wrong username or Password"}
	Responses["encryption_error"] = Response{ResponseCode: 284, ResponseMsg: "Failed to Encrypt password"}

	Responses["user_created"] = Response{ResponseCode: 0, ResponseMsg: "User Created"}
	Responses["logged_in"] = Response{ResponseCode: 0, ResponseMsg: "Logged In"}
}
