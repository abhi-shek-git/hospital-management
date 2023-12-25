package employee

import (
	"log"
	"net/http"
	"testing"
)

func TestList(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST

	//building request
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/v1/listEmployee?name=ABCD", nil)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}

	// building client
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
func TestListWrongInp(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST

	//building request
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/v1/listEmployee?name=123", nil)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}

	// building client
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
func TestListNoInp(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST

	//building request
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/v1/listEmployee", nil)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}

	// building client
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
func TestListWrongFieldInp(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	// utils.Database = utils.HMDB_TEST

	//building request
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/v1/listEmployee?mobieno=909", nil)
	if err != nil {
		log.Printf("error occured in test case during making request. Error =  %s", err)
		t.Fail()
	}

	// building client
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
