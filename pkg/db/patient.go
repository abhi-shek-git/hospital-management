package db

import (
	"context"
	"log"

	"github.com/hospital-management/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindPatientOneByMobileNo(collectionName *mongo.Collection, idMobileNo int) (models.Patient, error) {
	var pat models.Patient
	query := bson.M{"mobileno": idMobileNo}
	dbFindResult := collectionName.FindOne(context.TODO(), query)
	err := dbFindResult.Err()
	if err == mongo.ErrNoDocuments {
		log.Printf("no document found in db. Error = %s", err)
		return pat, mongo.ErrNoDocuments
	}
	if err != nil && err != mongo.ErrNoDocuments {
		log.Printf("error occured during finding data from db, error = %s", err)
		return pat, err
	}
	err = dbFindResult.Decode(&pat)
	if err != nil {
		log.Printf("error occured during decding the find result from db to models.patient . Error = %s", err)
		return pat, err
	}
	return pat, nil
}

func UpdateOnePatient(collection *mongo.Collection, idMobileNo int, updateData models.Patient) (models.Patient, error) {
	query := bson.M{"mobileno": idMobileNo}
	var inpPat models.Patient

	if updateData.Email != "" {
		inpPat.Email = updateData.Email
	}
	if updateData.HouseNo != 0 {
		inpPat.HouseNo = updateData.HouseNo
	}
	if updateData.MobileNo != 0 {
		inpPat.MobileNo = updateData.MobileNo
	}
	if updateData.Name != "" {
		inpPat.Name = updateData.Name
	}
	if updateData.Doctor != "" {
		inpPat.Doctor = updateData.Doctor
	}
	if updateData.Gender != "" {
		inpPat.Gender = updateData.Gender
	}
	if updateData.Department != "" {
		inpPat.Department = updateData.Department
	}

	updateDoc := bson.M{
		"$set": inpPat,
	}
	_, err := collection.UpdateOne(context.TODO(), query, updateDoc)
	if err != nil {
		log.Printf("error occured during performing update operation in db. Error = %s", err)
		return updateData, err
	}

	// finding updated data
	updatedFind, err := FindPatientOneByMobileNo(collection, idMobileNo)
	if err != nil {
		log.Printf("error occured in update func in db using findByMobileNo func execution. Error = %s", err)
		return updateData, err
	}
	return updatedFind, nil
}
