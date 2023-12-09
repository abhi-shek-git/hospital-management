package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hospital-management/pkg/api/doctor"
	"github.com/hospital-management/pkg/utils"
)

func main() {
	StartServer()
}

func StartServer() {
	// making router
	router := mux.NewRouter()

	doc := doctor.NewDoc()
	// make handle func after writing the func
	router.HandleFunc("/api/v1/doctor", doc.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/deleteDoctor/{id}", doc.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/fetchDoctor/{id}", doc.Fetch).Methods(http.MethodGet)
	// making machine up and running on 8080
	err := http.ListenAndServe(utils.PortNo, router)
	if err != nil {
		log.Printf("error occured in starting the server. Error =  %s", err)
		return
	}
}
