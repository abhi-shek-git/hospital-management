package models

type Doctor struct {
	Name     string `json:"Name,omitempty"`
	HouseNo  int    `json:"House_No,omitempty"`
	MobileNo int   
	Email    string `json:"Email,omitempty"`
	Patients string `json:"Patients,omitempty"`
}
