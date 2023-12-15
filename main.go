package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hospital-management/pkg/api/doctor"
	"github.com/hospital-management/pkg/api/patient"
	"github.com/hospital-management/pkg/utils"
)

func main() {
	StartServer()
}

func StartServer() {
	// making router
	router := mux.NewRouter()

	doc := doctor.NewDoc()
	pat := patient.NewPat()
	// make handle func after writing the func
	// doctor functions
	router.HandleFunc("/api/v1/doctor", doc.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/deleteDoctor/{id}", doc.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/fetchDoctor/{id}", doc.Fetch).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/listDoctor", doc.List).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/updateDoctor", doc.Update).Methods(http.MethodPut)

	// patient functions
	router.HandleFunc("/api/v1/createPatient", pat.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/deletePatient/{id}", pat.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/fetchPatient/{id}", pat.Fetch).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/listPatient", pat.List).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/updatePatient", pat.Update).Methods(http.MethodPut)
	// making machine up and running on 8080
	err := http.ListenAndServe(utils.PortNo, router)
	if err != nil {
		log.Printf("error occured in starting the server. Error =  %s", err)
		return
	}
}
