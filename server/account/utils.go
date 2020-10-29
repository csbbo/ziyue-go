package account

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
	Username string 	`json:"username"`
	Password string 	`json:"password"`
	Email string 		`json:"email"`
	Phone string		`json:"phone"`
	Remark string 		`json:"remark"`
	ImgPath string 		`json:"img_path"`

	Focus	[]primitive.ObjectID	`json:"focus"`
	Fans	[]primitive.ObjectID	`json:"fans"`
	article	[]primitive.ObjectID	`json:"article"`
	draft	[]primitive.ObjectID	`json:"draft"`
	words	int64		`json:"words"`
	assets	int64		`json:"assets"`

	CreateTime int64 	`json:"create_time"`
	UpdateTime int64 	`json:"update_time"`
	LastLoginTime int64 `json:"last_login_time"`
}
