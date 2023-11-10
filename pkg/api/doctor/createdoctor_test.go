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
		"name" : "vippin",
		"mobileno":9996506515
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	// if err != nil {
	// 	log.Println("error occured during marshalling")
	// 	return
	// }
	request, err := http.NewRequest(http.MethodPost, "/api/v1/doctor", reqBody)
	// assert.Error(t, err)
	if err != nil {
		log.Println("error occured in test case during making request", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	CreateDoctor(response, request)
	// assert.Equal(t, http.StatusBadRequest, response.Code)
	if response.Code != http.StatusOK {
		log.Println("status code does not match", response.Code)
		t.Fail()

	}
}
