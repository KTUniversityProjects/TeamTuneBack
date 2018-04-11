package main

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"../../core"
	"../../core/structures"
	"fmt"
	"time"
)


func deleteUser(userID bson.ObjectId) {

	core.Dao.C("users")
	err := core.Dao.Collection.Remove(bson.M{"_id":userID})
	if err != nil {
		fmt.Println("Error removing user")
		core.ThrowResponse("database_error")
	}

	//Visa kita gerai, ir turi veikti, bet jei pažiūrėsi projekto struktūroj, mes dar vienoj vietoj saugom UserID t.y. prie kiekvieno project'o yra users masyvas.
	//Tai iš to masyvo irgi reikia pašalinti tą vieną elementą, kur userID yra tas kurį čia turim.
	//Dar reikėtų pachekint ar tas userID nėra creator ir jei yra, tai reikia ir projektą trinti.

	//Reiškia turi sutikrint visus projektus. O jei gausis taip, kad trinsi projektą, turi ištrinti ir jo boards. O vėliau ir dar daugaiu visko reikės papildyti.
}

func deleteSessions(userID bson.ObjectId) {

	core.Dao.C("sessions")
	_ , err := core.Dao.Collection.RemoveAll(bson.M{"user":userID})
	if err != nil {
		fmt.Println("Error removing session")
		core.ThrowResponse("database_error")
	}
}

	func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//Check if correct username and password
func  checkCredentials(user structures.User) (bson.ObjectId) {
	core.Dao.C("users")

	var login = structures.User{}
	err := core.Dao.Collection.Find(bson.M{"username": user.Username}).Select(bson.M{"_id" :1, "password" : 1}).One(&login)
	if err != nil {
		fmt.Println("Wrong username")
		core.ThrowResponse("wrong_credentials")
	}

	if !CheckPasswordHash(user.Password, login.Password) {
		fmt.Println("Wrong password")
		core.ThrowResponse("wrong_credentials")
	}

	return login.Id
}

//Creates session / returns SessionID
func createSession(userID bson.ObjectId) {
	core.Dao.C("sessions")
	sessionID := bson.NewObjectId()

	//Create session object for insertion
	var session = structures.Session{
		SessionID:sessionID,UserID:userID,
		Expires:time.Now().Add(time.Duration(24 * time.Hour))}

	//Database insert
	err := core.Dao.Collection.Insert(&session)
	if err != nil {
		core.ThrowResponse("database_error")
	}

	core.SetData(sessionID)
}



//Adds User to Database
func addUser(user structures.User) {
	core.Dao.C("users")

	var err error
	user.Password,err = EncryptPassword(user.Password)
	if err != nil {
		core.ThrowResponse("encryption_error")
	}

	err = core.Dao.Collection.Insert(bson.M{"username":user.Username, "password":user.Password, "email":user.Email})
	if err != nil {
		core.ThrowResponse("database_error")
	}
}

//Checks if User and Email does not exists in Database
func checkFieldsExistance(user structures.User) {
	core.Dao.C("users")

	count, err := core.Dao.Collection.Find(bson.M{"username": user.Username}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}
	if count > 0 {
		core.ThrowResponse("username_exists")
	}

	count, err = core.Dao.Collection.Find(bson.M{"email": user.Email}).Count()
	if err != nil {
		core.ThrowResponse("database_error")
	}
	if count > 0 {
		core.ThrowResponse("email_exists")
	}
}

//Validates form fields
func validate(user structures.User) {

	if user.Username == "" || user.Password == "" || user.Email == "" {
		core.ThrowResponse("empty_fields")
	}

	if  user.Password2 != user.Password {
		core.ThrowResponse("password_match")
	}

	checkFieldsExistance(user)
}
