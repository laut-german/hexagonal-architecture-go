DATABASE_URL=postgres://admin:admin@localhost:5432/hex-db?sslmode=disable
MIGRATIONS_DIR=migrations


.PHONY: migrate-up migrate-down new-migration

new-migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database $(DATABASE_URL) up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database $(DATABASE_URL) down
