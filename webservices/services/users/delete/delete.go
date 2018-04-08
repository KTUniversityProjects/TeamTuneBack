package main

import (
	"../../../core"
	"gopkg.in/mgo.v2/bson"
	"../../projects"
)

var servicePort = "1340"

func do() {

	// čia turi būti ne PojectRequest, o User ar kažkas panašaus t.y. į kokio tipo struktūra bus suparsintas tavo siūstas request. ž
	// Realiai, kad ištrinti user'į tai tau reikia tik USER ID, bet user ID mes tiesiogiai nesiunčiam, siunčiam sesijos ID. Tai tau čia reik sesijos struktūrą nurodyt
	//O ji jau sukurta core/structures/structures.go, tai gali ir įrašyt čia structures.Session
	var data projects.ProjectRequest

	core.DecodeRequest(&data)

	//CreateSession(data, bson.ObjectId(userID)) //Sesija kuriama tik loginantis. Kad gaut sesijos ID. Tai šito nereikia.

	//Va čia gausi user id
	userID := core.Dao.CheckSession(data.Session)
	deletesUser(bson.ObjectId((userID))) //USer ID jau grįžta kaip bson.ObjectId, tai nereikia čia, užtenka userID parašyt
}

func deletesUser(userID bson.ObjectId) {

	//šit vieta bus gerai, tik
	err := core.Dao.Collection.Remove(bson.M{"user":bson.ObjectId(userID)})// vėl rašai bson.ObjectId, nors userID jau yra bson.ObjectId tipo, tai čia vėl užtenka userID tik parašyt

	//čia skliaustų nereikia
	if(err != nil){
		core.ThrowResponse("database_error")
	}

	//Visa kita gerai, ir turi veikti, bet jei pažiūrėsi projekto struktūroj, mes dar vienoj vietoj saugom UserID t.y. prie kiekvieno project'o yra users masyvas.
	//Tai iš to masyvo irgi reikia pašalinti tą vieną elementą, kur userID yra tas kurį čia turim.
	//Dar reikėtų pachekint ar tas userID nėra creator ir jei yra, tai reikia ir projektą trinti.

	//Reiškia turi sutikrint visus projektus. O jei gausis taip, kad trinsi projektą, turi ištrinti ir jo boards. O vėliau ir dar daugaiu visko reikės papildyti.
}

/*           Every Webservice             */
func main(){
	core.Initialize(do, servicePort)
}