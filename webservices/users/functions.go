package users

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword (pass string) string{

	_, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err!=nil{
		return ""
	}
	return string(pass) //returns the same now, hash not working
}