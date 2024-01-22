package main

// Import dependencies
import (
	// "bufio"
	"fmt"
	"log"
	"time"
	"net/http"
	"strings"
	"encoding/json"
	// "os"
	// "strconv"
	"dictionary/dictionary"
	"dictionary/middlewares"
	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
)
// Define the new line lenght for different env and machines
// 2 windows | 1 iOS
// var brLen int = 1
// Define a global dictionary instance
var dict *dictionary.Dictionary

// Define the user object structure
type User struct {
	Username string
	Password string
}

func main() {
    filePath := "dictionary.json"
    // Create a new instance of the dictionary from the package
    dict = dictionary.New(filePath)

    // Create a new router
    router := mux.NewRouter()

    // Use the logging middleware for all routes
    router.Use(middlewares.LoggingMiddleware)

    // Public routes
    router.HandleFunc("/", homeHandler).Methods("GET")
    router.HandleFunc("/home", homeHandler).Methods("GET")
    router.HandleFunc("/login", loginHandler).Methods("POST")
    router.HandleFunc("/about", aboutHandler).Methods("GET")

    // Create a subrouter for private routes without /private prefix
    privateRouter := router.PathPrefix("/").Subrouter()
    privateRouter.Use(middlewares.AuthMiddleware)

    // Private routes without /private prefix
    privateRouter.HandleFunc("/protected", protectedHandler).Methods("GET")
    privateRouter.HandleFunc("/addWord", addWordHandler).Methods("POST")
    privateRouter.HandleFunc("/getWord/{word}", getWordHandler).Methods("GET")
    privateRouter.HandleFunc("/deleteWord/{word}", deleteWordHandler).Methods("DELETE")

    fmt.Println("Server listening on :8080...")
    http.Handle("/", router)
    http.ListenAndServe(":8080", nil)
}

// PUBLIC ROUTES
// Home page - Handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Write a welcome message to the response writer
	w.Write([]byte("Welcome to the home page!"))
}
// About page - Handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	// Write a message about the page to the response writer
	w.Write([]byte("About us page"))
}
// Allows the user to authenticate to server
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginHandler")
	// Check if a session cookie exists, only browsers
    _, err := r.Cookie("session")
    if err == nil {
        // Session cookie exists, user is already authenticated
        log.Println("User is already authenticated")
        http.Error(w, "User is already authenticated", http.StatusConflict)
        return
    }


    // Check credentials (replace this with your authentication logic)
    username := r.FormValue("username")
    password := r.FormValue("password")

    demoUser := User{
        Username: "demo",
        Password: "password123",
    }

    if username == demoUser.Username && password == demoUser.Password {
        // Generate JWT token
        token := jwt.New(jwt.SigningMethodHS256)
        claims := token.Claims.(jwt.MapClaims)
        claims["username"] = demoUser.Username
        claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 1 day

        tokenString, err := token.SignedString(middlewares.JWTSecret)
        if err != nil {
            log.Println("Error generating token:", err)
            http.Error(w, "Error generating token", http.StatusInternalServerError)
            return
        }

        // Include the token in the response headers
        w.Header().Set("Authorization", "Bearer "+tokenString)
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Login successful"))

		// Set the token as a cookie to simulate a session
        http.SetCookie(w, &http.Cookie{
            Name:    "session",
            Value:   tokenString,
            Expires: time.Now().Add(time.Hour * 24),
        })

        // Print token information for debugging
        fmt.Println("Generated Token:", tokenString)
    } else {
        log.Println("Invalid credentials")
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    }
}

// PRIVATE ROUTES
// Handler for the "addWord" route (POST method)
func addWordHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into a map
	var requestData map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Extract word and definition from the request data
	word := requestData["word"]
	fmt.Printf("word : %s\n", word)
	definition := requestData["definition"]
	fmt.Printf("definition : %s\n", definition)

	// Add the word to the dictionary
	result, err := dict.Add(word, definition)
	if err != nil {
		fmt.Println("Error:", err)
		w.Write([]byte("Could not add word "+word+"."))
		return
	}
	// Send a success response
	w.Write([]byte(result))
}

// Handler for the "getWord" route (GET method)
func getWordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	if word == "" {
        http.Error(w, "Please provide a word", http.StatusBadRequest)
        return
    }

	if (strings.Replace(word, " ", "", -1) == "" || 
		len(strings.Replace(word, " ", "", -1)) < 1) {
		w.Write([]byte("Please enter a word."))
		return
	}

	entry, err := dict.Get(word)
	if err != nil {
		log.Printf("Error getting word '%s' from dictionary.\n", word)
		log.Println("Error message : ", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
	}

	if entry == (dictionary.Entry{}) {
        http.Error(w, "Word '"+word+"' not found in the dictionary.", http.StatusNotFound)
        return
    }

	// Convert the entry to JSON and send it as the response
	responseJSON, _ := json.Marshal(entry)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

// Handler for the "deleteWord" route (DELETE method)
func deleteWordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	if (strings.Replace(word, " ", "", -1) == "" || 
		len(strings.Replace(word, " ", "", -1)) < 1) {
		w.Write([]byte("Please enter a word to delete."))
		return
	}

	result, err := dict.Remove(word)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Send a success response
	w.Write([]byte(result))
}




// Only for the authorisation middleware test purpose
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// The request will only reach here if it passed the AuthMiddleware
	log.SetPrefix("INFO: ")
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("You are authorized!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("You are authorized!"))
}




