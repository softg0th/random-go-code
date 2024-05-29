package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConnection struct {
	client *mongo.Client
}

type DataBase struct {
	db              *mongo.Database
	postsCollection *mongo.Collection
}

func New() *mongoConnection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	return &mongoConnection{
		client: client,
	}
}

func InitDatabase(mongo *mongoConnection) (*DataBase, error) {
	datasDB := mongo.client.Database("crud")
	postsCollection := datasDB.Collection("posts")

	return &DataBase{
		db:              datasDB,
		postsCollection: postsCollection,
	}, nil
}

func GetDBInstance() *DataBase {
	db := &DataBase{}
	return db
}
