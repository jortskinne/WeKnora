package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/health", Handler())
	return r
}

func TestHealthHandler_ReturnsOK(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var s Status
	if err := json.NewDecoder(w.Body).Decode(&s); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if s.Status != "ok" {
		t.Errorf("expected status 'ok', got '%s'", s.Status)
	}

	if s.GoVersion == "" {
		t.Error("expected non-empty GoVersion")
	}

	if s.Uptime == "" {
		t.Error("expected non-empty Uptime")
	}
}

func TestHealthHandler_ContainsMemoryCheck(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	r.ServeHTTP(w, req)

	var s Status
	if err := json.NewDecoder(w.Body).Decode(&s); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if _, ok := s.Checks["memory"]; !ok {
		t.Error("expected 'memory' key in checks")
	}
}

func TestCheckMemory_ReturnsOK(t *testing.T) {
	result := checkMemory()
	if result != "ok" {
		t.Errorf("expected 'ok' from checkMemory, got '%s'", result)
	}
}
