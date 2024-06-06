package data

import (
	"database/sql"
	"errors"
)

// ErrRecordNotFound is a custom error that represents no record found.
var ErrRecordNotFound = errors.New("record not found")

// Models is a struct that wraps all the database models.
type Models struct {
	Movies MovieModel
}

// NewModels returns an initialized Models struct.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
