package data

import "time"

type Movie struct {
	ID        int64     // Unique integer ID for movie
	CreatedAt time.Time // Timestamp for when the movie is added to the database
	Title     string    // Movie title
	Year      int32     // Movie release year
	Runtime   int32     // Movie runtime (in minutes)
	Genres    []string  // Slice of genres for the movies (romance, comedy, etc.)
	Version   int32     // Version number starts with 1 and will be incremented each time movie info is updated
}
