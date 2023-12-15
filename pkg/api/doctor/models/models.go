package models

type Doctor struct {
	Name       string `json:"Name,omitempty" bson:"name,omitempty"`
	Gender     string `json:"Gender,omitempty" bson:"gender,omitempty"`
	Department string `json:"Department,omitempty" bson:"department,omitempty"`
	HouseNo    int    `json:"HouseNo,omitempty" bson:"houseno,omitempty"`
	MobileNo   int
	Email      string `json:"Email,omitempty" bson:"email,omitempty"`
	Patients   string `json:"Patients,omitempty" bson:"patients,omitempty"`
}
