package main

import (
	"context"
	"github.com/hayohtee/greenlight/internal/data"
	"net/http"
)

// Define a custom contextKey type, with the underlying type string.
type contextKey string

// Represent a key for storing and retrieving user information in the context.
const userContextKey = contextKey("user")

// contextSetup returns a new copy of the request with the provided User struct
// added to the context.
func (app *application) contextSetup(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

// contextGetUser retrieves the User struct from the request context.
func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
