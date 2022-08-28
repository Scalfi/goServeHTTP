package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleGET)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// func (s *serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch {
// 	case r.Method == "GET":
// 		HandleGET(w, r)
// 	default:
// 		fmt.Fprintf(w, "oi")
// 	}
// }

func HandleGET(w http.ResponseWriter, r *http.Request) {
	jsonResponse := struct {
		Name string `json:"name"`
	}{"Gabriel"}

	json.NewEncoder(w).Encode(jsonResponse)
}
