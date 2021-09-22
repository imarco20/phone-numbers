package main

import (
	"net/http"
)

// filterHandler handles requests to filter phone numbers in the database by Country Code and valid State
func (app *application) filterHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	customers, err := app.models.Customers.GetAll(code)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	response := make([]*CustomerResponse, 0)

	for _, customer := range customers {

		customerResponse := newCustomerResponse(customer.PhoneNumber)
		if state == customerResponse.State || state == "" {
			response = append(response, customerResponse)
		}
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"customers": response})
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
