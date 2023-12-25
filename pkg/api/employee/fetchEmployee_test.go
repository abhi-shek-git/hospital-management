package employee

import (
	"log"
	"net/http"
	"testing"
)

func TestFetchOk(t *testing.T) {

	// making request
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/v1/fetchEmployee/90", nil)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}

	// building  client
	cli := http.Client{}
	res, err := cli.Do(request)
	if err != nil {
		log.Printf("error in client response. Error =  %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println("status code does not match", res.StatusCode)
		t.Fail()

	}
}
func TestFetchWrongId(t *testing.T) {

	// making request
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/v1/fetchEmployee/90", nil)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}

	// building  client
	cli := http.Client{}
	res, err := cli.Do(request)
	if err != nil {
		log.Printf("error in client response. Error =  %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		log.Println("status code does not match", res.StatusCode)
		t.Fail()

	}
}
func TestFetchNilId(t *testing.T) {

	// making request
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/v1/fetchEmployee", nil)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}

	// building  client
	cli := http.Client{}
	res, err := cli.Do(request)
	if err != nil {
		log.Printf("error in client response. Error =  %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		log.Println("status code does not match", res.StatusCode)
		t.Fail()

	}
}
