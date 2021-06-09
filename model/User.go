package model

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Job       string `json:"job"`
	Gender    string `json:"gender"`
	City      string `json:"city"`
	Salary    int    `json:"salary"`
	BirthDate string `json:"birthdate"`
}
