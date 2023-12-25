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

func ListEmployee(collection *mongo.Collection, idName string) (models.Employee, error) {
	var employ models.Employee
	query := bson.M{"name": idName}
	findResult, err := collection.Find(context.TODO(), query)
	if err != nil {
		log.Printf("error occured during finding data from databade.")
		return employ, err
	}
	for findResult.Next(context.TODO()) {
		err = findResult.Decode(&employ)
		if err != nil {
			log.Printf("error occured during decoding findresult from db. Error = %s", err)
			return employ, err
		}
		err = findResult.Err()
		if err != nil {
			log.Printf("error occured durinf findresult error checking. Error = %s", err)
			break
		}
	}

	return employ, nil
}

func UpdateOneEmployee(collection *mongo.Collection, idMobileNo int, updateData models.Employee) (models.Employee, error) {
	query := bson.M{"mobileno": idMobileNo}
	var inpEmploy models.Employee

	if updateData.Email != "" {
		inpEmploy.Email = updateData.Email
	}
	if updateData.HouseNo != 0 {
		inpEmploy.HouseNo = updateData.HouseNo
	}
	if updateData.MobileNo != 0 {
		inpEmploy.MobileNo = updateData.MobileNo
	}
	if updateData.Name != "" {
		inpEmploy.Name = updateData.Name
	}
	if updateData.Designation != "" {
		inpEmploy.Designation = updateData.Designation
	}
	if updateData.Gender != "" {
		inpEmploy.Gender = updateData.Gender
	}
	if updateData.Department != "" {
		inpEmploy.Department = updateData.Department
	}
	if updateData.PostName != "" {
		inpEmploy.PostName = updateData.PostName
	}
	updateEmp := bson.M{
		"$set": inpEmploy,
	}
	_, err := collection.UpdateOne(context.TODO(), query, updateEmp)
	if err != nil {
		log.Printf("error occured during performing update operation in db. Error = %s", err)
		return updateData, err
	}

	// finding updated data
	updatedFind, err := FindOneByMobile(collection, idMobileNo)
	if err != nil {
		log.Printf("error occured in update func in db using findByMobileNo func execution. Error = %s", err)
		return updateData, err
	}
	return updatedFind, nil
}
