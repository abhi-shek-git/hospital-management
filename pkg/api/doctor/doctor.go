package doctor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hospital-management/db"
	"github.com/hospital-management/pkg/api/doctor/models"
	"github.com/hospital-management/pkg/utils"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var doc models.Doctor

	//  decoding input and checking if input is not empty
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		log.Printf("error occured during decoding input body, error = %s", err)
		w.WriteHeader(http.StatusBadRequest)
		formatedErr := fmt.Errorf("invalid input")
		_, err = w.Write([]byte(formatedErr.Error()))
		if err != nil {
			log.Printf("error occured during writing error data in response. Error =  %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	// checking input is valid or not
	validatedInput := validateInput(doc)
	if !validatedInput {
		w.WriteHeader(http.StatusBadRequest)
		formatedErr := fmt.Errorf("name or mobile number field can not be empty")
		_, err = w.Write([]byte(formatedErr.Error()))
		if err != nil {
			log.Printf("error occured during writing data to response name field should not be empty. Error =  %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	collection := db.Connect().Collection(utils.DoctorCollection)

	// finding data from database for validation
	dbFindResult := db.FindOneByMobileNo(collection, doc.MobileNo)
	if dbFindResult == "data already exists" {
		w.WriteHeader(http.StatusConflict)
		formatedErr := fmt.Errorf("data already exists")
		_, err = w.Write([]byte(formatedErr.Error()))
		if err != nil {
			log.Printf("error occured during writing data already exists to response . Error =  %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	if dbFindResult == "internal error" {
		w.WriteHeader(http.StatusInternalServerError)
		formatedErr := fmt.Errorf("internal error")
		_, err = w.Write([]byte(formatedErr.Error()))
		if err != nil {
			log.Printf("error occured during writing data to response when error occured in findresult from db. Error =  %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	//  inserting data into database
	dbInsertResult := db.InsertOne(collection, doc)
	if dbInsertResult == "not inserted" {
		log.Printf("data does not insert in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// sending stored data back to response
	ResponseByte, err := json.Marshal(doc)
	if err != nil {
		log.Printf("error occured during marshalling the data for response. Error =  %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(ResponseByte)
	if err != nil {
		log.Printf("error occured during writing response for sending data back. Error =  %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func validateInput(doc models.Doctor) bool {
	if doc.MobileNo == 0 {
		log.Printf("invalid input, mobile no. field is empty")
		return false
	}
	if doc.Name == "" {
		log.Printf("invalid input, name field is empty")
		return false
	}
	return true

}
