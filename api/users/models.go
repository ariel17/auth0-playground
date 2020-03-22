package users

import "time"

// User is the representation of a person using on the system.
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	Enabled   bool      `json:"enabled"`
}
