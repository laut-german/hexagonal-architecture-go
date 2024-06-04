package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	UserController "hexagonal-architecture-go/user/infrastructure/adapters/controller"
	UserRepository "hexagonal-architecture-go/user/infrastructure/adapters/db"
	"log"
	"net/http"
)

func main() {
	// PostgreSQL setup
	postgresDB, err := sql.Open("postgres", "postgres://admin:admin@localhost:5432/hex-db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer postgresDB.Close()

	userPostgresRepo := UserRepository.NewPotgresUserRepository(postgresDB)
	userPostgresController := UserController.NewUserController(userPostgresRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes endpoints
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userPostgresController.CreateUser)
		r.Get("/", userPostgresController.ListUsers)
		r.Get("/{id}", userPostgresController.GetUserByID)
		r.Put("/{id}", userPostgresController.UpdateUser)
	})
	log.Default().Println("Listening at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
