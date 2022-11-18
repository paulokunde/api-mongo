package db

import (
	"context"
	"fmt"
	"log"

	configs "github.com/paulokunde/api-mongo/Configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func OpenConection() (*mongo.Client, error) {
	conf := configs.GetDB()
	//const uri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)
	log.Printf("Formação da uri %v", uri)

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	//client.Database("infra")
	return client, err
}
