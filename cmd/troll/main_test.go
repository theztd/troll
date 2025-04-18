package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.com/theztd/troll/internal/server"
)

func TestHealthEndpoint(t *testing.T) {
	router := server.InitRoutes()
	req, _ := http.NewRequest("GET", "/_healthz/ready.json", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
}
