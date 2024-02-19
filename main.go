package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/mbp/endpoint", handlePostRequest)
	http.HandleFunc("/mbp/healthcheck", healthcheck)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body.", http.StatusBadRequest)
		return
	}

	fmt.Println("Received POST data:", string(body))

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "POST request processed successfully.")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("health check req recieved")
	w.WriteHeader(http.StatusOK)

}
