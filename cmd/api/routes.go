package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.Handle("/api", http.HandlerFunc(app.filterHandler))
	router.Handle("/api/health", http.HandlerFunc(app.healthCheckHandler))

	return router
}
