package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	MongoDB *mongo.Database
)

func InitMongoDB() error {
	opt := options.Client().ApplyURI(MongoConnectURI)
	opt.SetMaxPoolSize(MongoMaxPoolSize)

	client, err := mongo.NewClient(opt)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), MongoConnectTimeOut*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	MongoDB = client.Database(MongoDatabase)
	return nil
}
