package main

// Import dependencies
import (
	"fmt"
	"log"
	"time"
	"strconv"
	"strings"
	"net/http"
	"encoding/json"
	"dictionary/dictionary"
	"dictionary/middlewares"
	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
)

// Define a global dictionary instance
var dict *dictionary.Dictionary

// Define the user object structure
type User struct {
	Username string
	Password string
}

func main() {
	// Create a new instance of the dictionary from the package
	dict = dictionary.New("dictionary.json")

	// Create a main router
	router := mux.NewRouter()

	// Use the logging middleware for all routes
	router.Use(middlewares.Logging)

	// Public routes
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/home", HomeHandler).Methods("GET")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/about", AboutHandler).Methods("GET")

	// Create a subrouter for private routes
	privateRouter := router.PathPrefix("/").Subrouter()
	privateRouter.Use(middlewares.Authentication)

	// Add routes to the private router
	privateRouter.HandleFunc("/protected", protectedHandler).Methods("GET")
	privateRouter.HandleFunc("/addWord", addWordHandler).Methods("POST")
	privateRouter.HandleFunc("/getWord", getWordHandler).Methods("GET")
	privateRouter.HandleFunc("/getWord/{word}", getWordHandler).Methods("GET")
	privateRouter.HandleFunc("/listWords", listWordsHandler).Methods("GET")
	privateRouter.HandleFunc("/deleteWord", deleteWordHandler).Methods("DELETE")
	privateRouter.HandleFunc("/deleteWord/{word}", deleteWordHandler).Methods("DELETE")
	// REDIS routes
	privateRouter.HandleFunc("/addToRedis", AddToRedisHandler).Methods("POST")
	privateRouter.HandleFunc("/getFromRedis", GetFromRedisHandler).Methods("GET")
	privateRouter.HandleFunc("/getFromRedis/{word}", GetFromRedisHandler).Methods("GET")
	privateRouter.HandleFunc("/listRedis", listRedisHandler).Methods("GET")
	privateRouter.HandleFunc("/deleteFromRedis", deleteFromRedisHandler).Methods("DELETE")
	privateRouter.HandleFunc("/deleteFromRedis/{word}", deleteFromRedisHandler).Methods("DELETE")

	fmt.Println("Server listening on :8080...") // Log message
	http.Handle("/", router)                    // Default route and server instance
	http.ListenAndServe(":8080", nil)           // Start the server
}

// PUBLIC ROUTES
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Write a welcome message to the response writer
	w.Write([]byte("Welcome to the dictionary!"))
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Learn about our dictionary."))
}
// Allows the user to authenticate to server
func LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	account := User{
		Username: "cristian@mail.com",
		Password: "password123",
	}

	if username == account.Username && password == account.Password {
		// Generate JWT token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = account.Username
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
		fmt.Println("\nGenerated Token:\n",tokenString)
	} else {
		log.Println("Invalid credentials")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}


