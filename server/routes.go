package server

import (
	"fmt"
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
			n, err := fmt.Fprintf(w, "Error: Method not allowed")
			if err != nil {
				fmt.Println(n, err)
			}
			return
		}
	})

}
