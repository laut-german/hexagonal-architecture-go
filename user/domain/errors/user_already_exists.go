package errors

import "fmt"

type UserAlreadyExistsError struct {
    Email string
}

func (e *UserAlreadyExistsError) Error() string {
    return fmt.Sprintf("user with email %s already exists", e.Email)
}

func IsUserAlreadyExistsError(err error) bool {
    _, ok := err.(*UserAlreadyExistsError)
    return ok
}
