package main

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"io"
	"log"
	"marcode.io/phone-numbers/pkg/data"
	"testing"
)

type StubCustomerModel struct {
	db map[string]*data.Customer
}

var StubDatabase = map[string]*data.Customer{}

func (s StubCustomerModel) GetAll(code string) ([]*data.Customer, error) {
	customers := make([]*data.Customer, 0)
	if code == "" {
		return []*data.Customer{
			{
				ID:          1,
				Name:        "Customer1",
				PhoneNumber: "(251) 911168450",
			},
			{
				ID:          2,
				Name:        "Customer2",
				PhoneNumber: "(251) 91A1168450",
			},
			{
				ID:          3,
				Name:        "Customer3",
				PhoneNumber: "(258) 847651504",
			},
			{
				ID:          4,
				Name:        "Customer4",
				PhoneNumber: "(258) 042423566",
			},
			{
				ID:          5,
				Name:        "Customer5",
				PhoneNumber: "(237) 6A0311634",
			},
			{
				ID:          6,
				Name:        "Customer6",
				PhoneNumber: "(237) 6B0311634",
			},
			{
				ID:          7,
				Name:        "Customer7",
				PhoneNumber: "(237) 697151594",
			},
		}, nil
	} else if code == "(251)" {
		return []*data.Customer{
			{
				ID:          1,
				Name:        "Customer1",
				PhoneNumber: "(251) 911168450",
			},
			{
				ID:          2,
				Name:        "Customer2",
				PhoneNumber: "(251) 91A1168450",
			},
		}, nil

	} else if code == "(258)" {
		return []*data.Customer{
			{
				ID:          3,
				Name:        "Customer3",
				PhoneNumber: "(258) 847651504",
			},
			{
				ID:          4,
				Name:        "Customer4",
				PhoneNumber: "(258) 042423566",
			},
		}, nil
	} else if code == "(237)" {
		return []*data.Customer{
			{
				ID:          5,
				Name:        "Customer5",
				PhoneNumber: "(237) 6A0311634",
			},
			{
				ID:          6,
				Name:        "Customer6",
				PhoneNumber: "(237) 6B0311634",
			},
			{
				ID:          7,
				Name:        "Customer7",
				PhoneNumber: "(237) 697151594",
			},
		}, nil
	}
	return customers, nil
}

// newTestApplication returns an instance of our application
// composing mocked dependencies
func newTestApplication() *application {
	return &application{
		logger: log.New(io.Discard, "", 0),
		models: data.Models{Customers: StubCustomerModel{db: StubDatabase}},
	}
}

func assertCustomerResponse(t testing.TB, expected, actual CustomerResponse) {
	t.Helper()

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("expected the following Customer Response %v, but got %v, and difference as follows %v", expected, actual, diff)
	}
}

func assertResponseCode(t testing.TB, expected, actual int) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected status %d, but got status %d", expected, actual)
	}
}

func assertContentTypeHeader(t testing.TB, expected, actual string) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %s Content-Type Header, but got %s", expected, actual)
	}
}

func assertResponseRecordsCount(t testing.TB, expected, actual int) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected %d records but got %d", expected, actual)
	}
}

func decodeResponse(t testing.TB, body io.Reader, dto interface{}) {

	err := json.NewDecoder(body).Decode(&dto)
	if err != nil {
		t.Fatalf("couldn't decode json response, %v", err)
	}
}

func assertResponseBody(t testing.TB, expected, actual string) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected the following response %q, but got %q", expected, actual)
	}
}
