include .env

GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
GONAME=$(shell basename "$(PWD)")

.PHONY: deps
sqlboil:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) sqlboiler --wipe -o pkg/models psql

.PHONY: db.migrate
db.migrate:
	migrate -path ./db/migrations -database ${DB_URL} up

.PHONY: db.rollback
db.rollback:
	migrate -path ./db/migrations -database ${DB_URL} down 1
