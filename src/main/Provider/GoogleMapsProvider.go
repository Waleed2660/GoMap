package Provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Set up the API endpoint and parameters
	origin := "Manchester,UK"
	destination := "London,UK"
	apiKey := "" // Replace with your actual Google Maps API key
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/directions/json?origin=%s&destination=%s&key=%s", origin, destination, apiKey)

	// Send HTTP GET request to the API
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse the JSON response
	var directionsResponse struct {
		Routes []struct {
			Legs []struct {
				Duration struct {
					Text string `json:"text"`
				} `json:"duration"`
				Distance struct {
					Text string `json:"text"`
				} `json:"distance"`
			} `json:"legs"`
		} `json:"routes"`
		Status string `json:"status"`
	}
	err = json.Unmarshal(body, &directionsResponse)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		fmt.Println("Response -> ", response)
		return
	}

	// Extract and display the duration and distance of the shortest route
	if len(directionsResponse.Routes) > 0 && len(directionsResponse.Routes[0].Legs) > 0 {
		duration := directionsResponse.Routes[0].Legs[0].Duration.Text
		distance := directionsResponse.Routes[0].Legs[0].Distance.Text
		fmt.Println("Duration:", duration)
		fmt.Println("Distance:", distance)
	} else {
		fmt.Println("No route found")
	}
}
