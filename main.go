package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const apiURL = "http://cab.inta-csic.es/rems/wp-content/plugins/marsweather-widget/api.php"

func main() {
	// GET data from INTA-CSIC API
	res, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Parse JSON response into Go Struct
	var csicResponse CSICResponse
	if err := json.NewDecoder(res.Body).Decode(&csicResponse); err != nil {
		log.Fatal(err)
	}

	// TODO Parse strings into apropriate data types
}
