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

func FindAndDelete(collection *mongo.Collection, idMobileNo string) bool {
	findResult := collection.FindOneAndDelete(context.TODO(), idMobileNo)
	err := findResult.Err()
	if findResult == nil || err == mongo.ErrNoDocuments {
		log.Printf("no document found for deletion")
		return false
	}
	return true
}
