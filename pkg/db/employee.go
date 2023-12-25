package db

import (
	"context"
	"log"

	"github.com/hospital-management/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOneByMobile(collectionName *mongo.Collection, idMobileNo int) (models.Employee, error) {
	var employ models.Employee
	query := bson.M{"mobileno": idMobileNo}
	dbFindResult := collectionName.FindOne(context.TODO(), query)
	err := dbFindResult.Err()
	if err == mongo.ErrNoDocuments {
		log.Printf("no document found in db. Error = %s", err)
		return employ, mongo.ErrNoDocuments
	}
	if err != nil && err != mongo.ErrNoDocuments {
		log.Printf("error occured during finding data from db, error = %s", err)
		return employ, err
	}
	err = dbFindResult.Decode(&employ)
	if err != nil {
		log.Printf("error occured during decoding the find result from db to models.employee. Error = %s", err)
		return employ, err
	}
	return employ, nil
}
func Insert(collection *mongo.Collection, insertData interface{}) error {
	_, err := collection.InsertOne(context.TODO(), insertData)
	if err != nil {
		log.Printf("error occured during inserting the data into db %s", err)
		return err
	}
	return nil
}
