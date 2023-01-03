package MongoDB

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

// ConnectToDatabase
// Connects to a given database
func ConnectToDatabase() (*mongo.Client, context.Context) {
	var databaseUri = os.Getenv("DATABASE_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(databaseUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx
}

func GetAllThreads() ([]*Thread, error) {
	client, ctx := ConnectToDatabase()
	var databaseName = os.Getenv("DATABASE_NAME")
	var databaseCollection = os.Getenv("DATABASE_COLLECTION")
	var threads []*Thread

	database := client.Database(databaseName)
	collection := database.Collection(databaseCollection)
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var user Thread
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		threads = append(threads, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	err = cursor.Close(ctx)
	if err != nil {
		return nil, err
	}
	if len(threads) == 0 {
		return nil, errors.New("no threads found")
	}
	return threads, nil
}

func CreateThread(thread *Thread) error {
	client, ctx := ConnectToDatabase()
	var databaseName = os.Getenv("DATABASE_NAME")
	var databaseCollection = os.Getenv("DATABASE_COLLECTION")

	database := client.Database(databaseName)
	collection := database.Collection(databaseCollection)
	_, err := collection.InsertOne(ctx, thread)
	return err
}

func GetThread(id *string) (Thread, error) {
	var thread Thread
	client, ctx := ConnectToDatabase()
	var databaseName = os.Getenv("DATABASE_NAME")
	var databaseCollection = os.Getenv("DATABASE_COLLECTION")

	database := client.Database(databaseName)
	collection := database.Collection(databaseCollection)
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := collection.FindOne(ctx, query).Decode(&thread)
	return thread, err
}

func UpdateThread(thread *Thread) error {
	client, ctx := ConnectToDatabase()
	var databaseName = os.Getenv("DATABASE_NAME")
	var databaseCollection = os.Getenv("DATABASE_COLLECTION")

	database := client.Database(databaseName)
	collection := database.Collection(databaseCollection)
	filter := bson.D{bson.E{Key: "Frage", Value: thread.Frage}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "Frage", Value: thread.Frage}, bson.E{Key: "Titel", Value: thread.Titel}}}}
	result, _ := collection.UpdateOne(ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched thread found for update")
	}
	return nil
}
