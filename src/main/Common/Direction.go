package Common

// Direction for Google DirectionsAPI
type Direction struct {
	CurrentAddress string `json:"current_address"`
	Destination    string `json:"destination"`
}
