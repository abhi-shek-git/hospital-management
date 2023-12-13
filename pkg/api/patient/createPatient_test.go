package patient

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePatient(t *testing.T) {
	body := `{
		"name" : "ABC",
		"mobileno":90
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
	// utils.Database = utils.HMDB_TEST
	NewPat().Create(response, request)
	if response.Code != http.StatusOK {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}
func TestCreatePatientNilBody(t *testing.T) {
	request, err := http.NewRequest(http.MethodPost, "/api/v1/doctor", nil)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST
	NewPat().Create(response, request)
	if response.Code != http.StatusOK {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}
func TestCreatePatientWrongBody(t *testing.T) {
	body := `{
		"name" : 9,
		"mobileno":"abc"
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
	// utils.Database = utils.HMDB_TEST
	NewPat().Create(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}
}

func TestCreatePatientWrongBodyFieldName(t *testing.T) {
	body := `{
		"mobileno":9
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
	// utils.Database = utils.HMDB_TEST
	NewPat().Create(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}
}
func TestCreatePatientWrongBodyFieldMobileNo(t *testing.T) {
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
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST
	NewPat().Create(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}
}
