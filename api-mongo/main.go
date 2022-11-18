package main

import (
	"context"
	"fmt"

	db "github.com/paulokunde/api-mongo/DB"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	usersCollection := db.OpenConection().Collection("users")
	user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	result, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}
