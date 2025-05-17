package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", handleEcho)
	http.HandleFunc("/time", handleTime)
	http.HandleFunc("/schema", handleSchema)

	log.Println("🔧 MCP Reference Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
