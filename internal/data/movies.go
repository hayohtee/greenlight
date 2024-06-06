package data

import (
	"database/sql"
	"github.com/hayohtee/greenlight/internal/validator"
	"time"
)

type Movie struct {
	// Unique integer ID for movie.
	ID int64 `json:"id"`
	// Timestamp for when the movie is added to the database.
	CreatedAt time.Time `json:"-"`
	// Movie title.
	Title string `json:"title"`
	// Movie release year.
	Year int32 `json:"year,omitempty"`
	// Movie runtime (in minutes).
	Runtime Runtime `json:"runtime,omitempty"`
	// Slice of genres for the movies (romance, comedy, etc.).
	Genres []string `json:"genres,omitempty"`
	// Version number starts with 1 and will be incremented each time movie info is updated.
	Version int32 `json:"version"`
}

// ValidateMovie adds validation check on the movie.
func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater or equal to 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}

// MovieModel is a struct which wraps a sql.DB connection pool.
type MovieModel struct {
	DB *sql.DB
}

// Insert a movie into the database.
func (m MovieModel) Insert(movie *Movie) error {
	return nil
}

// Get a specific movie from the database or return an error.
func (m MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

// Update specific record in the movie database.
func (m MovieModel) Update(movie *Movie) error {
	return nil
}

// Delete a specific movie from the database or return an error.
func (m MovieModel) Delete(id int64) error {
	return nil
}
