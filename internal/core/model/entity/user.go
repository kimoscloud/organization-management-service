package organization

import "time"

type User struct {
	ID        string    `json:"id"`
	LastName  string    `json:"lastName"`
	FirstName string    `json:"firstName"`
	Email     string    `json:"email"`
	LastLogin time.Time `json:"lastLogin"`
	CreatedAt time.Time `json:"createdAt"`
}
