package doctor

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hospital-management/db"
	"github.com/hospital-management/pkg/utils"
)

func TestCreateDoctor(t *testing.T) {
	body := `{
		"name" : "ABCD",
		"mobileno":123456
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/doctor", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST
	Create(response, request)
	if response.Code != http.StatusOK {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}

func TestCreateDoctorNilBody(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST
	reqBody := bytes.NewBuffer(nil)
	request, err := http.NewRequest(http.MethodPost, "/api/v1/doctor", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	Create(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusBadRequest, response.Code)
		t.Fail()

	}
}

func TestCreateDoctorWrongBodyFieldMobileno(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST
	body := `{
		"name" : "ABC"
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/doctor", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()

	Create(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusBadRequest, response.Code)
		t.Fail()

	}
}
func TestCreateDoctorWrongBodyFieldName(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST
	body := `{
		"mobileno":123
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/doctor", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()

	Create(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusBadRequest, response.Code)
		t.Fail()

	}
}
