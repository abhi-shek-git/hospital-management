package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hospital-management/pkg/api/models"
	"github.com/hospital-management/pkg/db"
	"github.com/hospital-management/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	Create(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)

}

type user struct {
	collection *mongo.Collection
}

func NewUser() User {
	return &user{collection: db.Connect().Collection(utils.UserCollection)}
}

func (u *user) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	//  decoding input and checking if input is not empty
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("error occured during decoding input body, error = %s", err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// checking input is valid or not
	validate := validateInput(user)
	if !validate {
		http.Error(w, "name, mobile, department, gender, post name and password field can not be empty", http.StatusBadRequest)
		return
	}

	// finding data from database for validation
	_, dbUserErr := db.FindOneUserByMobileNo(u.collection, user.MobileNo)
	if dbUserErr != nil && dbUserErr != mongo.ErrNoDocuments {
		http.Error(w, dbUserErr.Error(), http.StatusInternalServerError)
		return
	}
	if dbUserErr == nil {
		http.Error(w, "data already exists", http.StatusBadRequest)
		return
	}

	//  inserting data into database
	insertResult := db.InsertOneUser(u.collection, user)
	if insertResult != nil {
		http.Error(w, insertResult.Error(), http.StatusInternalServerError)
		return
	}

	// sending stored data back to response
	ResponseByte, err := json.Marshal(user)
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

func validateInput(user models.User) bool {
	if user.MobileNo == 0 {
		log.Printf("invalid input, mobile no. field is empty")
		return false
	}
	if user.Name == "" {
		log.Printf("invalid input, name field is empty")
		return false
	}
	if user.Gender == "" {
		log.Printf("invalid input, gender field can not be empty")
		return false
	}
	if user.Department == "" {
		log.Printf("invalid input, department field can not be empty")
		return false
	}
	if user.Password == "" {
		log.Printf("invalid input, password field can not be empty")
		return false
	}
	if user.PostName == "" {
		log.Printf("invalid input, post name field can not be empty")
		return false
	}
	return true
}

func (u *user) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	//  decoding input and checking if input is not empty
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("error occured during decoding input body, error = %s", err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// checking input is valid or not
	validate := validateLoginInput(user)
	if !validate {
		http.Error(w, "mobile and password field can not be empty", http.StatusBadRequest)
		return
	}

	// validating password with db
	boolVar, err := db.CheckPassword(u.collection, user.MobileNo, user.Password)
	if err != nil && boolVar == false {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err == nil && boolVar == false {
		http.Error(w, "invalid password", http.StatusBadRequest)
		return
	}
	_, err = w.Write([]byte("login successful"))
	if err != nil {
		log.Printf("error occured during writing response for sending data back. Error =  %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func validateLoginInput(user models.User) bool {
	if user.MobileNo == 0 {
		log.Printf("invalid input, mobile no. field is empty")
		return false
	}
	if user.Password == "" {
		log.Printf("invalid input, password field can not be empty")
		return false
	}
	return true
}


