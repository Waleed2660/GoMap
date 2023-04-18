package Helper

import (
	"fmt"
	"os"
	"sync"
)

var cache sync.Map

// GetApiKey method will cache the API key on first call
func GetApiKey() (string, error) {
	// Check if the data is already in cache
	val, ok := cache.Load("API_KEY")
	if ok {
		return val.(string), nil
	}
	// Fetch the data from the original source
	data, err := getKeyFromSystemEnv()
	if err != nil {
		return "Failed to get API_KEY from system", err
	}
	// Store the fetched data in cache
	cache.Store("API_KEY", data)
	return data, nil
}

func getKeyFromSystemEnv() (string, error) {
	return os.Getenv("API_KEY"), nil
}

func ApiKeyExists() bool {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API_KEY environment variable not set")
		return false
	}
	// Store the fetched data in cache
	cache.Store("API_KEY", apiKey)
	return true
}
