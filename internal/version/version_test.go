package version_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/Tencent/WeKnora/internal/version"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/version", version.Handler())
	return r
}

func TestVersionHandler_ReturnsOK(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/version", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestVersionHandler_ContainsExpectedFields(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/version", nil)
	r.ServeHTTP(w, req)

	var resp map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp["status"] != "ok" {
		t.Errorf("expected status 'ok', got %v", resp["status"])
	}

	data, ok := resp["data"].(map[string]interface{})
	if !ok {
		t.Fatal("expected 'data' field to be an object")
	}

	for _, field := range []string{"version", "commit", "build_date", "go_version", "os", "arch"} {
		if _, exists := data[field]; !exists {
			t.Errorf("expected field %q in version data", field)
		}
	}
}

func TestVersionGet_GoVersionPrefix(t *testing.T) {
	info := version.Get()
	if !strings.HasPrefix(info.GoVersion, "go") {
		t.Errorf("expected GoVersion to start with 'go', got %q", info.GoVersion)
	}
}

func TestVersionGet_DefaultValues(t *testing.T) {
	info := version.Get()
	if info.Version == "" {
		t.Error("expected Version to be non-empty")
	}
	if info.OS == "" {
		t.Error("expected OS to be non-empty")
	}
	if info.Arch == "" {
		t.Error("expected Arch to be non-empty")
	}
}
