package projects

type Project struct {
	Id string        `json:"id,omitempty" bson:"_id,omitempty"`
	Name string  `json:"name,omitempty" bson:"name,omitempty"`
	Description string  `json:"description,omitempty" bson:"description,omitempty"`
	User string `json:"user,omitempty" bson:"user,omitempty"`
}

type ProjectCreation struct{
	Project Project    `json:"project,omitempty"`
	SessionID string       `json:"session,omitempty"`
}