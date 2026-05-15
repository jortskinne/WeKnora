package info

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/info", Handler())
	return r
}

func TestInfoHandler_ReturnsOK(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/info", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestInfoHandler_ContainsExpectedFields(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/info", nil)
	r.ServeHTTP(w, req)

	var body map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	if body["status"] != "ok" {
		t.Errorf("expected status 'ok', got %v", body["status"])
	}

	if _, ok := body["info"]; !ok {
		t.Error("expected 'info' field in response")
	}
}

func TestInfoGet_GoVersionPrefix(t *testing.T) {
	info := Get()
	if !strings.HasPrefix(info.GoVersion, "go") {
		t.Errorf("expected GoVersion to start with 'go', got %q", info.GoVersion)
	}
}

func TestInfoGet_OSAndArch(t *testing.T) {
	info := Get()
	if info.OS != runtime.GOOS {
		t.Errorf("expected OS %q, got %q", runtime.GOOS, info.OS)
	}
	if info.Arch != runtime.GOARCH {
		t.Errorf("expected Arch %q, got %q", runtime.GOARCH, info.Arch)
	}
}

func TestInfoGet_CPUsPositive(t *testing.T) {
	info := Get()
	if info.CPUs <= 0 {
		t.Errorf("expected CPUs > 0, got %d", info.CPUs)
	}
}

func TestInfoGet_UptimeNonEmpty(t *testing.T) {
	info := Get()
	if info.Uptime == "" {
		t.Error("expected non-empty Uptime")
	}
	if info.StartedAt == "" {
		t.Error("expected non-empty StartedAt")
	}
}
