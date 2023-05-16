package Google

import (
	"encoding/json"
	"fmt"
	"github.com/Waleed2660/GoMap/src/main/Helper"
	"io/ioutil"
	"net/http"
)

type PlaceSearchResponse struct {
	Results []struct {
		PlaceID string `json:"place_id"`
	} `json:"results"`
}

func GetPlaceID(address string) (string, error) {

	// Replace with your actual Google Maps API key
	apiKey, err := Helper.GetApiKey()

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=%s&inputtype=textquery&key=%s", address, apiKey)

	//url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=martin,%s&inputtype=textquery&fields=name,place_id,geometry,formatted_address&key=%s", address, apiKey)
	//url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/textsearch/json?query=%s&key=%s", address, apiKey)

	// call API
	resp, err := http.Get(url)

	if err != nil {
		return resp.Status, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading PlaceID response body:", err)
	}
	fmt.Println("Place ID - Response: ", string(bodyBytes))

	var placeSearchResponse PlaceSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&placeSearchResponse); err != nil {
		return "", err
	}

	if len(placeSearchResponse.Results) == 0 {
		return "", fmt.Errorf("no results found")
	}

	return placeSearchResponse.Results[0].PlaceID, nil
}
