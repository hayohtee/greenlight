package jsonlog

import (
	"encoding/json"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// Level is a type to represent the severity level of a log entry.
type Level int8

// Initialize constants which represent specific severity level.
const (
	LevelInfo Level = iota
	LevelError
	LevelFatal
	LevelOff
)

// Return human-friendly string for the severity level.
func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

// Logger is a type which holds the output destination that the log entries
// will be written to, the minimum severity level that log entries will be
// written for, plus mutex for coordinating writes.
type Logger struct {
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

// New return a new Logger instance which writes log entries at or above
// minimum severity level to a specific output destination.
func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
	}
}

func (l *Logger) Write(message []byte) (int, error) {
	return l.print(LevelError, string(message), nil)
}

func (l *Logger) PrintInfo(message string, properties map[string]string) {
	_, _ = l.print(LevelInfo, message, properties)
}

func (l *Logger) PrintError(err error, properties map[string]string) {
	_, _ = l.print(LevelError, err.Error(), properties)
}

func (l *Logger) PrintFatal(err error, properties map[string]string) {
	_, _ = l.print(LevelFatal, err.Error(), properties)
	os.Exit(1)
}

// print is an internal method for writing the log entry.
func (l *Logger) print(level Level, message string, properties map[string]string) (int, error) {
	if level < l.minLevel {
		return 0, nil
	}

	// Define anonymous struct holding the data for the log entry
	aux := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:"message"`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Level:      level.String(),
		Time:       time.Now().Format(time.RFC3339),
		Message:    message,
		Properties: properties,
	}

	// Include a stack trace for entries at the ERROR and FATAL levels.
	if level >= LevelError {
		aux.Trace = string(debug.Stack())
	}

	// Declare a line variable for holding the actual log entry text
	var line []byte

	// Marshal anonymous struct to JSON and store it in the line variable. If there
	// was a problem creating the JSON, set the contents of the log entry to be
	// a plain-text error message instead.
	line, err := json.Marshal(aux)
	if err != nil {
		line = []byte(LevelError.String() + ": unable to marshal log message: " + err.Error())
	}

	// Lock the mutex so that no two writes to the output destination can
	// happen concurrently.
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.out.Write(append(line, '\n'))
}
