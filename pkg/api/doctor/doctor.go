package doctor

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hospital-management/db"
	"github.com/hospital-management/pkg/api/doctor/models"
	"github.com/hospital-management/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var doc models.Doctor

	//  decoding input and checking if input is not empty
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		log.Printf("error occured during decoding input body, error = %s", err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	// checking input is valid or not
	validate := validateInput(doc)
	if !validate {
		http.Error(w, "name or mobile number field can not be empty", http.StatusBadRequest)
		return
	}

	collection := db.Connect().Collection(utils.DoctorCollection)

	// finding data from database for validation
	dbDoctorErr := db.FindOneByMobileNo(collection, doc.MobileNo)
	if dbDoctorErr != nil && dbDoctorErr != mongo.ErrNoDocuments {
		http.Error(w, dbDoctorErr.Error(), http.StatusInternalServerError)
		return
	}
	if dbDoctorErr == nil {
		http.Error(w, "data already exists", http.StatusBadRequest)
		return
	}

	//  inserting data into database
	insertResult := db.InsertOne(collection, doc)
	if insertResult != nil {
		http.Error(w, insertResult.Error(), http.StatusInternalServerError)
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
func Delete(w http.ResponseWriter, r *http.Request) {
	var2 := mux.Vars(r)

	// validating input id
	id, err := ValidateDeleteInput(var2)
	if err != nil {
		log.Printf("error occured during validating input id. Error = %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := db.Connect().Collection(utils.DoctorCollection)

	// finding and deleting data from db

	deleteErr := db.FindOneAndDelete(collection, id)

	if deleteErr == mongo.ErrNoDocuments {
		log.Printf("no documents found. Error = %s", deleteErr)
		http.Error(w, deleteErr.Error(), http.StatusBadRequest)
		return
	}

	if deleteErr != nil {
		log.Printf("error occured during finding data from db. Error = %s", deleteErr)
		http.Error(w, deleteErr.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte("Record deleted"))
	if err != nil {
		log.Printf("error occured during writing data in response body. Error = %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ValidateDeleteInput(inputId map[string]string) (int, error) {
	id := ""

	for key, value := range inputId {
		if key == "id" {
			id = value
			break
		}
	}
	if id == "" {
		log.Printf("delete id is missing in input")
		return 0, mongo.ErrNoDocuments
	}
	findId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error occured during converting input string to integer. Error = %s", err)
		return 0, err
	}
	return findId, nil
}
