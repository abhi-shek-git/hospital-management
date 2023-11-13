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
		log.Printf("error occured during decoding input body, error = %s", err)
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("invalid input"))
		if err != nil {
			log.Printf("error occured during writing data to response in decoding input body. Error =  %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	if doc.Name == "" {
		log.Println("invalid input, name field is empty")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("Name field should not be empty"))
		if err != nil {
			log.Printf("error occured during writing data to response name field should not be empty. Error =  %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	if doc.MobileNo == 0 {
		log.Println("invalid input, mobile no. field is empty")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("mobile number field should not be empty"))
		if err != nil {
			log.Printf("error occured during writing data to response mobileno fiekd should not be empty. Error =  %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	collection := db.ConnectDB().Collection(utils.Doctors)
	query := bson.M{"mobileno": doc.MobileNo}
	dbFindResult := collection.FindOne(context.TODO(), query)
	err = dbFindResult.Err()
	if err == nil {
		log.Printf("data already exists")
		w.WriteHeader(http.StatusConflict)
		_, err = w.Write([]byte("data already exists"))
		if err != nil {
			log.Printf("error occured during writing data already exists to response . Error =  %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	if err != nil && err != mongo.ErrNoDocuments {
		log.Printf("error occured during finding data from db, error = %s /n", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("internal error"))
		if err != nil {
			log.Printf("error occured during writing data to response when error occured in findresult from db. Error =  %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	_, err = collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Printf("error occured during inserting the data into db %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ResponseByte, err := json.Marshal(doc)
	if err != bson.ErrDecodeToNil {
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
