package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	//_ "github.com/lib/pq"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	UserController "hexagonal-architecture-go/user/infrastructure/adapters/controller"
	UserRepository "hexagonal-architecture-go/user/infrastructure/adapters/db"
)

func main() {
	// PostgreSQL setup
	postgresDB, err := sql.Open("postgres", "postgres://user:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer postgresDB.Close()

	// MongoDB setup
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(context.TODO())

	//mongoDB := mongoClient.Database("dbname")

	// Initialize repositories
	userPostgresRepo := UserRepository.NewPotgresUserRepository(postgresDB)
	//userMongoRepo := UserRepository.NewMongoUserRepository(mongoDB, "users")

	// Initialize controllers
	userPostgresController := UserController.NewUserController(userPostgresRepo)
	//userMongoController := UserController.NewUserController(userMongoRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes for PostgreSQL-backed endpoints
	r.Route("/postgres/users", func(r chi.Router) {
		r.Post("/", userPostgresController.CreateUser)
		r.Get("/", userPostgresController.ListUsers)
		r.Get("/{id}", userPostgresController.GetUserByID)
		r.Put("/{id}", userPostgresController.UpdateUser)
	})

	// Routes for MongoDB-backed endpoints
	// r.Route("/mongo/users", func(r chi.Router) {
	//     r.Post("/", userMongoController.CreateUser)
	//     r.Get("/", userMongoController.ListUsers)
	//     r.Get("/{id}", userMongoController.GetUserByID)
	//     r.Put("/{id}", userMongoController.UpdateUser)
	// })

	log.Fatal(http.ListenAndServe(":8080", r))
}
