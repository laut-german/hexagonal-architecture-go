package responses

type CreateUserResponse struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func NewUserResponse(id, name, email string) CreateUserResponse {
    return CreateUserResponse{
        ID:    id,
        Name:  name,
        Email: email,
    }
}