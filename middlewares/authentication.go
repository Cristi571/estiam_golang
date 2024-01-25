// middlewares.go
package middlewares

import (
	_ "fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var JWTSecret = []byte("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU5NDE0NjMsInVzZXJuYW1lIjoiZGVtbyJ9.mHkQYWTMPpVJgAaPLdGejty8rgsl6NL5581j5QCCrgo")

// AuthMiddleware checks the Authorization header for a valid JWT token
func Authentication(next http.Handler) http.Handler {
	// log.Print("[PROTECTED] - ")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the Authorization header
		tokenString := extractToken(r)

		log.SetPrefix("\n[PROTECTED] | ")
		if tokenString == "" {
			// log.SetPrefix("WARNING: ")
			log.SetFlags(log.Ldate | log.Ltime)
			log.Println("Unauthorized: Missing token")

			http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
			return
		}

		// Validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})

		if err != nil || !token.Valid {
			// log.SetPrefix("\n[PROTECTED] | ")
			log.SetFlags(log.Ldate | log.Ltime)
			log.Println("Unauthorized: Invalid token")

			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}
		// log.SetPrefix("\n[PROTECTED]: ")
		log.SetFlags(log.Ldate | log.Ltime)
		log.Println("Authorized.")
		
		next.ServeHTTP(w, r)
	})
}

func extractToken(r *http.Request) string {
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        return ""
    }

    authHeaderParts := strings.Split(authHeader, " ")
    if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
        return ""
    }

    return authHeaderParts[1]
}
