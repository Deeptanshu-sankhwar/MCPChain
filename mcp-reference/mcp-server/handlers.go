package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type EchoRequest struct {
	Text string `json:"text"`
}

type EchoResponse struct {
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

type TimeResponse struct {
	Time string `json:"time"`
}

func handleEcho(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req EchoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp := EchoResponse{
		Text:      req.Text,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	writeJSON(w, resp)
}

func handleTime(w http.ResponseWriter, r *http.Request) {
	resp := TimeResponse{
		Time: time.Now().UTC().Format(time.RFC3339),
	}
	writeJSON(w, resp)
}

func handleSchema(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "schema.json")
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
