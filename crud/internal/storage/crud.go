package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertPost(mongo *DataBase, pd PostDocument) error {
	ctx := context.Background()
	_, err := mongo.postsCollection.InsertOne(ctx, pd)

	return err
}

func GetPost(mongo *DataBase, fs FieldSearch) ([]PostDocument, error) {
	ctx := context.Background()

	query := bson.M{
		fs.Field: fs.Value,
	}
	var output []PostDocument

	cursor, err := mongo.postsCollection.Find(ctx, query)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &output); err != nil {
		return nil, err
	}

	return output, nil
}

func GrepPosts(mongo *DataBase) ([]PostDocument, error) {
	ctx := context.Background()
	var output []PostDocument

	query := bson.M{"_id": bson.M{"$ne": nil}}

	cursor, err := mongo.postsCollection.Find(ctx, query)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &output); err != nil {
		return nil, err
	}

	return output, nil
}

func UpdatePost(mongo *DataBase, fu FieldUpdate) error {
	ctx := context.Background()

	filter := bson.M{"_id": fu.ID}
	query := bson.M{"$set": fu.NewValues}

	_, err := mongo.postsCollection.UpdateOne(ctx, filter, query)
	return err
}

func DeletePost(mongo *DataBase, sfs StrictFieldSearch) error {
	ctx := context.Background()

	filter := bson.M{"_id": sfs.ID}

	_, err := mongo.postsCollection.DeleteOne(ctx, filter)
	return err
}
