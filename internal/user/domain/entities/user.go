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

func (u *User) Update(name, email string) {
	// Aqui se puede aplicar la logica de negocio necesaria que se necesite para un update
	if name != "" {
		u.Name = name
	}
	if email != "" {
		u.Email = email
	}
}