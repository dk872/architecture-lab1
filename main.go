package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func timeManager(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	currTime := map[string]string{"time": time.Now().Format(time.RFC3339)}
	responseWriter.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(responseWriter).Encode(currTime); err != nil {
		http.Error(responseWriter, "Failed to encode response", http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
	}
}

func main() {
	http.HandleFunc("/time", timeManager)
	PORT := ":8795"
	log.Println("Server is running on http://localhost" + PORT)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
