package entity

type Person struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
