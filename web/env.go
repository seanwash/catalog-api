package web

// Env facilitates injecting dependencies into route handlers.
import "database/sql"

type Env struct {
	DB *sql.DB
}
