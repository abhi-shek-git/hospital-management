package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOneByMobileNo(collectionName *mongo.Collection, idMobileNo int) error {
	query := bson.M{"mobileno": idMobileNo}
	dbFindResult := collectionName.FindOne(context.TODO(), query)
	err := dbFindResult.Err()
	if err == mongo.ErrNoDocuments {
		log.Printf("no document found in db")
		return mongo.ErrNoDocuments
	}
	if err != nil && err != mongo.ErrNoDocuments {
		log.Printf("error occured during finding data from db, error = %s", err)
		return err
	}
	return nil
}

func InsertOne(collection *mongo.Collection, insertData interface{}) error {

	_, err := collection.InsertOne(context.TODO(), insertData)
	if err != nil {
		log.Printf("error occured during inserting the data into db %s", err)
		return err
	}
	return nil
}

func FindOneAndDelete(collection *mongo.Collection, idMobileNo int) error {
	query := bson.M{"mobileno": idMobileNo}
	findResult := collection.FindOneAndDelete(context.TODO(), query)
	err := findResult.Err()

	if err != nil {
		log.Printf("error occured during finding data from db. Error=%s", err)
		return err
	}

	return nil
}
