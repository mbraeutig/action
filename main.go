package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mbraeutig/afh/api"
)

func main() {
	http.HandleFunc("/action.jpg", api.Action)
	http.HandleFunc("/health", api.Health)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
