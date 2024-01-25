package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Logs basic information about each HTTP request to "app.log"
func Logging(next http.Handler) http.Handler {
	fmt.Printf("\n---- ---- ---- ----")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("Error opening log file:", err)
		}
		defer logFile.Close()

		// Add the log time of the request
		logger := log.New(logFile, "", log.Ldate|log.Ltime)
		logger.Printf("[%s] %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		fmt.Printf("\n[%s] %s", r.Method, r.RequestURI)

		// Serve the request to the next handler
		next.ServeHTTP(w, r)
	})
}
