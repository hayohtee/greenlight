package main

import (
	"fmt"
	"net/http"
)

// logError is a generic helper for logging an error message.
func (app *application) logError(r *http.Request, err error) {
	app.logger.PrintError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

// errorResponse is a generic helper for sending JSON-formatted error messages to the client
// with the given status code.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// The serverErrorResponse logs the detailed error message and uses the
// errorResponse helper method to send a 500 Internal Server Error status code
// and JSON response (containing a generic error message) to the client.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := http.StatusText(http.StatusInternalServerError)
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse uses the errorResponse helper method to send 404 Not Found status code
// and JSON response to the client.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := http.StatusText(http.StatusNotFound)
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// methodNotAllowedResponse uses the errorResponse helper method to send a 405 Method Not Allowed
// status code and JSON response to the client.
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// badRequestResponse uses the errorResponse helper method to send a 400 Bad Request status code
// and JSON response containing the error message to the client.
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

// failedValidationResponse writes 422 Unprocessable Entity and the contents of the error as JSON response.
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

// editConflictResponse writes 409 Conflict and a message describing the error as JSON response.
func (app *application) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again."
	app.errorResponse(w, r, http.StatusConflict, message)
}

// rateLimitExceededResponse writes 429 Too Many Requests and a message describing the error
// as JSON response
func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded."
	app.errorResponse(w, r, http.StatusTooManyRequests, message)
}

// invalidCredential response writes 401 Unauthorized header and a message describing the error
// as JSON response.
func (app *application) invalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials."
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *application) invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *application) authenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *application) inactiveAccountResponse(w http.ResponseWriter, r *http.Request) {
	message := "your account must be activated to access this resource"
	app.errorResponse(w, r, http.StatusForbidden, message)
}
