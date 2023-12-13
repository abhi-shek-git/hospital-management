package doctor

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateDoctor(t *testing.T) {
	body := `{
		"name" : "ABCDE",
		"mobileno":909000
		}
	`
	reqBody := bytes.NewBuffer([]byte(body))
	request, err := http.NewRequest(http.MethodPut, "/api/v1/updateDoctor", reqBody)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}
	response := httptest.NewRecorder()
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST
	NewDoc().Update(response, request)
	if response.Code != http.StatusOK {
		log.Printf("status code does not match. Needed =  %d  Got =  %d", http.StatusOK, response.Code)
		t.Fail()

	}

}
