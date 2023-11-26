package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOneByMobileNo(collectionName *mongo.Collection, idMobileNo int) string {
	query := bson.M{"mobileno": idMobileNo}
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

func InsertOne (collection *mongo.Collection, insertData interface{}) string {

	_, err := collection.InsertOne(context.TODO(), insertData)
	if err != nil {
		log.Printf("error occured during inserting the data into db %s", err)
		return "not inserted"
	}
	return "inserted"
}
