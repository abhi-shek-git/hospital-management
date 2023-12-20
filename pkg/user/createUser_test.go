package user

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	body := `{
		"name" : "ABCDE",
		"mobileno":909,
		"gender" : "male",
		"department":"cardiology",
		"password":"abc123",
		"postname" : "helper"
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/createUser", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST
	NewUser().Create(response, request)
	if response.Code != http.StatusOK {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}
func TestCreateUserWrongBodyFieldMobileno(t *testing.T) {
	
	body := `{
		"name" : "ABC"
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/createUser", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()

	NewUser().Create(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusBadRequest, response.Code)
		t.Fail()

	}
}
func TestCreateUserWrongBodyFieldName(t *testing.T) {
	
	body := `{
		"mobileno":123
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/createUser", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()

	NewUser().Create(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusBadRequest, response.Code)
		t.Fail()

	}
}
