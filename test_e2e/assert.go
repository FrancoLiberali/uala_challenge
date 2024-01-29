package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/stretchr/testify/assert"
)

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter

	a(&t, expected, actual, msgAndArgs...)

	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...) //nolint:goerr113 // necessary to implement assert.TestingT
}

func assertResponseStatus(resp *http.Response, status int) error {
	err := assertExpectedAndActual(assert.Equal, status, resp.StatusCode)
	if err != nil {
		respBytes, _ := io.ReadAll(resp.Body)

		return assertExpectedAndActual(assert.Equal, status, resp.StatusCode, string(respBytes))
	}

	return nil
}
