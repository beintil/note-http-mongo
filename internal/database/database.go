package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	login = "mongodb+srv://name:pass@cluster0.mgskbie.mongodb.net/?retryWrites=true&w=majority"
)

func ConnectDB() *mongo.Collection {
	log.Println("Database start...")

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(login))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database start OK\n", "")

	collection := client.Database("notesdb").Collection("Mynotes")

	return collection
}
