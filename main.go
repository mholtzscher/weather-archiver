package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func ecowittHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	log.Printf("Received Ecowitt data: %s", string(body))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Request received successfully")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /ecowitt", ecowittHandler)

	port := "8080"
	log.Printf("Starting server on port %s", port)
	log.Println("Listening for POST requests on http://localhost:8080/ecowitt")

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
