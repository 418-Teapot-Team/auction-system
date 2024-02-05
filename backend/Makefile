DATABASE_URI="postgres://postgres:root@localhost:5432/app?sslmode=disable"

MIGRATIONS_DIR="migrations"

.PHONY: server fmt linter migrate-up migrate-down

server:
	go run cmd/main.go

fmt:
	gofumpt -l -w .

linter:
	golangci-lint run --deadline 5m --verbose

migrate-up:
	migrate -path migrations -database $(DATABASE_URI) -verbose up

migrate-down:
	migrate -path migrations -database $(DATABASE_URI) -verbose down
