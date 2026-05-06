package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s %s %s", r.Proto, r.Method, r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("/score", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s %s %s", r.Proto, r.Method, r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"score": 700}`))
	})

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
