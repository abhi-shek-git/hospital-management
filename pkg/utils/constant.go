package utils

const (
	PortNo  = ":8080"
	HMDB    = "HMDB"
	Doctors = "Doctors"
)

type Doctor struct {
	Name     string `json:"Name,omitempty"`
	HouseNo  int    `json:"HouseNo,omitempty"`
	MobileNo int    `json:"MobileNo,omitempty"`
	Email    string `json:"Email,omitempty"`
	Patients string `json:"Patients,omitempty"`
}
