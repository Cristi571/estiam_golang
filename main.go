package main

// Import dependencies
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"dictionary/dictionary"
)

func main() {
	// Create a new instance of the dictionary from the package
	d := dictionary.New()

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
		nextStep, err := strconv.Atoi(inputText[:len(inputText)-2]) // Trim newline character
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
}




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
	word = word[:len(word)-2] // Trim newline character

	// Let the user give the definition of the item
	fmt.Print("Enter the definition: ")
	definition, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	definition = definition[:len(definition)-2] // Trim newline character

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
	word = word[:len(word)-2] // Trim newline character

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
	word = word[:len(word)-2] // Trim newline character

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