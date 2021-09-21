package main

import "net/http"

// healthCheckHandler handles requests to check the application is up and running
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	err := app.writeJSON(w, http.StatusOK, envelope{"health": "the application is working properly"})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
