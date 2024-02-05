MIGRATIONS_DIR="migrations"

DATABASE_URI="postgres://doadmin:AVNS_dsLxBC4W4w_dvG0ykUC@db-postgresql-auction-do-user-14102421-0.c.db.ondigitalocean.com:25060/auction"

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
