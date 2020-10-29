package article

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

func articleCreate(c *gin.Context)  {
	data := &ArticleCreateParm{}
	if err := common.Check(c, data, true); err != nil {
		common.ResponseError(c, err)
		return
	}

	username := sessions.Default(c).Get("user")
	user := struct {
		ID primitive.ObjectID	`bson:"_id"`
	}{}
	err := database.MongoDB.Collection("user").FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		common.ResponseError(c, "用户不存在")
		return
	}

	timestamp := time.Now().Unix()
	article := Article{
		Title:	data.Title,
		Content:	data.Content,
		Creator:	user.ID,
		CreateTime: timestamp,
		UpdateTime: timestamp,
	}
	collection := database.MongoDB.Collection("article")
	res, _ := collection.InsertOne(context.Background(), article)
	common.ResponseSuccess(c, map[string]interface{}{"id": res.InsertedID})
}

func articleList(c *gin.Context) {
	data := &ArticleListParm{}
	if err := common.Check(c, data, true); err != nil {
		common.ResponseError(c, err)
		return
	}

	filter := bson.M{}
	if data.Creator != "" {
		user := struct {
			ID	primitive.ObjectID	`bson:"_id"`
		}
		collection := database.MongoDB.Collection("user")
		collection.FindOne(context.Background(), bson.M{"username": data.Creator}).Decode(&user)
		filter["creator"] = user.ID
	}

	collection := database.MongoDB.Collection("article")
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		common.ResponseError(c, "查询异常")
		return
	}
	var article Article
	var results []Article
	for cur.Next(context.Background()) {
		err = cur.Decode(&article)
		if err != nil {
			common.ResponseError(c, "解析失败")
			return
		}
		results = append(results, article)
	}
	common.ResponseSuccess(c, results)
}
func Setup(r *gin.Engine) {
	r.POST("/api/article/create", articleCreate)
	r.GET("/api/article/list", articleList)
}
