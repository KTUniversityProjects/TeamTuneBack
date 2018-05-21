package core

type Response struct {
	ResponseCode int         `json:"code"`
	ResponseMsg  string      `json:"message,omitempty"`
	ReturnData   interface{} `json:"data,omitempty"`
}

type ResponseTest struct {
	ResponseCode int         		`json:"code"`
	ResponseMsg  string      		`json:"message,omitempty"`
	ReturnData   map[string]string 	`json:"data,omitempty"`
}

var Responses = make(map[string]Response)

//Error responses used in WEB services
func loadResponses() {

	//System
	Responses["database_error"] = Response{ResponseCode: 200, ResponseMsg: "Failed to make database query"}
	Responses["decode_failure"] = Response{ResponseCode: 201, ResponseMsg: "Failed to decode Request data"}
	Responses["system_mistake"] = Response{ResponseCode: 202, ResponseMsg: "Mistake in Web Service"}
	Responses["routing_mistake"] = Response{ResponseCode: 203, ResponseMsg: "Routing mistake in Web Service"}
	Responses["database_connection"] = Response{ResponseCode: 204, ResponseMsg: "Not connected to database"}

	//Request & validation
	Responses["empty_fields"] = Response{ResponseCode: 220, ResponseMsg: "No Empty Fields"}
	Responses["name_exists"] = Response{ResponseCode: 221, ResponseMsg: "Name Already Exists"}
	Responses["no_permission"] = Response{ResponseCode: 222, ResponseMsg: "No permissions"}

	//Session
	Responses["wrong_session"] = Response{ResponseCode: 240, ResponseMsg: "Wrong session ID"}

	//Projects
	Responses["project_not_exists"] = Response{ResponseCode: 250, ResponseMsg: "Project does not exists"}

	//Users
	Responses["username_exists"] = Response{ResponseCode: 260, ResponseMsg: "Username Already Exists"}
	Responses["email_exists"] = Response{ResponseCode: 261, ResponseMsg: "Email Already Exists"}
	Responses["password_match"] = Response{ResponseCode: 262, ResponseMsg: "Passwords do not match"}
	Responses["wrong_credentials"] = Response{ResponseCode: 263, ResponseMsg: "Wrong username or Password"}
	Responses["encryption_error"] = Response{ResponseCode: 264, ResponseMsg: "Failed to Encrypt password"}

	//Boards
	Responses["boards_not_exists"] = Response{ResponseCode: 280, ResponseMsg: "Wrong Project ID"}
}
