package doctor

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_createDoctor(t *testing.T) {
	body := `{
		"name" : "",
		"mobileno":0
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
	if response.Code != http.StatusOK {
		log.Println("status code does not match", response.Code)
		t.Fail()

	}
}
