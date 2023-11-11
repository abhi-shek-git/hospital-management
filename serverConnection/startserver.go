package serverConnection

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hospital-management/pkg/api/doctor"
	"github.com/hospital-management/pkg/utils"
)

func StartServer() {
	// making router
	router := mux.NewRouter()
	// make handle func after writing the func
	router.HandleFunc("/api/v1/doctor", doctor.CreateDoctor).Methods(http.MethodPost)
	// making machine up and running on 8080
	err := http.ListenAndServe(utils.PortNo, router)
	if err != nil {
		log.Fatal("error occured during listen and serve process", err)
	}
}
