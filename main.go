package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/climate", handleClimateRequest)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
