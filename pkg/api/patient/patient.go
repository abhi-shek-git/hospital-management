package patient

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hospital-management/pkg/api/patient/models"
	"github.com/hospital-management/pkg/db"
	"github.com/hospital-management/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type Patient interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type pat struct {
	collection *mongo.Collection
}

func NewPat() Patient {
	return &pat{collection: db.Connect().Collection(utils.PatientCollection)}
}

func (d *pat) Create(w http.ResponseWriter, r *http.Request) {
	var pat models.Patient

	//  decoding input and checking if input is not empty
	err := json.NewDecoder(r.Body).Decode(&pat)
	if err != nil {
		log.Printf("error occured during decoding input body, error = %s", err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// checking input is valid or not
	validate := validateInput(pat)
	if !validate {
		http.Error(w, "name or mobile number field can not be empty", http.StatusBadRequest)
		return
	}

	// finding data from database for validation
	_, dbPatientErr := db.FindOneByMobileNo(d.collection, pat.MobileNo)
	if dbPatientErr != nil && dbPatientErr != mongo.ErrNoDocuments {
		http.Error(w, dbPatientErr.Error(), http.StatusInternalServerError)
		return
	}
	if dbPatientErr == nil {
		http.Error(w, "data already exists", http.StatusBadRequest)
		return
	}

	//  inserting data into database
	insertResult := db.InsertOne(d.collection, pat)
	if insertResult != nil {
		http.Error(w, insertResult.Error(), http.StatusInternalServerError)
		return
	}

	// sending stored data back to response
	ResponseByte, err := json.Marshal(pat)
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

func validateInput(pat models.Patient) bool {
	if pat.MobileNo == 0 {
		log.Printf("invalid input, mobile no. field is empty")
		return false
	}
	if pat.Name == "" {
		log.Printf("invalid input, name field is empty")
		return false
	}
	return true
}

// func (d *doc) Delete(w http.ResponseWriter, r *http.Request) {
// 	var2 := mux.Vars(r)

// 	// validating input id
// 	id, err := ValidateInputId(var2)
// 	if err != nil {
// 		log.Printf("error occured during validating input id. Error = %s", err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// finding and deleting data from db
// 	deleteErr := db.FindOneAndDelete(d.collection, id)

// 	if deleteErr == mongo.ErrNoDocuments {
// 		log.Printf("no documents found. Error = %s", deleteErr)
// 		http.Error(w, deleteErr.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if deleteErr != nil {
// 		log.Printf("error occured during finding data from db. Error = %s", deleteErr)
// 		http.Error(w, deleteErr.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err = w.Write([]byte("Record deleted"))
// 	if err != nil {
// 		log.Printf("error occured during writing data in response body. Error = %s", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func ValidateInputId(inputId map[string]string) (int, error) {
// 	id := ""

// 	for key, value := range inputId {
// 		if key == "id" {
// 			id = value
// 			break
// 		}
// 	}
// 	if id == "" {
// 		log.Printf("delete id is missing in input")
// 		return 0, mongo.ErrNoDocuments
// 	}
// 	findId, err := strconv.Atoi(id)
// 	if err != nil {
// 		log.Printf("error occured during converting input string to integer. Error = %s", err)
// 		return 0, err
// 	}
// 	return findId, nil
// }
