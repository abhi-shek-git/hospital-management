package models

type User struct {
	Department string `json:"Department,omitempty" bson:"department,omitempty"`
	PostName   string `json:"PostName,omitempty" bson:"postname,omitempty"`
	Name       string `json:"Name,omitempty" bson:"name,omitempty"`
	Gender     string `json:"Gender,omitempty" bson:"gender,omitempty"`
	HouseNo    int    `json:"HouseNo,omitempty" bson:"houseno,omitempty"`
	MobileNo   int
	Email      string `json:"Email,omitempty" bson:"email,omitempty"`
	Password   string `json:"Password,omitempty" bson:"password,omitempty"`
}

type Doctor struct {
	Name       string `json:"Name,omitempty" bson:"name,omitempty"`
	Gender     string `json:"Gender,omitempty" bson:"gender,omitempty"`
	Department string `json:"Department,omitempty" bson:"department,omitempty"`
	HouseNo    int    `json:"HouseNo,omitempty" bson:"houseno,omitempty"`
	MobileNo   int
	Email      string `json:"Email,omitempty" bson:"email,omitempty"`
	Patients   string `json:"Patients,omitempty" bson:"patients,omitempty"`
}
type Patient struct {
	Name       string `json:"Name,omitempty" bson:"name,omitempty"`
	Gender     string `json:"Gender,omitempty" bson:"gender,omitempty"`
	Department string `json:"Department,omitempty" bson:"department,omitempty"`
	HouseNo    int    `json:"HouseNo,omitempty" bson:"houseno,omitempty"`
	MobileNo   int
	Email      string `json:"Email,omitempty" bson:"email,omitempty"`
	Doctor     string `json:"Patients,omitempty" bson:"doctor,omitempty"`
}

type Department struct {
	Id      string `bson: "_id"`
	Name    string `json:"Name,omitempty" bson:"name,omitempty"`
	Doctor  Doctor
	Patient Patient
	User    User
}

type Employee struct {
	Department  string
	PostName    string
	Designation string
	Name        string
	Gender      string
	HouseNo     int
	MobileNo    int
	Email       string
}
