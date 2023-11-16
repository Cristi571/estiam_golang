package dictionary

// Import dependencies
import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	Definition string
}

func (e Entry) String() string {
	return ""
}

type Dictionary struct {
	// The dictionary json filepath
	filePath string
	// The entries stands for data
	entries map[string]Entry
}
func New(filePath string) *Dictionary {
	d := &Dictionary{
		filePath: filePath,
		entries: make(map[string]Entry),
	}
	d.getDataFromFile()
	return d
}
func (d *Dictionary) getDataFromFile() error {
	file, err := os.Open(d.filePath)
	if err != nil {
		// In case the file does not exist yet
		if os.IsNotExist(err) {
			fmt.Println("File (.json) not found.")
			for {
				fmt.Println("Would you like to create the file now?")
				fmt.Println("+-------------------------------------+")
				fmt.Println("yes/y : Create the file")
				fmt.Println("no/n : Do not create now")
				fmt.Println("+-------------------------------------+")

				fmt.Print("Your choice: ")
				// Let the user to make a choice
				fmt.Print("Your choice: ")
				// Get the user input with a scanner that reads from os.Stdin
				reader := bufio.NewReader(os.Stdin)
				// Get the text that was read
				userInput, err := reader.ReadString('\n')
				if err != nil {
					fmt.Printf("Error reading input, please retry.\n%s", err)
					continue
				}
				userInput = strings.ToUpper(userInput)
				switch userInput {
				case "Y", "YES":
					fmt.Println("Creating the file..")
				}
				break
			}

			return nil
		}
		return err
	}

	decoder := json.NewDecoder(file)
	return decoder.Decode(&d.entries)
}

func (d *Dictionary) writeDataToFile() error {
	file, err := os.Create(d.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(d.entries)
}

func (d *Dictionary) Add(word string, definition string) {
	entry := Entry{Definition: definition}
	d.entries[word] = entry
	fmt.Printf("Word '%s' added to the dictionary.\n", word)
	d.writeDataToFile()
}

func (d *Dictionary) Get(word string) (Entry, error) {
	// Try to find the word in dictionary
	entry, found := d.entries[word]
	// Case word not found
	if !found {
		fmt.Printf("Word '%s' not found in the dictionary.", word)
		return Entry{}, nil
	}
	return entry, nil
}

func (d *Dictionary) Remove(word string) {
	// Delete the word from the dictionary
	delete(d.entries, word)
	// Update dictionary data
	d.writeDataToFile()
	fmt.Printf("Word '%s' removed from the dictionary.\n", word)
}

func (d *Dictionary) List() ([]string, map[string]Entry) {
	words := make([]string, 0, len(d.entries))
	for word := range d.entries {
		words = append(words, word)
	}
	return words, d.entries
}