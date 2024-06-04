package entities

import "time"

type User struct {
	ID string
	Name string
	Email string
	Created_At time.Time
}