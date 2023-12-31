package interceptors

import (
	"gobatchbuilder/internal/interceptors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggerInterceptor(t *testing.T) {
	// Create a test HTTP handler that uses the LoggerInterceptor
	handler := interceptors.LoggerInterceptor(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your test logic here
		w.WriteHeader(http.StatusOK)
	}))

	// Create a test HTTP request
	req := httptest.NewRequest("GET", "/test", nil)

	// Create a test HTTP response recorder
	w := httptest.NewRecorder()

	// Serve the request using the handler
	handler.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}
