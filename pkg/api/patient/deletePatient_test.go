package patient

import (
	"log"
	"net/http"
	"testing"
)

func TestDelete(t *testing.T) {

	// making request
	request, err := http.NewRequest(http.MethodDelete, "http://127.0.0.1:8080/api/v1/deletePatient/9", nil)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}

	// makingclient
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
func TestDeleteNilId(t *testing.T) {

	// making request
	request, err := http.NewRequest(http.MethodDelete, "http://127.0.0.1:8080/api/v1/deletePatient", nil)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}

	// makingclient
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
func TestDeleteWrongId(t *testing.T) {

	// making request
	request, err := http.NewRequest(http.MethodDelete, "http://127.0.0.1:8080/api/v1/deletePatient/abc", nil)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}

	// makingclient
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
