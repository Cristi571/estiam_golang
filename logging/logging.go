package logging

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

// Flusher interface for manually flushing log buffer
type Flusher interface {
	Flush() error
}

// Middleware function for logging requests
func LoggingMiddleware(next http.Handler) http.Handler {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	defer logFile.Close()

	// Wrap logFile with a FlushableFile type
	flushableFile := &FlushableFile{logFile, sync.Mutex{}}

	logger := log.New(flushableFile, "", log.Ldate|log.Ltime)
	logger.Printf("Logging 1")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Method : [%s]\n", r.Method)
		fmt.Printf("RequestURI : %s\n", r.RequestURI)
		fmt.Printf("RemoteAddr : %s\n", r.RemoteAddr)

		logger.Printf("Logging 2 [%s] %s %s", r.Method, r.RequestURI, r.RemoteAddr)

		// Serve the request to the next handler
		next.ServeHTTP(w, r)

		// Log after handling the request
		logger.Printf("Logging 3 [%s] %s %s", r.Method, r.RequestURI, r.RemoteAddr)

		// Manually flush the log data to the file using the Flusher interface
		flushableFile.Flush()
	})
}

// FlushableFile is a wrapper around *os.File that implements the Flusher interface
type FlushableFile struct {
	*os.File
	mu sync.Mutex
}

// Flush manually flushes the log buffer to the file
func (f *FlushableFile) Flush() {
	f.mu.Lock()
	f.File.Sync()
	f.mu.Unlock()
}
