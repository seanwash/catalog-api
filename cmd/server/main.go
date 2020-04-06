package main

import (
	"log"
	"net/http"
	"os"
	"time"

	internalMiddleware "github.com/seanwash/catalog-api/api/middleware"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/joho/godotenv"
	"github.com/seanwash/catalog-api/db"
	"github.com/seanwash/catalog-api/graph"
	"github.com/seanwash/catalog-api/graph/generated"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// TODO: Move this to a middleware so that it can be accessed through context?
	port := os.Getenv("PORT")
	dbConn := db.Connection()
	defer dbConn.Close()
	if err != nil {
		log.Fatal(err)
	}

	config := generated.Config{Resolvers: &graph.Resolver{DB: dbConn}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	r := chi.NewRouter()

	// A good base middleware stack
	// Adds a unique ID for each request. Useful for logging.
	r.Use(middleware.RequestID)
	// Adds remote IP addr for each request. Useful for logging.
	r.Use(middleware.RealIP)
	// Adds basic logging for each request. Reports timing, RequestID, and remote IP.
	r.Use(middleware.Logger)
	// Recovers from panics and logs them. Useful for logging, error tracking.
	r.Use(middleware.Recoverer)
	// Adds a route for responding to uptime reporters and zero downtime deployment checks.
	r.Use(middleware.Heartbeat("/healthz"))
	// Sets an optional User into the request context based on a API key header.
	r.Use(internalMiddleware.Auth())

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Routes
	r.Handle("/", playground.Handler("Catalog API", "/graphql"))
	r.Handle("/graphql", srv)

	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
