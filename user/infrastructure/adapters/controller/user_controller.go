package controller

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"hexagonal-architecture-go/user/application/commands"
	"hexagonal-architecture-go/user/application/queries"
	"hexagonal-architecture-go/user/domain/ports/repositories"
	userDomainErrors "hexagonal-architecture-go/user/domain/errors"
)

type UserController struct {
	createUserHandler  *commands.CreateUserHandler
	updateUserHandler  *commands.UpdateUserHandler
	getUserByIDHandler *queries.GetUserByIDHandler
	listUsersHandler   *queries.ListUsersHandler
}

func NewUserController(repo repositories.UserRepository) *UserController {
	return &UserController{
		createUserHandler:  commands.NewCreateUserHandler(repo),
		updateUserHandler:  commands.NewUpdateUserHandler(repo),
		getUserByIDHandler: queries.NewGetUserByIDHandler(repo),
		listUsersHandler:   queries.NewListUsersHandler(repo),
	}

}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userCmd commands.CreateUserCommand
	if err := json.NewDecoder(r.Body).Decode(&userCmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCreated, err := uc.createUserHandler.Handle(userCmd)
	if err != nil {
		if userDomainErrors.IsUserAlreadyExistsError(err) {
			http.Error(w, "Invalid email", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(userCreated)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userCmd commands.UpdateUserCommand
	if err := json.NewDecoder(r.Body).Decode(&userCmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := uc.updateUserHandler.Handle(userCmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	user, err := uc.getUserByIDHandler.Handle(queries.GetUserByIDQuery{ID: idStr})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.listUsersHandler.Handle(queries.ListUsersQuery{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
