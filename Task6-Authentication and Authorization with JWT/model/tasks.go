package model

type Task struct {
	ID          string `bson:"id" json:"id"`
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Status      string `bson:"status" json:"status"`
	CreaterID   string `bson:"creater_id" json:"creater_id"`
}
