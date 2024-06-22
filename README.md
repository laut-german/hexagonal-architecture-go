# Hexagonal Architecture Go Application

![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue)
![License: GPL v3.0](https://img.shields.io/badge/license-GPL%20v3.0-red)

## Hexagonal Architecture

Hexagonal architecture, also known as ports and adapters architecture, is a design pattern used in software development. This architecture promotes the separation of concerns and allows the application to be independent from external agencies, such as user interfaces and databases.

For more information, you can read about hexagonal architecture on [Wikipedia](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)).

## Description
This repository contains a Go application following the Hexagonal Architecture pattern. Below are the steps to set up and run the application locally.

## Prerequisites

- Docker
- Docker Compose
- Go 1.22.3 or higher

## Getting Started

### 1. Clone the Repository

```bash
git clone git@github.com:laut-german/hexagonal-architecture-go.git
```

### 2. Set Up Docker Containers

Use Docker Compose to set up the PostgreSQL and RabbitMQ containers.

```bash
docker-compose up -d
```

This will start the following services:
- PostgreSQL on port 5432
- RabbitMQ on port 5672 (and management UI on port 15672)

### 3. Run Database Migrations

Ensure you have `golang-migrate` installed. If not, install it using the following command:

```bash
brew install golang-migrate
```

Run the migrations to set up the database schema:

```bash
make migrate-up
```

This command applies all pending migrations to the PostgreSQL database.

### 4. Create a Test User

To create a test user, you can use a tool like `curl` or Postman to send a POST request to the `/users` endpoint.

Using `curl`:

```bash
curl -X POST http://localhost:8000/users \
     -H "Content-Type: application/json" \
     -d '{
           "email": "testuser@example.com"
           "name": "Test User"
         }'
```

This will create a new user in the database.

### 5. Run the Application

Start the Go application:

```bash
go run cmd/hex-app/main.go
```

The application will start and listen on port 8000.

## Endpoints

### User Endpoints

- **Create User**: `POST /users`
- **List Users**: `GET /users`
- **Get User by ID**: `GET /users/{id}`
- **Update User**: `PUT /users/{id}`

### Shopping Cart Endpoints

- **Add Item to Cart**: `POST /cart/add-item`
- **Remove Item from Cart**: `POST /cart/remove-item`
- **Get Cart by User ID**: `GET /cart/user/{id}`
- **List Cart Items by User ID**: `GET /cart/list-items/user/{id}`

## Project Structure

```
.
├── Makefile
├── cmd
│   └── hex-app
│       └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── migrations
│   ├── 000001_init_schema.down.sql
│   └── 000001_init_schema.up.sql
├── internal
   ├── shopping-cart
   │   ├── application
   │   │   ├── commands
   │   │   │   ├── add-item-to-cart.go
   │   │   │   └── remove-item-from-cart.go
   │   │   ├── queries
   │   │   │   ├── get_cart_by_user_id.go
   │   │   │   └── list_cart_items.go
   │   │   └── responses
   │   │       └── cart_response.go
   │   ├── domain
   │   │   ├── entities
   │   │   │   └── cart.go
   │   │   └── ports
   │   │       ├── queue
   │   │       │   ├── queue_consumer.go
   │   │       │   └── queue_producer.go
   │   │       └── repositories
   │   │           └── cart_repository.go
   │   └── infrastructure
   │       └── adapters
   │           ├── controller
   │           │   └── shopping_cart_controller.go
   │           ├── db
   │           │   └── postgres_cart_repository.go
   │           └── queue
   │               ├── rabbitmq_consumer.go
   │               └── rabbitmq_producer.go
   ├── users
      ├── application
      │   ├── commands
      │   │   ├── create_user.go
      │   │   └── update_user.go
      │   ├── queries
      │   │   ├── get_user_by_id.go
      │   │   └── list_users.go
      │   └── responses
      │       └── create_user.response.go
      ├── domain
      │   ├── entities
      │   │   └── user.go
      |            └── postgres_user_repository.go
      │   ├── errors
      │   │   └── user_already_exists.go
      │   └── ports
      │       └── repositories
      │           └── user_repository.go
      └── infrastructure
          └── adapters
              │   └── user_controller.go
              ├── controller
              └── db
```

## Running the Application

To start the Go application locally, use the following command:

```bash
go run cmd/hex-app/main.go
```

## Conclusion

You now have a running Go application following the Hexagonal Architecture pattern. You can interact with the application via the provided endpoints. Feel free to explore and extend the application as needed.

