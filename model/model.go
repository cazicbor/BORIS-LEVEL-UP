package model

type Task struct {
	ID          interface{} `json:"id" bson:"_id"`
	Description string      `json:"description" bson:"description"`
	Deadline    string      `json:"deadline" bson:"deadline"`
	Status      string      `json:"status" bson:"status"`
}
