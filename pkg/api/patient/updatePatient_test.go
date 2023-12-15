package patient

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdatePatient(t *testing.T) {
	body := `{
		"name" : "",
		"mobileno":909,
		"houseno": 11,
		"email": "abhi",
		"patients": "koi nhin",
		"gender" : "",
		"department" :"neuro"
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPut, "/api/v1/updatePatient", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST
	NewPat().Update(response, request)
	if response.Code != http.StatusOK {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}
func TestUpdatePatientWrongId(t *testing.T) {
	body := `{
		"name" : "ABCDE",
		"mobileno":a ,
		"houseno": 1,
		"email": "abhi",
		"patients": "koi nhin"
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPut, "/api/v1/updatePatient", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST
	NewPat().Update(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}
func TestUpdatePatientNilId(t *testing.T) {
	body := `{
		"name" : "ABCDE",
		"houseno": 1,
		"email": "abhi",
		"patients": "koi nhin"
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPut, "/api/v1/updatePatient", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST
	NewPat().Update(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}
func TestUpdatePatientWrongUpdate(t *testing.T) {
	body := `{
		"name" : "ABCDE",
		"mobileno":9,
		"houseno": 1,
		"email": "abhi",
		"patients": 0
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPut, "/api/v1/updatePatient", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST
	NewPat().Update(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}
