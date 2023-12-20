package db

import (
	"context"
	"log"

	"github.com/hospital-management/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOneUserByMobileNo(collectionName *mongo.Collection, idMobileNo int) (models.User, error) {
	var user models.User
	query := bson.M{"mobileno": idMobileNo}
	dbFindResult := collectionName.FindOne(context.TODO(), query)
	err := dbFindResult.Err()
	if err == mongo.ErrNoDocuments {
		log.Printf("no document found in db. Error = %s", err)
		return user, mongo.ErrNoDocuments
	}
	if err != nil && err != mongo.ErrNoDocuments {
		log.Printf("error occured during finding data from db, error = %s", err)
		return user, err 
	}
	err = dbFindResult.Decode(&user)
	if err != nil {
		log.Printf("error occured during decding the find result from db to models.user. Error = %s", err)
		return user, err
	}
	return user, nil
}

func InsertOneUser(collection *mongo.Collection, insertData interface{}) error {

	_, err := collection.InsertOne(context.TODO(), insertData)
	if err != nil {
		log.Printf("error occured during inserting the data into db %s", err)
		return err
	}
	return nil
}

func CheckPassword(collectionName *mongo.Collection, idMobileNo int, password string) (bool, error) {
	dbFindResult, err := FindOneUserByMobileNo(collectionName, idMobileNo)
	if err != nil {
		log.Printf("error occured during finding data from db by func findOneUserByMobieNo. Error =%s", err)
		return false, err
	}
	if dbFindResult.Password != password{
		log.Printf("password does not match")
		return false, nil
	}
	return true, nil
}
