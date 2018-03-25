package core

type Response struct{
	ResponseCode int
	ResponseMsg string
}

var Responses = make(map[string]Response)

func loadReponses(){
	Responses["decode_failure"] = Response{200, "Failed to decode Request data"}

	Responses["database_error"] = Response{300, "Failed to make database query"}

	Responses["no_response"] = Response{500, "No Response Return"}

	Responses["username_exists"] = Response{400, "Username Already Exists"}
	Responses["email_exists"] = Response{401, "Email Already Exists"}

	Responses["user_created"] = Response{0, "User Created"}
}


