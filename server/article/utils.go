package article

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	Title string 	`json:"title"`
	Content string 	`json:"content"`
	Creator primitive.ObjectID	`json:"creator"`
	CreateTime int64 	`json:"create_time"`
	UpdateTime int64 	`json:"update_time"`
}
