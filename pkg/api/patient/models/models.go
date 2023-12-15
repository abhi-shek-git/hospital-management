package models

type Patient struct {
	Name     string `json:"Name,omitempty" bson:"name,omitempty"`
	Gender string`json:"Gender,omitempty" bson:"gender,omitempty"`
	Department string`json:"Department,omitempty" bson:"department,omitempty"`
	HouseNo  int    `json:"HouseNo,omitempty" bson:"houseno,omitempty"`
	MobileNo int   
	Email    string `json:"Email,omitempty" bson:"email,omitempty"`
	Doctor string `json:"Patients,omitempty" bson:"doctor,omitempty"`
}
