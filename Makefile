migrateup:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5432/postgres" -verbose up

test:
	go test -v -cover ./...

.PHONY: test