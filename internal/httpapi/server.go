package httpapi

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

const (
	serviceName = "devops-lab-micro"
	version     = "1.0.0"
)

type jsonResponse map[string]any

func NewServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/version", versionHandler)
	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}

	respondJSON(w, http.StatusOK, jsonResponse{
		"service": serviceName,
		"message": "hello from go microservice",
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}

	respondJSON(w, http.StatusOK, jsonResponse{
		"status":  "ok",
		"service": serviceName,
		"time":    time.Now().UTC().Format(time.RFC3339),
	})
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}

	respondJSON(w, http.StatusOK, jsonResponse{
		"service":    serviceName,
		"version":    version,
		"go_version": runtime.Version(),
	})
}

func methodNotAllowed(w http.ResponseWriter) {
	respondJSON(w, http.StatusMethodNotAllowed, jsonResponse{
		"error": "method not allowed",
	})
}

func respondJSON(w http.ResponseWriter, status int, payload jsonResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
