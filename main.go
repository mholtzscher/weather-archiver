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

func wundergroundHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	log.Printf("Received Weather Underground data: %v", query)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "WU request received successfully")
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 not found")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /ecowitt", ecowittHandler)
	mux.HandleFunc("GET /wunderground", wundergroundHandler)
	mux.HandleFunc("/", defaultHandler)

	port := "8080"
	log.Printf("Starting server on port %s", port)
	log.Println("Listening for POST requests on http://localhost:8080/ecowitt")

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
