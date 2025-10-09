package models

type Employee struct {
	ID       int    `json:"id"`
	Company  string `json:"company"`
	Name     string `json:"name"`
	JoinDate string `json:"join_date"`
	Role     string `json:"role"`
	Paygrade string `json:"paygrade"`
}

type NewEmployeeRequest struct {
	Company  string `json:"company"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Paygrade string `json:"paygrade"`
}
