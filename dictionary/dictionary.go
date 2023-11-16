package dictionary
// Import dependencies
import (
	"fmt"
)

type Entry struct {
	Definition string
}

func (e Entry) String() string {

	return ""
}

type Dictionary struct {
	entries map[string]Entry
}

func New() *Dictionary {
	return &Dictionary{
		entries: make(map[string]Entry),
	}
}

func (d *Dictionary) Add(word string, definition string) {
	entry := Entry{Definition: definition}
	d.entries[word] = entry
	fmt.Printf("Word '%s' added to the dictionary.\n", word)
}

func (d *Dictionary) Get(word string) (Entry, error) {
	entry, found := d.entries[word]
	if !found {
		return Entry{}, fmt.Errorf("word '%s' not found in the dictionary", word)
	}
	return entry, nil
}

func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
	fmt.Printf("Word '%s' removed from the dictionary.\n", word)
}

func (d *Dictionary) List() ([]string, map[string]Entry) {
	words := make([]string, 0, len(d.entries))
	for word := range d.entries {
		words = append(words, word)
	}

	return words, d.entries
}