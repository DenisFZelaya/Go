package models

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	FullName string `json:"fullname"`
}
