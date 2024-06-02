package main

import (
	"flag"
	"fmt"
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
}

// A type to hold the dependencies for HTTP handlers, helpers,
// middlewares.
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// Reads the value of port and env commandline flags into the config struct
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new logger which writes message to the standard output stream, prefixed with
	// current date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Create an instance of application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)

	// Create HTTP server with timeouts and listen on the port provided in the config struct.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