// PRIVATE ROUTES
// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// JSON LOCAL FILE HANDLERS
// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// Add a key-value pair to the json dictionary
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
	definition := requestData["definition"]
	fmt.Printf("word : %s\n", word)
	fmt.Printf("definition : %s\n", definition)

	if word == "" {
		http.Error(w, "Word is required", http.StatusBadRequest)
		return
	}
	if definition == "" {
		http.Error(w, "No definition was provided for '"+word+"'", http.StatusBadRequest)
		return
	}

	// Add the word to the dictionary
	result, err := dict.Add(word, definition)
	if err != nil {
		log.Println("Error adding '"+word+"' to the dictionary:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Send a success response
	w.Write([]byte(result))
}

// Get a requested word from the dictionary if it exists
func getWordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	if word == "" {
		http.Error(w, "Please provide a word", http.StatusBadRequest)
		return
	}

	if strings.Replace(word, " ", "", -1) == "" ||
		len(strings.Replace(word, " ", "", -1)) < 1 {
		w.Write([]byte("Please enter a word."))
		return
	}

	defs, err := dict.Get(word)
	if err != nil {
		http.Error(w, "Error getting word '%s' from dictionary.", http.StatusInternalServerError)
		return
	}

	if defs == (dictionary.Entry{}) {
		http.Error(w, "Word '"+word+"' not found in the dictionary.", http.StatusNotFound)
		return
	}

	// Convert the entry to JSON and send it as the response
	responseJSON, _ := json.Marshal(defs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

// Get access to all the words from dictionary with pagination
func listWordsHandler(w http.ResponseWriter, r *http.Request) {
	// Get pagination parameters from the query string
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")

	// Convert and get page and pageSize to integers
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		// Method 1 : return error
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
		// Method 2 : set default values
		// page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		// Method 1 : return error
		// http.Error(w, "Invalid pageSize parameter", http.StatusBadRequest)
		// return
		// Method 2 : set default values
		pageSize = 10
	}

	// Get the data from dictionary
	words, entries := dict.List()

	// Case: less words than words number indicated
	if pageSize > len(words) {
		pageSize = len(words) // Return all words
	}
	// Calculate the start and end indices based on pagination parameters
	start := (page - 1) * pageSize
	end := start + pageSize
	if end > pageSize {
		end = len(words) // Return remaining words
	}

	// Ensure that start and end are within bounds
	if start < 0 || start >= len(words) || end > len(words) {
		http.Error(w, "Invalid pagination parameters", http.StatusBadRequest)
		return
	}

	// Extract the entries for the requested page interval
	pagedEntries := make(map[string]dictionary.Entry)
	for _, word := range words[start:end] {
		pagedEntries[word] = entries[word]
	}

	// Create the response JSON structure
	responseJSON := map[string]interface{}{
		"request": map[string]interface{}{
			"page":     page,
			"pageSize": pageSize,
			"start":    start,
			"end":      end,
		},
		"data": pagedEntries,
	}

	// Convert the response JSON to a string
	response, err := json.Marshal(responseJSON)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	// Set the response content type and send the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Delete a given word with its definition if it exists
func deleteWordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	if strings.Replace(word, " ", "", -1) == "" ||
		len(strings.Replace(word, " ", "", -1)) < 1 {
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
// Allow to test the authentication middleware
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// The request will only reach here if it passed the AuthMiddleware
	log.SetPrefix("INFO: ")
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("You are authorized!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("You are authorized!"))
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// REDIS HANDLERS
// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// CREATE request

func AddToRedisHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into a map
	var requestData map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Extract word from the request data
	word := requestData["word"]
	fmt.Printf("word : %s\n", word)
	if word == "" {
		http.Error(w, "Word is required", http.StatusBadRequest)
		return
	}

	// Extract definition from the request data
	definition := requestData["definition"]
	fmt.Printf("definition : %s\n", definition)
	if definition == "" {
		http.Error(w, "No definition was provided for '"+word+"'", http.StatusBadRequest)
		return
	}

	// Save data to Redis
	result, err := dict.SaveToRedis(word, definition)
	if err != nil {
		log.Println("Error adding '"+word+"' to the dictionary:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.Write([]byte(result))
}

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
// GET : a word from the dictionary if it exists
func GetFromRedisHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	if word == "" {
		http.Error(w, "Please provide a word", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(word) == "" ||
		len(strings.TrimSpace(word)) < 1 {
		w.Write([]byte("Please enter a word."))
		return
	}

	defs, err := dict.GetFromRedis(word)
	if err != nil {
		log.Printf("Error getting word '%s' from dictionary.\n", word)
		log.Println("Error message : ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if defs == (dictionary.Entry{}) {
		http.Error(w, "Word '"+word+"' not found in the dictionary.", http.StatusNotFound)
		return
	}

	// // Convert the entry to JSON and send it as the response
	responseJSON, _ := json.Marshal(defs)
	fmt.Printf("%s : %s\n", word, defs.Definition)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
// GET LIST : access all the words from dictionary with pagination
func listRedisHandler(w http.ResponseWriter, r *http.Request) {
	// Get pagination parameters from the query string
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")

	// Convert and get page and pageSize to integers
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		// Method 1 : return error
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
		// Method 2 : set default values
		// page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		// Method 1 : return error
		// http.Error(w, "Invalid pageSize parameter", http.StatusBadRequest)
		// return
		// Method 2 : set default values
		pageSize = 10
	}

	// Get the data from the Redis database
	words, err := dict.ListRedis(page, pageSize)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving data from Redis: %v", err), http.StatusInternalServerError)
		return
	}

	// Case: less words than words number indicated
	if pageSize > len(words) {
		pageSize = len(words) // Return all words
	}
	// Calculate the start and end indices based on pagination parameters
	start := (page - 1) * pageSize
	end := start + pageSize
	if end > pageSize {
		end = len(words) // Return remaining words
	}

	// Ensure that start and end are within bounds
	if start < 0 || start >= len(words) || end > len(words) {
		http.Error(w, "Invalid pagination parameters", http.StatusBadRequest)
		return
	}

	// Extract the entries for the requested page interval
	pagedEntries := make(map[string]dictionary.Entry)
	counter := 0
	for word, entry := range words {
		if counter >= start && counter < end {
			pagedEntries[word] = entry
		}
		counter++
	
		if counter >= end {
			break
		}
	}

	// Create the response JSON structure
	responseJSON := map[string]interface{}{
		"request": map[string]interface{}{
			"page":     page,
			"pageSize": pageSize,
			"start":    start,
			"end":      end,
		},
		"data": pagedEntries,
	}

	// Convert the response JSON to a string
	response, err := json.Marshal(responseJSON)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
// DELETE : a word with its definition if it exists
func deleteFromRedisHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	if strings.TrimSpace(word) == "" || len(strings.TrimSpace(word)) < 1 {
		w.Write([]byte("Please enter a word."))
		return
	}

	if err := dict.DeleteFromRedis(word); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Send a success response
	w.Write([]byte(word + " deleted successfully."))
}
