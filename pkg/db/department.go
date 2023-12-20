package db

import (
	"context"
	"log"

	"github.com/hospital-management/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDept(collection *mongo.Collection, insertData interface{}) error {

	_, err := collection.InsertOne(context.TODO(), insertData)
	if err != nil {
		log.Printf("error occured during inserting the data into db %s", err)
		return err
	}
	return nil
}
func FindOne(collection *mongo.Collection, id string) (models.Department, error) {
	var dept models.Department
	indexmodel := mongo.IndexModel{
		Keys: bson.D{{Key: "mobileno", Value: 1}},
	}
	nameOfIndex, err := collection.Indexes().CreateOne(context.TODO(), indexmodel)
	if err != nil {
		log.Printf("error occured in indexing. Error = %s", err)
		return dept, err
	}
	log.Printf("nameOfIndex = %s", nameOfIndex)
	query := bson.M{"id": id}
	findResult := collection.FindOne(context.TODO(), query)
	err = findResult.Err()

	if err != nil {
		log.Printf("error occured during finding data from db. Error=%s", err)
		return dept, err
	}
	err = findResult.Decode(&dept)
	if err != nil {
		log.Printf("error occured during decoding findresult to dept. Error =%s", err)
		return dept, err
	}
	return dept, nil
}
