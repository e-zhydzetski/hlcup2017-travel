package domain

import "time"

type User struct {
	ID        uint32    `json:"id"`         // unique
	Email     string    `json:"email"`      // unique, len <= 100
	FirstName string    `json:"first_name"` // len <= 50
	LastName  string    `json:"last_name"`  // len <= 50
	Gender    string    `json:"gender"`     // m/f
	BirthDate time.Time `json:"birth_date"` // timestamp
}
