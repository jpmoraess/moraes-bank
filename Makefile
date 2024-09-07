migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -source=db/sqlc/store.go -destination=db/mock/store.go -package=mock

.PHONY: migrateup migrateup1 migratedown migratedown1 sqlc test server mock