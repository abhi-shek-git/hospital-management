package department

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hospital-management/pkg/api/models"
	"github.com/hospital-management/pkg/db"
	"github.com/hospital-management/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type Department interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type dept struct {
	collection *mongo.Collection
}

func NewDept() Department {
	return &dept{collection: db.Connect().Collection(utils.DepartmentCollection)}
}

func (d *dept) Create(w http.ResponseWriter, r *http.Request) {
	var dept models.Department

	// decoding input and checking if input is not empty
	err := json.NewDecoder(r.Body).Decode(&dept)
	if err != nil {
		log.Printf("error occured during decoding input body, error = %s", err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// validate input
	if dept.Name == "" {
		log.Printf("name field can not be empty.")
		http.Error(w, "name of department can not be empty", http.StatusBadRequest)
		return
	}

	// insert in database
	err = db.InsertOneDept(d.collection, dept)
	if err != nil {
		log.Printf("error occured during inserting in db. Error = %s", err)
		http.Error(w, "internal server error.", http.StatusInternalServerError)
		return
	}

	// finding data and sending back
	outData, err := db.FindOne(d.collection, dept.Id)
	if err == mongo.ErrNoDocuments {
		log.Printf("no document found for sending back. %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err != nil {
		log.Printf("error occured during finding data for sending back. %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	// making output byte
	outByte, err := json.Marshal(outData)
	if err != nil {
		log.Printf("error occured during making output byte. Error = %s", err)
		http.Error(w, "internal error.", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(outByte)
	if err != nil {
		log.Printf("error occured during writing data for output on write body. Error = %s", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
