package main

import (
	"fmt"
	"main/src/main/Helper"
	"os"
)

func main() {

	// Terminates if no valid API Key is provided
	if !Helper.ApiKeyExists() {
		fmt.Println("Please set up a valid API key")
		os.Exit(1)
	}

}
