package account

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/common"
	"server/database"
	"time"
)

func regist(c *gin.Context) {
	form := &RegistParam{}
	if err := c.ShouldBind(form); err != nil {
		common.ResponseError(c, err)
		return
	}

	collection := database.MongoDB.Collection("user")

	count, _ := collection.CountDocuments(context.Background(), bson.M{"username": form.Username})
	if count != 0 {
		common.ResponseError(c, "用户已存在")
		return
	}

	count, _ = collection.CountDocuments(context.Background(), bson.M{"email": form.Email})
	if count != 0 {
		common.ResponseError(c, "邮箱已使用")
		return
	}

	timestamp := time.Now().Unix()
	user := User{
		Username: form.Username,
		Email:	  form.Email,
		Password: common.Hash(form.Password),
		CreateTime: timestamp,
		UpdateTime: timestamp,
	}
	res, _ := collection.InsertOne(context.Background(), user)
	common.ResponseSuccess(c, map[string]interface{}{"id": res.InsertedID})
}

func login(c *gin.Context) {
	form := &LoginParam{}
	if err := c.ShouldBind(form); err != nil {
		common.ResponseError(c, err)
		return
	}

	user := struct {
		Username string
		Password string
	}{}

	collection := database.MongoDB.Collection("user")
	filter := bson.M{"username": form.Username}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		common.ResponseError(c, "用户不存在")
		return
	}

	if user.Password != common.Hash(form.Password) {
		common.ResponseError(c, "密码错误")
		return
	}

	_, err = collection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"lastlogintime": time.Now().Unix()}})
	if err != nil {
		common.ServerError(c, "更新登录时间失败")
		return
	}

	session := sessions.Default(c)
	session.Set("user", form.Username)
	session.Save()
	common.ResponseSuccess(c, nil)
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	common.ResponseSuccess(c, nil)
}



func checkAuth(c *gin.Context) {
	if err := common.Check(c, nil, true); err != nil {
		common.ResponseError(c, err)
		return
	}
	common.ResponseSuccess(c, nil)
}

func user(c *gin.Context) {
	if err := common.Check(c, nil, true); err != nil {
		common.ResponseError(c, err)
		return
	}
	username := sessions.Default(c).Get("user")
	type User struct {
		Username string 	`json:"username"`
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
	var user User
	collection := database.MongoDB.Collection("user")

	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		common.ServerError(c, nil)
	}
	common.ResponseSuccess(c, user)
}

func userSave(c *gin.Context) {
	form := &UserSaveParm{}
	if err := common.Check(c, form, true); err != nil {
		common.ResponseError(c, err)
		return
	}

	update := bson.M{}
	if form.Username != "" {
		update["username"] = form.Username
	}
	if form.Password != "" {
		update["password"] = form.Password
	}
	if form.Email != "" {
		update["email"] = form.Email
	}
	if form.Phone != "" {
		update["phone"] = form.Phone
	}
	if form.Remark != "" {
		update["remark"] = form.Remark
	}
	if form.ImgPath != "" {
		update["imgpath"] = form.ImgPath
	}
	update = bson.M{"$set": update}

	username := sessions.Default(c).Get("user")
	collection := database.MongoDB.Collection("user")
	_, err := collection.UpdateOne(context.Background(), bson.M{"username": username}, update)
	if err != nil {
		common.ServerError(c, err.Error())
		return
	}
	common.ResponseSuccess(c, nil)
}

func Setup(r *gin.Engine) {
	r.POST("/api/regist", regist)
	r.POST("/api/login", login)
	r.POST("/api/logout", logout)
	r.GET("/api/checkAuth", checkAuth)
	r.GET("/api/user", user)
	r.POST("/api/user/save/", userSave)
}
