package main

import (
	"fmt"
	"net/http"
)

// healthcheckHandler write application information as a fixed-format json response
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(js))
}
