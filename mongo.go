package main

import (
	"fmt"
	"time"
	"log"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

type Attack struct {
	Method	string
	Payload string
	Time	time.Time
}

func MongoSend(method string, data string) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("test").Collection("testCollection")
	document := Attack{method, data, time.Now()}

	insertResult, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document: ", insertResult.InsertedID)
}

