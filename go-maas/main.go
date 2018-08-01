package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const apiURL = "http://cab.inta-csic.es/rems/wp-content/plugins/marsweather-widget/api.php"

func testRequest() {
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

	// Parse soles
	for index, sol := range csicResponse.Soles {
		fmt.Println(index, sol2fsol(sol))
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleIndex)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%q", html.EscapeString(r.URL.Path))
}
