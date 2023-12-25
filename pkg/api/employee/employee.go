package employee

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hospital-management/pkg/api/models"
	"github.com/hospital-management/pkg/db"
	"github.com/hospital-management/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type Employee interface {
	Create(w http.ResponseWriter, r *http.Request)
	Fetch(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)

}

type employee struct {
	collection *mongo.Collection
}

func NewEmployee() Employee {

	return &employee{collection: db.Connect().Collection(utils.EmployeeCollection)}
}

func (e *employee) Create(w http.ResponseWriter, r *http.Request) {
	var employ models.Employee

	//  decoding input and checking if input is not empty
	err := json.NewDecoder(r.Body).Decode(&employ)
	if err != nil {
		log.Printf("error occured during decoding input body, error = %s", err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// checking input is valid or not
	validate := validateInput(employ)
	if !validate {
		http.Error(w, "name, mobile number, department, designation, gender field can not be empty", http.StatusBadRequest)
		return
	}

	// finding data from database for validation
	_, dbEmployErr := db.FindOneByMobile(e.collection, employ.MobileNo)
	if dbEmployErr != nil && dbEmployErr != mongo.ErrNoDocuments {
		http.Error(w, dbEmployErr.Error(), http.StatusInternalServerError)
		return
	}
	if dbEmployErr == nil {
		http.Error(w, "data already exists", http.StatusBadRequest)
		log.Printf("data already exists")
		return
	}

	//  inserting data into database
	insertResult := db.Insert(e.collection, employ)
	if insertResult != nil {
		http.Error(w, insertResult.Error(), http.StatusInternalServerError)
		return
	}

	// sending stored data back to response
	ResponseByte, err := json.Marshal(employ)
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

func validateInput(employ models.Employee) bool {
	if employ.MobileNo == 0 {
		log.Printf("invalid input, mobile no. field is empty")
		return false
	}
	if employ.Name == "" {
		log.Printf("invalid input, name field is empty")
		return false
	}
	if employ.Gender == "" {
		log.Printf("invalid input, gender field can not be empty")
		return false
	}
	if employ.Department == "" {
		log.Printf("invalid input, department field can not be empty")
		return false
	}
	if employ.Designation == "" {
		log.Printf("invalid input, designation field can not be empty")
		return false
	}
	if employ.PostName == "" {
		log.Printf("invalid input, designation field can not be empty")
		return false
	}
	if employ.PostName != "sweeper" {
		log.Printf("invalid input, post Name does not match")
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

func (e *employee) Fetch(w http.ResponseWriter, r *http.Request) {
	inputId := mux.Vars(r)

	// validating input data
	id, err := ValidateInputId(inputId)
	if err != nil {
		log.Printf("input id should not be empty. Error = %s", err)
		http.Error(w, "invalid input id", http.StatusBadRequest)
		return
	}

	// fetching data from db
	empl, err := db.FindOneByMobile(e.collection, id)
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
	ResponseByte, err := json.Marshal(empl)
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

func (e *employee) List(w http.ResponseWriter, r *http.Request) {
	inpId := r.URL.Query().Get("name")

	// checking if input query is not empty
	if inpId == "" {
		log.Printf("empty input query")
		http.Error(w, "query field can not be empty", http.StatusBadRequest)
		return
	}

	//fetching data from databas

	emp, err := db.ListEmployee(e.collection, inpId)
	if err != nil {
		log.Printf("error occured during fnding data from db. Error = %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	outByte, err := json.Marshal(emp)
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

func (e *employee) Update(w http.ResponseWriter, r *http.Request) {
	var employ models.Employee

	//  decoding input and checking if input is not empty
	err := json.NewDecoder(r.Body).Decode(&employ)
	if err != nil {
		log.Printf("error occured during decoding input body, error = %s", err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	log.Println("??????????", employ.MobileNo)
	// validating input data
	if employ.MobileNo == 0 {
		log.Printf("mobile No field can not be empty.")
		http.Error(w, "mobile No field can not be empty.", http.StatusBadRequest)
		return
	}

	// fetching data from db
	updatedEmp, err := db.UpdateOneEmployee(e.collection, employ.MobileNo, employ)
	if err != nil {
		log.Printf("error occured during updating data. Error = %s", err)
		http.Error(w, "internal server error.", http.StatusInternalServerError)
		return
	}

	// sending response
	ResponseByte, err := json.Marshal(updatedEmp)
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