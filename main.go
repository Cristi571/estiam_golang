package main

// Import dependencies
import (
	// "bufio"
	"fmt"
	"strings"
	"encoding/json"
	// "os"
	// "strconv"
	"dictionary/dictionary"
	"dictionary/logging"
	"net/http"
	"github.com/gorilla/mux"
)
// Define the new line lenght for different env and machines
// 2 windows | 1 iOS
// var brLen int = 1
// Define a global dictionary instance
var dict *dictionary.Dictionary

func main() {
	filePath := "dictionary.json"
	// Create a new instance of the dictionary from the package
	dict = dictionary.New(filePath)


	// Create a new instance of the Gorilla Mux router
	router := mux.NewRouter()

	// Use the logging middleware for all routes
	router.Use(logging.LoggingMiddleware)

	// Define a handler for the home page
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/home", homeHandler).Methods("GET")
	// Define a handler for the about page
	router.HandleFunc("/about", aboutHandler).Methods("GET")

	// 1. Une pour ajouter une entrée au dictionnaire (POST)
	router.HandleFunc("/addWord", addWordHandler).Methods("POST")
	// 2. Une pour récupérer une définition par mot (GET)
	router.HandleFunc("/getWord/{word}", getWordHandler).Methods("GET")
	// 3. Une pour supprimer une entrée par mot (DELETE)
	router.HandleFunc("/deleteWord/{word}", deleteWordHandler).Methods("DELETE")
	

	// Start the server on port 8080
	fmt.Println("Server listening on :8080...")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)



	/*
	// Set the app running status
	var appRunning = true
	// Make the app running loop
	for appRunning {
		// Store the menu oprions in an array
		buttons := [5]int{0, 1, 2, 3, 4}
		// Display menu navigation instructions for user
		fmt.Println("+-------------------------------------+")
		fmt.Println("1 : Add new item to dictionary")
		fmt.Println("2 : Get definition of a word")
		fmt.Println("3 : Delete an item from the dictionary")
		fmt.Println("4 : Display the dictionary content")
		fmt.Println("0 : Quit the app")
		fmt.Println("+-------------------------------------+")
		
		// ---------------------------------------------------------------------
		// Let the user to make a choice
		fmt.Print("Your choice: ")
		// Get the user input with a scanner that reads from os.Stdin
		reader := bufio.NewReader(os.Stdin)
		// Get the text that was read
		inputText, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// ---------------------------------------------------------------------
		// Convert the input text to an integer
		nextStep, err := strconv.Atoi(inputText[:len(inputText)-brLen]) // Trim newline character
		// Check if user printed an string typed valid integer
		if err != nil {
			fmt.Println("/!\\ Please enter a valid number")
			fmt.Printf("Error: %s\n", err)
			continue	// prevent the app to continue this iteration
		}

		// ---------------------------------------------------------------------
		// Check if the user choice is in the range of allowed values
		validButton := false
		for _, value := range buttons {
			if value == nextStep {
				validButton = true
				break
			}
		}
		if !validButton {
			fmt.Printf("/!\\ %d is not in the menu. \nPlease retry.\n", nextStep)
			continue
		}

		// ---------------------------------------------------------------------
		// Perform actions based on user choice
		switch nextStep {
		case 0:
			fmt.Println("Exiting the application..")
			appRunning = false
		case 1:
			actionAdd(d, reader)
		case 2:
			actionDefine(d, reader)
		case 3:
			actionRemove(d, reader)
		case 4:
			actionList(d)
			
		}

		// fmt.Printf("You entered: %s\n", inputText)
	}
	
	fmt.Println("Application closed.")
	*/
}

// Handler function for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Write a welcome message to the response writer
	w.Write([]byte("Welcome to the home page!"))
}
// Handler function for the about page
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



