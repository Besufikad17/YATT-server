package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"message": "hi :)",
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Write(data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ping)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
