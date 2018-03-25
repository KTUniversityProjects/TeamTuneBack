package users

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword (pass string) string{

	pass2, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err!=nil{
		return ""
	}
	return string(pass2)
}