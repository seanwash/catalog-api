package db

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/volatiletech/sqlboiler/boil"

	_ "github.com/lib/pq"
)

func Connection() *sql.DB {
	dbUrl := os.Getenv("DB_URL")
	logSqlStr := os.Getenv("LOG_SQL")
	logSql, err := strconv.ParseBool(logSqlStr)
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// Enable/Disable SQL query logging. Useful for debugging and performance
	// insights during development.
	boil.DebugMode = logSql

	return db
}
