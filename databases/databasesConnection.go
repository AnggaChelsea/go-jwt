package databases

import (
"fmt"
"log"
"time"
"os"
"context"
"github.com/joho/godotenv"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error when loadiung ")
	}
	MongoDb := os.Getenv("MONGODB_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal("error when creating")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	if err != nil {
		log.Fatal("error when connecting")
	}
	fmt.Println("Connected to mongodb")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) (*mongo.Collection, error) {
	var collection *mongo.Collection = client.Database("").Collection(collectionName)
	return collection, nil
}