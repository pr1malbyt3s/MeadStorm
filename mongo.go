package main

import (
	"fmt"
	"time"
	"log"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

type ContactAttack struct {
	Method	string
    Name string
    Email string
    Message string
	Time time.Time
}

type LoginAttack struct {
    Method string
    UserName string
    Password string
    Time time.Time
}

func MongoSendContact(method string, name string, email string, message string) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("test").Collection("testCollection")
	document := ContactAttack{method, name, email, message, time.Now()}

	insertResult, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document: ", insertResult.InsertedID)
}

func MongoSendLogin(method string, username string, password string) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("test").Collection("testCollection")
	document := LoginAttack{method, username, password, time.Now()}

	insertResult, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document: ", insertResult.InsertedID)
}
