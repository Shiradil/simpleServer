package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const port = ":8080"

func main() {
	http.HandleFunc("/", handlePostRequest)
	fmt.Printf("Server listening on port %s...\n", port)
	fmt.Printf("http://localhost:8080/\n")
	http.ListenAndServe(port, nil)
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if requestBody.Message == "" {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received message: %s\n", requestBody.Message)

	response := Response{
		Status:  "success",
		Message: "Data successfully received",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
