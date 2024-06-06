package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/hayohtee/greenlight/internal/data"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

// Holds the application version number.
const version = "1.0.0"

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
}

// A type to hold the dependencies for HTTP handlers, helpers,
// middlewares.
type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config

	// Reads the value of port and env commandline flags into the config struct
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// Reads the DSN value from db-dsn command-line flag into the config struct.
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL DSN")

	// Reads the connection pool settings from the command-line flags into the config struct
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max idle time")
	flag.Parse()

	// Initialize a new logger which writes message to the standard output stream, prefixed with
	// current date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(db)

	logger.Print("database connection established")

	// Create an instance of application struct
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	// Create HTTP server with timeouts and listen on the port provided in the config struct.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	// Create context with a 5-seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
