package main

import (
	"fmt"
	"net/http"
)

// logError is a generic helper for logging an error message.
func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
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
