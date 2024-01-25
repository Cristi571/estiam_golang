package dictionary

// Import dependencies
import (
	"os"
	"fmt"
	"encoding/json"
	"context"
	"github.com/go-redis/redis/v8"
)

type Entry struct {
	Definition string
}
func (e Entry) String() string {
	return ""
}

// Context var is used for redis pack
var ctx = context.Background()
// Create the Redis client
var redisClient = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379", // Replace with your Redis server address
	Password: "",               // No password for local Redis instance
	DB:       0,                // Default DB
})

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
		// Try to create a new json file if it doesn't exist
		if os.IsNotExist(err) {
			fmt.Println("File (.json) not found. Creating new file..")
			newFile, err := os.Create(d.filePath)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return nil
			}
			defer newFile.Close() // Close the file on exit
			data := map[string]interface{}{} // Define initial data
			// Encode the data and write it to the new file
			encoder := json.NewEncoder(newFile)
			if err := encoder.Encode(data); err != nil {
				fmt.Println("Error encoding JSON:", err)
				return nil
			}
			fmt.Println("JSON file created successfully.")
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

func (d *Dictionary) Add(word string, definition string)(string, error) {
	entry := Entry{Definition: definition}
	d.entries[word] = entry
	d.writeDataToFile()
	// fmt.Printf("Word '%s' added to the dictionary.\n", word)
	return "Word "+word+" added to the dictionary.", nil
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

func (d *Dictionary) Remove(word string)(string, error) {
	// Delete the word from the dictionary
	delete(d.entries, word)
	// Update dictionary data
	d.writeDataToFile()
	// fmt.Printf("Word '%s' removed from the dictionary.\n", word)
	return "Word "+word+" deleted from the dictionary.", nil
}

func (d *Dictionary) List() ([]string, map[string]Entry) {
	words := make([]string, 0, len(d.entries))
	for word := range d.entries {
		words = append(words, word)
	}
	return words, d.entries
}



// Save data to Redis
func (d *Dictionary) SaveToRedis(word string, definition string)(string, error) {
	fmt.Println("Writing data to Redis..")
	entry := Entry{Definition: definition}
	d.entries[word] = entry
	// Convert the entry to JSON
	data, err := json.Marshal(entry)
	if err != nil {
		return "Failed to add "+word+" to Redis.", err
	}

	// Set the key-value pair in Redis
	err = redisClient.Set(ctx, word, data, 0).Err()
	if err != nil {
		fmt.Println("Writing data to Redis FAILED.")
		return "Failed to add "+word+" to Redis.", err
	}

	fmt.Println("Writing data to Redis SUCCEEDED.")
	return "Word "+word+" added to Redis.", nil
}

// Retrieve data from Redis
func (d *Dictionary) GetFromRedis(word string) (Entry, error) {
	// Get the value from Redis

	data, err := redisClient.Get(ctx, word).Result()
	if err != nil {
		return Entry{}, err
	}

	// Unmarshal the JSON data into the Entry struct
	var entry Entry
	err = json.Unmarshal([]byte(data), &entry)
	if err != nil {
		return Entry{}, err
	}

	return entry, nil
}
