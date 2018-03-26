package users

type UserStruct struct {
	Id string        `json:"id" bson:"_id,omitempty"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Password2 string `json:"password2"`
	Email    string  `json:"email"`
}