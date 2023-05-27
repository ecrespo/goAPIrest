package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Error: Method not allowed")
		return
	}
	fmt.Fprintf(w, "Hello World")

}

func getCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(countries)
	if err != nil {
		log.Println(err)
		return
	}
}

func addCountry(w http.ResponseWriter, r *http.Request) {
	country := &Country{}
	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err)
		log.Println(err)
		return
	}
	countries = append(countries, country)
	log.Println("Country added: ", country.Name)
	_, err = fmt.Fprintf(w, "Country added: %s", country.Name)
	if err != nil {
		log.Println(err)
		return
	}

}
