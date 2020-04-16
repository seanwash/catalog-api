package main

import (
	"log"
	"net/http"
	"os"
	"time"

	internalMiddleware "github.com/seanwash/catalog-api/api/middleware"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

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

	port := os.Getenv("PORT")
	dbConn := db.Connection()
	defer dbConn.Close()

	config := generated.Config{Resolvers: &graph.Resolver{DB: dbConn}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
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

	log.Println("Server listening on port " + port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
