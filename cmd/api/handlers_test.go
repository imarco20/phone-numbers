package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFilterHandler(t *testing.T) {
	app := newTestApplication()

	t.Run("it returns status 200 OK and application/json Content-Type Header", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api", nil)
		response := httptest.NewRecorder()

		app.filterHandler(response, request)

		assertResponseCode(t, http.StatusOK, response.Code)
		assertContentTypeHeader(t, "application/json", response.Header().Get("Content-Type"))

	})
	t.Run("it returns all customers with code (251)", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api", nil)

		q := request.URL.Query()
		q.Add("code", "(251)")
		request.URL.RawQuery = q.Encode()

		response := httptest.NewRecorder()

		app.filterHandler(response, request)

		var jsonResponse = map[string][]CustomerResponse{}

		decodeResponse(t, response.Body, &jsonResponse)

		expected := []CustomerResponse{
			{
				Country:     "Ethiopia",
				State:       "OK",
				CountryCode: "+251",
				PhoneNumber: "911168450",
			},
			{
				Country:     "Ethiopia",
				State:       "NOK",
				CountryCode: "+251",
				PhoneNumber: "91A1168450",
			},
		}

		actual := jsonResponse["customers"]
		assertResponseRecordsCount(t, len(expected), len(actual))
		for i := 0; i < len(expected); i++ {
			assertCustomerResponse(t, expected[i], actual[i])
		}

	})
	t.Run("it returns 1 customer with code (258) and state OK", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api", nil)

		q := request.URL.Query()
		q.Add("code", "(258)")
		q.Add("state", "OK")
		request.URL.RawQuery = q.Encode()

		response := httptest.NewRecorder()

		app.filterHandler(response, request)

		var jsonResponse = map[string][]CustomerResponse{}

		decodeResponse(t, response.Body, &jsonResponse)

		expected := []CustomerResponse{
			{
				Country:     "Mozambique",
				State:       "OK",
				CountryCode: "+258",
				PhoneNumber: "847651504",
			},
		}

		actual := jsonResponse["customers"]
		assertResponseRecordsCount(t, len(expected), len(actual))
		for i := 0; i < len(expected); i++ {
			assertCustomerResponse(t, expected[i], actual[i])
		}
	})
	t.Run("it returns 2 customers with code (237) and state NOK", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api", nil)

		q := request.URL.Query()
		q.Add("code", "(237)")
		q.Add("state", "NOK")
		request.URL.RawQuery = q.Encode()

		response := httptest.NewRecorder()

		app.filterHandler(response, request)

		var jsonResponse = map[string][]CustomerResponse{}

		decodeResponse(t, response.Body, &jsonResponse)

		expected := []CustomerResponse{
			{
				Country:     "Cameron",
				State:       "NOK",
				CountryCode: "+237",
				PhoneNumber: "6A0311634",
			},
			{
				Country:     "Cameron",
				State:       "NOK",
				CountryCode: "+237",
				PhoneNumber: "6B0311634",
			},
		}

		actual := jsonResponse["customers"]
		assertResponseRecordsCount(t, len(expected), len(actual))
		for i := 0; i < len(expected); i++ {
			assertCustomerResponse(t, expected[i], actual[i])
		}
	})
	t.Run("it returns 4 customers with state NOK", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api", nil)

		q := request.URL.Query()
		q.Add("state", "NOK")
		request.URL.RawQuery = q.Encode()

		response := httptest.NewRecorder()

		app.filterHandler(response, request)

		var jsonResponse = map[string][]CustomerResponse{}

		decodeResponse(t, response.Body, &jsonResponse)

		expected := []CustomerResponse{
			{
				Country:     "Ethiopia",
				State:       "NOK",
				CountryCode: "+251",
				PhoneNumber: "91A1168450",
			},
			{
				Country:     "Mozambique",
				State:       "NOK",
				CountryCode: "+258",
				PhoneNumber: "042423566",
			},

			{
				Country:     "Cameron",
				State:       "NOK",
				CountryCode: "+237",
				PhoneNumber: "6A0311634",
			},
			{
				Country:     "Cameron",
				State:       "NOK",
				CountryCode: "+237",
				PhoneNumber: "6B0311634",
			},
		}

		actual := jsonResponse["customers"]
		assertResponseRecordsCount(t, len(expected), len(actual))
		for i := 0; i < len(expected); i++ {
			assertCustomerResponse(t, expected[i], actual[i])
		}
	})
	t.Run("it returns 3 customers with state OK", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api", nil)

		q := request.URL.Query()
		q.Add("state", "OK")
		request.URL.RawQuery = q.Encode()

		response := httptest.NewRecorder()

		app.filterHandler(response, request)

		var jsonResponse = map[string][]CustomerResponse{}

		decodeResponse(t, response.Body, &jsonResponse)

		expected := []CustomerResponse{
			{
				Country:     "Ethiopia",
				State:       "OK",
				CountryCode: "+251",
				PhoneNumber: "911168450",
			},
			{
				Country:     "Mozambique",
				State:       "OK",
				CountryCode: "+258",
				PhoneNumber: "847651504",
			},
			{
				Country:     "Cameron",
				State:       "OK",
				CountryCode: "+237",
				PhoneNumber: "697151594",
			},
		}

		actual := jsonResponse["customers"]
		assertResponseRecordsCount(t, len(expected), len(actual))
		for i := 0; i < len(expected); i++ {
			assertCustomerResponse(t, expected[i], actual[i])
		}
	})
}

func TestHealthCheckHandler(t *testing.T) {
	app := newTestApplication()

	t.Run("it returns status 200 OK", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api/health", nil)
		response := httptest.NewRecorder()

		app.healthCheckHandler(response, request)

		var jsonResponse = map[string]string{}
		decodeResponse(t, response.Body, &jsonResponse)
		assertResponseCode(t, http.StatusOK, response.Code)
		assertResponseBody(t, "the application is working properly", jsonResponse["health"])
	})
}
