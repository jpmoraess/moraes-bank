## Moraes Bank

## migrate

curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/bin/
which migrate

migrate create -ext sql -dir db/migration -seq init_schema

## sqlc

sqlc generate

## viper

go get github.com/spf13/viper

## go-mock

go install go.uber.org/mock/mockgen@latest

mockgen -source=db/sqlc/store.go -destination=db/mock/store.go -package=mock