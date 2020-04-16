-include .env

gen.boil:
	sqlboiler --wipe --no-hooks -o models psql

gen.gql:
	go run github.com/99designs/gqlgen

build.migrate:
	go build -o ./bin/migrate ./cmd/migrate

db.up:
	goose -dir ./db/migrations postgres ${DATABASE_URL} up

db.down:
	goose -dir ./db/migrations postgres ${DATABASE_URL} down

db.status:
	goose -dir ./db/migrations postgres ${DATABASE_URL} status

db.reset:
	goose -dir ./db/migrations postgres ${DATABASE_URL} reset

db.version:
	goose -dir ./db/migrations postgres ${DATABASE_URL} version

run:
	go run cmd/server/main.go
