package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/seanwash/catalog-api/web"

	"github.com/seanwash/catalog-api/db"

	"github.com/joho/godotenv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")
	dbConn := db.Connection()
	env := &web.Env{DB: dbConn}
	defer dbConn.Close()

	r := chi.NewRouter()

	// Middleware
	//
	// Injects a unique request ID into each request's context.
	r.Use(middleware.RequestID)
	// Basic request logging, Logs the start/end of each request to stdout.
	r.Use(middleware.Logger)
	// Automatically recover from panics, log them, and return an HTTP 500 status.
	r.Use(middleware.Recoverer)
	// Stores a URL's extension in each request context.
	r.Use(middleware.URLFormat)
	// Always set content-type JSON header.
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further processing
	// should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Routes
	registerRoutes(env, r)

	// Serve
	log.Println("Application started on port :" + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func registerRoutes(env *web.Env, r *chi.Mux) {
	// Misc
	r.Get("/healthz", env.HealthCheck)

	// Track Routing
	r.Get("/tracks", env.TracksIndex)
	r.Get("/tracks/{id}", env.TracksShow)
	r.Post("/tracks", env.TracksCreate)
	r.Put("/tracks/{id}", env.TracksUpdate)
	r.Delete("/tracks/{id}", env.TracksDelete)
}
