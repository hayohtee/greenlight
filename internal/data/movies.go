package data

import "time"

type Movie struct {
	// Unique integer ID for movie.
	ID int64 `json:"id"`
	// Timestamp for when the movie is added to the database.
	CreatedAt time.Time `json:"created_at"`
	// Movie title.
	Title string `json:"title"`
	// Movie release year.
	Year int32 `json:"year"`
	// Movie runtime (in minutes).
	Runtime int32 `json:"runtime"`
	// Slice of genres for the movies (romance, comedy, etc.).
	Genres []string `json:"genres"`
	// Version number starts with 1 and will be incremented each time movie info is updated.
	Version int32 `json:"version"`
}
