package data

import (
	"database/sql"
	"errors"
)

var (
	// ErrRecordNotFound is a custom error that represents no record found.
	ErrRecordNotFound = errors.New("record not found")

	// ErrEditConflict is a custom error that represents there is an edit conflict in the database.
	ErrEditConflict = errors.New("edit conflict")
)

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
