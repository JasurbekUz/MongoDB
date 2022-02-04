package main

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}

func DB() *mongo.Database {
	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	db := client.Database("users")
	fmt.Println("Successfuly connected to the database.")

	return db

	
}

func Find(email string) (*User, error) {

	db := DB()
	var user User
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := usersC.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func main () {

	var email string

	fmt.Scanln(email)

	user, _ := Find("2smvkm")

	fmt.Println(user.ID, "\n", user.Name)
}