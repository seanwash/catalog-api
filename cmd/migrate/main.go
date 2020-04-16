// Builds a custom Goose binary that includes any of the Go based migrations from db/migrations.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/pressly/goose"

	"github.com/seanwash/catalog-api/db"
	// Required so that the Go migrations are bundled into the binary. https://github.com/pressly/goose#go-migrations
	_ "github.com/seanwash/catalog-api/db/migrations"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./db/migrations", "directory with migration files")
)

const dialect = "postgres"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	dbConn := db.Connection()
	defer dbConn.Close()

	err = goose.SetDialect(dialect)
	if err != nil {
		log.Fatal(err)
	}

	if err = goose.Run(command, dbConn, *dir, args[1:]...); err != nil {
		log.Fatalf("migrate run: %v", err)
	}
}
