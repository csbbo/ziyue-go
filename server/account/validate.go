package account

type RegistParam struct {
	Username   string `json:"username" binding:"required"`
	Email		string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type LoginParam struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type UserSaveParm struct {
	Username string 	`json:"username"`
	Password string 	`json:"password"`
	Email string 		`json:"email"`
	Phone string		`json:"phone"`
	Remark string 		`json:"remark"`
	ImgPath string 		`json:"img_path"`
}