package server

import (
	"fmt"
	"log"
	"net/http"
)

func initRoutes() {

	http.HandleFunc("/", index)
	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCountry(w, r)
		case http.MethodPost:
			addCountry(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			log.Println("Error: Method not allowed")
			n, err := fmt.Fprintf(w, "Error: Method not allowed")
			if err != nil {
				log.Println(err)
				fmt.Println(n, err)
			}
			return
		}
	})

}
