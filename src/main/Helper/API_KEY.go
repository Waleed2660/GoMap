package Helper

import (
	"fmt"
	"os"
)

func getApiKey() string {
	return os.Getenv("API_KEY")
}

func ApiKeyExists() bool {
	apiKey := os.Getenv("API_KEY")
	fmt.Println("ApiKeyExists ---> ", apiKey)
	if apiKey == "" {
		fmt.Println("API_KEY environment variable not set")
		return false
	}
	return true
}
