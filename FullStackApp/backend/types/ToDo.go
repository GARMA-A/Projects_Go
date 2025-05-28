package types

type ToDo struct {
	ID        string `json:"id" bson:"_id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}
