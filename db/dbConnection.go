package db

import (
	"context"
	"log"

	"github.com/hospital-management/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database = utils.HMDB

func Connect() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("error occured during connecting to db", err)
	}

	return client.Database(Database)
}
