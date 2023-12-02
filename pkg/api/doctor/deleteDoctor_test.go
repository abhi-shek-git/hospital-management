package doctor

import (
	"log"
	"net/http"
	"testing"

	"github.com/hospital-management/db"
	"github.com/hospital-management/pkg/utils"
)

func TestDelete(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST

	// making request
	request, err := http.NewRequest(http.MethodDelete, "http://127.0.0.1:8080/api/v1/deleteDoctor/1234", nil)
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
func TestDeleteWrongId(t *testing.T) {
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST

	// making request
	request, err := http.NewRequest(http.MethodDelete, "http://127.0.0.1:8080/api/v1/deleteDoctor/abc", nil)
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
	// updating database variable from HMDB to HMDB_TEST
	db.Database = utils.HMDB_TEST

	// making request
	request, err := http.NewRequest(http.MethodDelete, "http://127.0.0.1:8080/api/v1/deleteDoctor", nil)
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
