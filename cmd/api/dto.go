package main

import (
	"fmt"
	"marcode.io/phone-numbers/pkg/data"
	"marcode.io/phone-numbers/pkg/validators"
	"strings"
)

type CustomerResponse struct {
	Country     string `json:"country"`
	State       string `json:"state"`
	CountryCode string `json:"countryCode"`
	PhoneNumber string `json:"phoneNumber"`
}

func newCustomerResponse(phoneNumber string) *CustomerResponse {
	splittedPhoneNumber := strings.Split(phoneNumber, " ")
	validator := validators.PhoneNumberValidatorFactory(phoneNumber)

	customerResponse := &CustomerResponse{
		Country:     data.CountryCodes[splittedPhoneNumber[0]],
		CountryCode: formatCountryCode(splittedPhoneNumber[0]),
		PhoneNumber: splittedPhoneNumber[1],
	}

	if validator.ValidatePhoneNumber(phoneNumber) {
		customerResponse.State = "OK"
	} else {
		customerResponse.State = "NOK"
	}

	return customerResponse
}

func formatCountryCode(countryCode string) string {
	n := len(countryCode)
	return fmt.Sprintf("+%s", countryCode[1:n-1])
}
