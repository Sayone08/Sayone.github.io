package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type contactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type apiResponse struct {
	Message string `json:"message"`
}

type profileResponse struct {
	Name   string   `json:"name"`
	Role   string   `json:"role"`
	Skills []string `json:"skills"`
	Status string   `json:"status"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", healthHandler)
	mux.HandleFunc("/api/profile", profileHandler)
	mux.HandleFunc("/api/contact", contactHandler)
	mux.Handle("/", http.FileServer(http.Dir("public")))

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           logging(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("portfolio server listening on http://localhost:%s", port)
	log.Fatal(server.ListenAndServe())
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	writeJSON(w, http.StatusOK, apiResponse{Message: "ok"})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	writeJSON(w, http.StatusOK, profileResponse{
		Name:   "Sayone Dey",
		Role:   "Lecturer, researcher and AI educator",
		Skills: []string{"Go", "Python", "Machine Learning", "Data Science"},
		Status: "Available for thoughtful collaborations",
	})
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var request contactRequest
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	if err := decoder.Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, apiResponse{Message: "Please send a valid message."})
		return
	}

	request.Name = strings.TrimSpace(request.Name)
	request.Email = strings.TrimSpace(request.Email)
	request.Message = strings.TrimSpace(request.Message)
	if request.Name == "" || request.Message == "" || !strings.Contains(request.Email, "@") {
		writeJSON(w, http.StatusBadRequest, apiResponse{Message: "Please complete your name, email, and message."})
		return
	}

	log.Printf("contact request from %s <%s>: %s", request.Name, request.Email, request.Message)
	writeJSON(w, http.StatusAccepted, apiResponse{Message: "Thanks! Your message reached the Go service."})
}

func writeJSON(w http.ResponseWriter, status int, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s (%s)", r.Method, filepath.Clean(r.URL.Path), time.Since(start))
	})
}
