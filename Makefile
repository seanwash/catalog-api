-include .env

boil:
	sqlboiler --wipe --no-hooks -o models psql

gql:
	go run github.com/99designs/gqlgen

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
