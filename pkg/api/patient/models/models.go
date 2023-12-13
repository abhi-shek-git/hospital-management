package models

type Patient struct {
	Name     string `json:"Name,omitempty"`
	HouseNo  int    `json:"House_No,omitempty"`
	MobileNo int   
	Email    string `json:"Email,omitempty"`
	Doctor string `json:"Patients,omitempty"`
}
