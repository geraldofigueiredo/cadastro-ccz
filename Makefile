migrate-up:
	migrate -path internal/db/migrations -database "postgresql://postgres:postgres@localhost:5432/ccz?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/db/migrations -database "postgresql://postgres:postgres@localhost:5432/ccz?sslmode=disable" -verbose down