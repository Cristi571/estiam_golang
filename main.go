package main

// Import dependencies
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Create a map to store dictionary content
	dictionary := make(map[string]string)

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
			actionAdd(dictionary, reader)
		case 2:
			actionDefine(dictionary, reader)
		case 3:
			actionRemove(dictionary, reader)
		case 4:
			actionList(dictionary)
			
		}

		// fmt.Printf("You entered: %s\n", inputText)
	}
	
	fmt.Println("Application closed.")
}




// -----------------------------------------------------------------------------
// Add new items to the dictionary (the key)
func actionAdd(d map[string]string, reader *bufio.Reader) {
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

	// Add the new item with it's definition to the dictionary
	d[word] = definition
	fmt.Printf("Word '%s' added to the dictionary.\n", word)
}

// -----------------------------------------------------------------------------
// Get the definition of a given word
func actionDefine(d map[string]string, reader *bufio.Reader) {
	fmt.Print("Enter a word to define: ")
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	word = word[:len(word)-2] // Trim newline character

	definition, found := d[word]
	if found {
		fmt.Printf("%s: %s\n", word, definition)
	} else {
		fmt.Printf("Word '%s' not found in the dictionary.\n", word)
	}
}

// -----------------------------------------------------------------------------
// Remove an item from the dictionary
func actionRemove(d map[string]string, reader *bufio.Reader) {
	fmt.Print("Enter a word to remove: ")
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	word = word[:len(word)-2] // Trim newline character

	delete(d, word)
	fmt.Printf("Word '%s' removed from the dictionary.\n", word)
}


// -----------------------------------------------------------------------------
// Display the dictionary content
func actionList(d map[string]string) {
	fmt.Println("Dictionary Content:")
	for word, definition := range d {
		fmt.Printf("%s: %s\n", word, definition)
	}
}