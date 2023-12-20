package department

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDept(t *testing.T) {
	body := `{
		"name" : "ABCDE"
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/createDept", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST
	NewDept().Create(response, request)
	if response.Code != http.StatusOK {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}
func TestCreateDeptWithObjectBody(t *testing.T) {
	body := `{
		"name" : "ABCDE",
		"doctor":{
			"name":"abc",
			"mobileno":123
		}
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/createDept", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST
	NewDept().Create(response, request)
	if response.Code != http.StatusOK {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}
func TestCreateDeptNilBody(t *testing.T) {
	request, err := http.NewRequest(http.MethodPost, "/api/v1/createDept", nil)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST
	NewDept().Create(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}
func TestCreateDeptWrongBody(t *testing.T) {
	body := `{
		"name" : 0
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/createDept", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST
	NewDept().Create(response, request)
	if response.Code != http.StatusBadRequest {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}


