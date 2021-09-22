package main

import (
	"testing"
)

func TestNewCustomerResponse(t *testing.T) {
	t.Run("it returns expected Customer response with OK status", func(t *testing.T) {
		validPhoneNumber := "(237) 697151594"

		expected := CustomerResponse{
			Country:     "Cameron",
			State:       "OK",
			CountryCode: "+237",
			PhoneNumber: "697151594",
		}

		assertCustomerResponse(t, expected, *newCustomerResponse(validPhoneNumber))
	})

	t.Run("it returns expected Customer response with NOK status", func(t *testing.T) {
		invalidPhoneNumber := "(237) 6A0311634"

		expected := CustomerResponse{
			Country:     "Cameron",
			State:       "NOK",
			CountryCode: "+237",
			PhoneNumber: "6A0311634",
		}

		assertCustomerResponse(t, expected, *newCustomerResponse(invalidPhoneNumber))

	})
}
