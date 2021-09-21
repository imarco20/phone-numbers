package main

import (
	"net/http"
)

func (app *application) filterHandler(w http.ResponseWriter, r *http.Request) {
	customers, err := app.models.Customers.GetAll("")
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"customers": customers})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// healthCheckHandler handles requests to check the application is up and running
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	err := app.writeJSON(w, http.StatusOK, envelope{"health": "the application is working properly"})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
