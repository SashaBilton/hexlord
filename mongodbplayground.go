package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect() {
	//write hello world to the screen
	fmt.Println("Hello Mongo?")

	//To start mongodb/brew/mongodb-community now and restart at login:
	//brew services start mongodb/brew/mongodb-community
	//brew services stop mongodb-community

	//connect to a mongodb database
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//clientOptions.SetConnectTimeout(30)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

}
