package db

import (
	"context"
	"log"

	"github.com/hospital-management/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database = utils.HMDB

func ConnectDB() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("error occured during connecting to db", err)
	}

	return client.Database(Database)
}

func FindDb(collectionName *mongo.Collection, query bson.M) string {
	dbFindResult := collectionName.FindOne(context.TODO(), query)
	err := dbFindResult.Err()
	if err == nil {
		log.Printf("data already exists")
		return "data already exists"
	}
	if err != nil && err != mongo.ErrNoDocuments {
		log.Printf("error occured during finding data from db, error = %s", err)
		return "internal error"
	}
	return "found"

}

func InsertDb(collection *mongo.Collection, insertData interface{}) string {

	_, err := collection.InsertOne(context.TODO(), insertData)
	if err != nil {
		log.Printf("error occured during inserting the data into db %s", err)
		return "not inserted"
	}
	return "inserted"
}
