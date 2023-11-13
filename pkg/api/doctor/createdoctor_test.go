package doctor

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hospital-management/db"
	"github.com/hospital-management/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateDoctor(t *testing.T) {
	body := `{
		"name" : "ABC",
		"mobileno":123
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/doctor", reqBody)
	if err != nil {
		log.Println("error occured in test case during making request", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST
	DoctorDataAdded(response, request)
	if response.Code != http.StatusOK {
		log.Println("status code does not match", response.Code)
		t.Fail()

	}
	filter := bson.M{"name": "ABC", "mobileno": 123}
	_, err = db.ConnectDB().Collection(utils.Doctors).DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("error occured in test case during deleting data. Erro = %s", err)
	}
}

func TestCreateDoctorNilBody(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST
	reqBody := bytes.NewBuffer(nil)
	request, err := http.NewRequest(http.MethodPost, "/api/v1/doctor", reqBody)
	if err != nil {
		log.Println("error occured in test case during making request", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	DoctorDataAdded(response, request)
	if response.Code != http.StatusBadRequest {
		log.Println("status code does not match", response.Code)
		t.Fail()

	}
}

func TestCreateDoctorWrongBody(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST
	body := `{
		"name" : "ABC"
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/doctor", reqBody)
	if err != nil {
		log.Println("error occured in test case during making request", err)
		t.Fail()
	}
	response := httptest.NewRecorder()

	DoctorDataAdded(response, request)
	if response.Code != http.StatusBadRequest {
		log.Println("status code does not match", response.Code)
		t.Fail()

	}
}
func TestCreateDoctorWrongBody1(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST
	body := `{
		"mobileno":123
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/doctor", reqBody)
	if err != nil {
		log.Println("error occured in test case during making request", err)
		t.Fail()
	}
	response := httptest.NewRecorder()

	DoctorDataAdded(response, request)
	if response.Code != http.StatusBadRequest {
		log.Println("status code does not match", response.Code)
		t.Fail()

	}
}
