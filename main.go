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
	router.Use(middlewares.AuthMiddleware)

	// Public routes
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/home", homeHandler).Methods("GET")
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/about", aboutHandler).Methods("GET")

	// Private routes
	router.HandleFunc("/private/protected", protectedHandler).Methods("GET")
	router.HandleFunc("/private/addWord", addWordHandler).Methods("POST")
	router.HandleFunc("/private/getWord/{word}", getWordHandler).Methods("GET")
	router.HandleFunc("/private/deleteWord/{word}", deleteWordHandler).Methods("DELETE")

	fmt.Println("Server listening on :8080...")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)


}

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

	if (strings.Replace(word, " ", "", -1) == "" || 
		len(strings.Replace(word, " ", "", -1)) < 1) {
		w.Write([]byte("Please enter a word."))
		return
	}

	entry, err := dict.Get(word)
	if err != nil {
		http.Error(w, "Error getting word from dictionary", http.StatusInternalServerError)
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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginHandler")
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
        claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires in 1 hour

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

        // Print token information for debugging
        fmt.Println("Generated Token:", tokenString)
    } else {
        log.Println("Invalid credentials")
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    }
}




func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// The request will only reach here if it passed the AuthMiddleware
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("You are authorized!"))
}



/*

// -----------------------------------------------------------------------------
// Add new items to the dictionary (the key)
func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	// Let user type the item name
	fmt.Print("Enter a word to add: ")
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	word = word[:len(word)-brLen] // Trim newline character

	// Let the user give the definition of the item
	fmt.Print("Enter the definition: ")
	definition, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	definition = definition[:len(definition)-brLen] // Trim newline character

	d.Add(word, definition)
}







// -----------------------------------------------------------------------------
// Get the definition of a given word
func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word to define: ")
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	word = word[:len(word)-brLen] // Trim newline character

	entry, err := d.Get(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Definition of '%s': %s\n", word, entry.Definition)
	}
}

// -----------------------------------------------------------------------------
// Remove an item from the dictionary
func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word to remove: ")
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	word = word[:len(word)-brLen] // Trim newline character

	// delete(d, word)
	d.Remove(word)
}


// -----------------------------------------------------------------------------
// Display the dictionary content
func actionList(d *dictionary.Dictionary) {
	// Get the content of dictionary as key value pairs
	words, entries := d.List()

	fmt.Printf("Dictionary contains %d item(s).\n", len(words))
	for _, word := range words {
		entry := entries[word]
		fmt.Printf("%s: %s\n", word, entry.Definition)
	}
}
*/



