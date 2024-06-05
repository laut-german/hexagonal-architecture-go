package main

import (
	"context"
	"database/sql"
	UserController "hexagonal-architecture-go/user/infrastructure/adapters/controller"
	UserRepository "hexagonal-architecture-go/user/infrastructure/adapters/db"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"

	CartController "hexagonal-architecture-go/shopping-cart/infrastructure/adapters/controller"
	CartRepository "hexagonal-architecture-go/shopping-cart/infrastructure/adapters/db"
	CartQueue "hexagonal-architecture-go/shopping-cart/infrastructure/adapters/queue"

	"github.com/rabbitmq/amqp091-go"
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

	// RabbitMQ setup
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	queueName := "cart_queue"

	cartRepo := CartRepository.NewPostgresCartRepository(postgresDB)
	cartConsumer := CartQueue.NewRabbitMQConsumer(channel, queueName, cartRepo)
	cartProducer := CartQueue.NewRabbitMQProducer(channel, queueName)
	cartController := CartController.NewShoppingCartController(cartRepo, cartProducer)

	// Start consuming messages in a separate goroutine
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := cartConsumer.StartConsuming(ctx, cartConsumer.HandleMessage); err != nil {
			log.Fatalf("Failed to start consumer: %v", err)
		}
	}()

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

	// Shopping Cart routes endpoints
	r.Route("/cart", func(r chi.Router) {
		r.Post("/add-item", cartController.AddItemToCart)
		r.Post("/remove-item", cartController.RemoveItemFromCart)
		r.Get("/get", cartController.GetCartByUserID)
		r.Get("/list-items", cartController.ListCartItems)
	})
	log.Default().Println("Listening at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
