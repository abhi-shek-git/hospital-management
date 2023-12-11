package db

import (
	"context"
	"log"

	"github.com/hospital-management/pkg/api/doctor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOneByMobileNo(collectionName *mongo.Collection, idMobileNo int) (models.Doctor, error) {
	var doc models.Doctor
	query := bson.M{"mobileno": idMobileNo}
	dbFindResult := collectionName.FindOne(context.TODO(), query)
	err := dbFindResult.Err()
	if err == mongo.ErrNoDocuments {
		log.Printf("no document found in db")
		return doc, mongo.ErrNoDocuments
	}
	if err != nil && err != mongo.ErrNoDocuments {
		log.Printf("error occured during finding data from db, error = %s", err)
		return doc, err
	}
	err = dbFindResult.Decode(&doc)
	if err != nil {
		log.Printf("error occured during decding the find result from db to models.doctor. Error = %s", err)
		return doc, err
	}
	return doc, nil
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

func List(collection *mongo.Collection, idName string) (
	models.Doctor, error) {
	var doc models.Doctor
	query := bson.M{"name": idName}
	findResult, err := collection.Find(context.TODO(), query)
	if err != nil {
		log.Printf("error occured during finding data from databade.")
		return doc, err
	}
	for findResult.Next(context.TODO()) {
		err = findResult.Decode(&doc)
		if err != nil {
			log.Printf("error occured during decoding findresult from db. Error = %s", err)
			return doc, err
		}
		err = findResult.Err()
		if err != nil {
			log.Printf("error occured durinf findresult error checking. Error = %s", err)
			break
		}
	}

	return doc, nil
}
