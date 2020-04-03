package web

// Env facilitates injecting dependencies into route handlers.
import "database/sql"

// Env facilitates dependency injection for web specific dependencies.
type Env struct {
	DB *sql.DB
}
