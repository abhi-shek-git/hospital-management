package doctor

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/hospital-management/db"
	"github.com/hospital-management/pkg/api/doctor/models"
	"github.com/hospital-management/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doc models.Doctor
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		log.Println("error occured during decoding input body", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid input"))
		return
	}
	if doc.Name == "" {
		log.Println("invalid input, name field is empty")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name field should not be empty"))
		return
	}
	if doc.MobileNo == 0 {
		log.Println("invalid input, mobile no. field is empty")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("mobile number field should not be empty"))
		return
	}
	collection := db.ConnectDB().Collection(utils.Doctors)
	query := bson.M{"mobileno": doc.MobileNo}
	dbFindResult := collection.FindOne(context.TODO(), query)
	err = dbFindResult.Err()
	if err == nil {
		log.Println("data already exists")
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("data already exists"))
		return
	}
	if err != nil && err != mongo.ErrNoDocuments {
		log.Println("error occured during finding data from db", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal error"))
		return
	}

	_, err = collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Printf("error occured during inserting the data into db %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
