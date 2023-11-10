package doctor

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/hospital-management/db"
	"github.com/hospital-management/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doc utils.Doctor
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		log.Println("error occured during decoding input body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if doc.Name == "" || doc.MobileNo == 0 {
		log.Println("invalid credentials")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	collection := db.ConnectDB().Collection(utils.Doctors)
	query := bson.M{"mobileno": doc.MobileNo}
	findResult := collection.FindOne(context.TODO(), query)
	var findData utils.Doctor
	err = findResult.Decode(&findData)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Println("error occured during finding data from db", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if findData.MobileNo == doc.MobileNo {
		log.Println("data already exists")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println("error occured during inserting the data in db", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
