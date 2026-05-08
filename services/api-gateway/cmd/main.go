package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	hostname, _ := os.Hostname()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s %s %s %s", hostname, r.Proto, r.Method, r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	type ScoreResponse struct {
		Score string `json:"score"`
	}

	mux.HandleFunc("/score", func(w http.ResponseWriter, r *http.Request) {
		score := os.Getenv("DEFAULT_SCORE")
		if score == "" {
			score = "770"
		}
		response := ScoreResponse{
			Score: score,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	mux.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s %s %s", r.Proto, r.Method, r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("readOK"))
	})

	log.Printf("Server started on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
