package main

import (
	"github.com/rs/cors"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	router.Handle("/api", http.HandlerFunc(app.filterHandler))
	router.Handle("/api/health", http.HandlerFunc(app.healthCheckHandler))

	return c.Handler(router)
}
