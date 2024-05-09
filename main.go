package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("PORT environment variable is required")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/{content...}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!" + port + "\n" + r.PathValue("content")))
	})
	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
