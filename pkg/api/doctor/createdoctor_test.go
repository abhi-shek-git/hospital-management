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

func Test_createDoctor(t *testing.T) {
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
	CreateDoctor(response, request)
	if response.Code != http.StatusOK {
		log.Println("status code does not match", response.Code)
		t.Fail()

	}
}

func Test_createDoctorNilBody(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST
	reqBody := bytes.NewBuffer(nil)
	request, err := http.NewRequest(http.MethodPost, "/api/v1/doctor", reqBody)
	if err != nil {
		log.Println("error occured in test case during making request", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	CreateDoctor(response, request)
	if response.Code != http.StatusBadRequest {
		log.Println("status code does not match", response.Code)
		t.Fail()

	}
}

func Test_createDoctorNegetive(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST
	body := `{
		"name" : "ABC",
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPost, "/api/v1/doctor", reqBody)
	if err != nil {
		log.Println("error occured in test case during making request", err)
		t.Fail()
	}
	response := httptest.NewRecorder()

	CreateDoctor(response, request)
	if response.Code != http.StatusBadRequest {
		log.Println("status code does not match", response.Code)
		t.Fail()

	}
}
func Test_createDoctorNegetive1(t *testing.T) {
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

	CreateDoctor(response, request)
	if response.Code != http.StatusBadRequest {
		log.Println("status code does not match", response.Code)
		t.Fail()

	}
}
