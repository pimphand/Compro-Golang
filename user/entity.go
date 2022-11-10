package user

import "time"

type User struct {
	Id    int
	Name  string
	Email string
	// EmailVerifiedAt time.Time
	Password  string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
