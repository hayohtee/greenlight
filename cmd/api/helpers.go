package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// readIDParam retrieve the "id" URL parameter from the current context, then
// convert it to an integer and return it.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, err
	}
	return id, nil
}
