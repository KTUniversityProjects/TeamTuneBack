package create

import (
	"net/http"
	"../../../core"
	_ "gopkg.in/mgo.v2/bson"
	"../../"
	_ "fmt"
	_ "golang.org/x/crypto/bcrypt"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}
var Database = ServiceDatabase{&core.Dao}

//Connects to database and listens to port
func main() {
	Database.Dao.Connect("localhost:27017", "teamtune")
	http.HandleFunc("/", do)
	http.ListenAndServe("localhost:1338", nil)
}

func do(w http.ResponseWriter, r *http.Request) {
	core.CORS(w)

	//Parses request data to
	var data users.UserStruct
	if core.DecodeRequest(&data, r){
		//Checks Username and Email
		if userID := Database.checkCredentials(data); userID != "" {
			Database.CreateSession(data, userID) //Logs in
		}
	}

	//Prints R
	core.PrintReponse(w)
}
