migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -source=db/sqlc/store.go -destination=db/mock/store.go -package=mock

.PHONY: migrateup migratedown sqlc test server mock