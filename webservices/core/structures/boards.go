package structures
import (
	"gopkg.in/mgo.v2/bson"
)

type Board struct {
	ID bson.ObjectId        `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`
	Description string      `json:"description,omitempty" bson:"description,omitempty"`
	ProjectID bson.ObjectId `json:"project,omitempty" bson:"project"`
}

type BoardCreation struct{
	Board Board    `json:"board,omitempty"`
	Session Session       `json:"session,omitempty"`
}

type BoardListRequest struct{
	ProjectID bson.ObjectId       `json:"project"`
	Session Session    `json:"session"`
}