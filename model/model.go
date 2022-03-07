package model

type Task struct {
	ID          interface{} `json:"id"`
	Description string      `json:"description"`
	Deadline    string      `json:"deadline"`
	Status      string      `json:"status"`
}
