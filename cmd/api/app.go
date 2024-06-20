package main

import (
	"github.com/hayohtee/greenlight/internal/data"
	"github.com/hayohtee/greenlight/internal/jsonlog"
	"github.com/hayohtee/greenlight/internal/mailer"
	"github.com/hayohtee/greenlight/internal/vcs"
	"sync"
)

// Holds the application version number.
var version = vcs.Version()

// A type to holds all configuration settings for the app.
type config struct {
	// A network port to listen on
	port int
	// Name of the current operating environment(development, staging,production, etc.)
	env string
	// Holds configuration settings for the database.
	db struct {
		// Holds the data source name.
		dsn string
		// Holds the maximum number of open connections.
		maxOpenConns int
		// Holds the maximum number of idle connections.
		maxIdleConns int
		// Holds the time duration for idle connections.
		maxIdleTime string
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
	cors struct {
		trustedOrigins []string
	}
}

// A type to hold the dependencies for HTTP handlers, helpers,
// middlewares.
type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
}
