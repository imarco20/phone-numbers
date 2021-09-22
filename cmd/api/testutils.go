package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func assertCustomerResponse(t testing.TB, expected, actual CustomerResponse) {
	t.Helper()

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("expected the following Customer Response %v, but got %v, and difference as follows %v", expected, actual, diff)
	}
}
