package doctor

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hospital-management/pkg/api/doctor/models"
	"github.com/hospital-management/pkg/db"
	"github.com/hospital-management/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type Doctor interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Fetch(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type doc struct {
	collection *mongo.Collection
}

func NewDoc() Doctor {
	return &doc{collection: db.Connect().Collection(utils.DoctorCollection)}
}

func (d *doc) Create(w http.ResponseWriter, r *http.Request) {
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

	// finding data from database for validation
	_, dbDoctorErr := db.FindOneByMobileNo(d.collection, doc.MobileNo)
	if dbDoctorErr != nil && dbDoctorErr != mongo.ErrNoDocuments {
		http.Error(w, dbDoctorErr.Error(), http.StatusInternalServerError)
		return
	}
	if dbDoctorErr == nil {
		http.Error(w, "data already exists", http.StatusBadRequest)
		return
	}

	//  inserting data into database
	insertResult := db.InsertOne(d.collection, doc)
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
	if doc.Gender == "" {
		log.Printf("invalid input, gender field can not be empty")
		return false
	}
	if doc.Department == "" {
		log.Printf("invalid input, department field can not be empty")
		return false
	}
	return true
}

func ValidateInputId(inputId map[string]string) (int, error) {
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

func (d *doc) Fetch(w http.ResponseWriter, r *http.Request) {
	inputId := mux.Vars(r)

	// validating input data
	id, err := ValidateInputId(inputId)
	if err != nil {
		log.Printf("input id should not be empty. Error = %s", err)
		http.Error(w, "invalid input id", http.StatusBadRequest)
		return
	}

	// fetching data from db
	doc, err := db.FindOneByMobileNo(d.collection, id)
	if err == mongo.ErrNoDocuments {
		log.Printf("no documents found. Error = %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Printf("error occured during finding data from db. Error = %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// sending response
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

}

func (d *doc) List(w http.ResponseWriter, r *http.Request) {
	inpId := r.URL.Query().Get("name")

	// checking if input query is not empty
	if inpId == "" {
		log.Printf("empty input query")
		http.Error(w, "query field can not be empty", http.StatusBadRequest)
		return
	}

	//fetching data from databas

	doc, err := db.List(d.collection, inpId)
	if err != nil {
		log.Printf("error occured during fnding data from db. Error = %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	outByte, err := json.Marshal(doc)
	if err != nil {
		log.Printf("error occured in marshalling the find data. Error = %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// sending response
	_, err = w.Write(outByte)
	if err != nil {
		log.Printf("error occured during sending data for output. Error = %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (d *doc) Update(w http.ResponseWriter, r *http.Request) {
	var doc models.Doctor

	//  decoding input and checking if input is not empty
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		log.Printf("error occured during decoding input body, error = %s", err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	log.Println("??????????", doc.MobileNo)
	// validating input data
	if doc.MobileNo == 0 {
		log.Printf("mobile No field can not be empty.")
		http.Error(w, "mobile No field can not be empty.", http.StatusBadRequest)
		return
	}

	// fetching data from db
	updatedDoc, err := db.UpdateOne(d.collection, doc.MobileNo, doc)
	if err != nil {
		log.Printf("error occured during updating data. Error = %s", err)
		http.Error(w, "internal server error.", http.StatusInternalServerError)
		return
	}

	// sending response
	ResponseByte, err := json.Marshal(updatedDoc)
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

}
func (d *doc) Delete(w http.ResponseWriter, r *http.Request) {
	var2 := mux.Vars(r)

	// validating input id
	id, err := ValidateInputId(var2)
	if err != nil {
		log.Printf("error occured during validating input id. Error = %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// finding and deleting data from db
	deleteErr := db.FindOneAndDelete(d.collection, id)

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
