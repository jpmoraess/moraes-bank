migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres" -verbose up

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrateup sqlc test