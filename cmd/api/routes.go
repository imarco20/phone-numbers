package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.Handle("/", http.HandlerFunc(app.healthCheckHandler))

	return router
}
