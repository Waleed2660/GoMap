package main

import (
	"fmt"
	"github.com/Waleed2660/GoMap/src/main/Gateway"
	"github.com/Waleed2660/GoMap/src/main/Helper"
	"os"
)

func main() {

	// Terminates if no valid API Key is provided
	if !Helper.ApiKeyExists() {
		fmt.Println("Please set up a valid API key")
		os.Exit(1)
	}

	inboundGateway := Gateway.NewInboundGateway()
	inboundGateway.Run()

	fmt.Println("GoMap is running")
}
