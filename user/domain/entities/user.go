package entities

import "time"

type User struct {
	ID string
	Name string
	Email string
	Created_At time.Time
}

func Create(name, email string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Created_At: time.Now(),
	}
}