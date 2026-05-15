package metrics

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/metrics", Handler)
	return r
}

func TestMetricsHandler_ReturnsOK(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/metrics", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMetricsHandler_ContainsExpectedFields(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/metrics", nil)
	r.ServeHTTP(w, req)

	var body map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &body)
	require.NoError(t, err)

	assert.Equal(t, "ok", body["status"])

	metrics, ok := body["metrics"].(map[string]interface{})
	require.True(t, ok, "metrics field should be an object")

	assert.Contains(t, metrics, "uptime")
	assert.Contains(t, metrics, "goroutines")
	assert.Contains(t, metrics, "alloc_mb")
	assert.Contains(t, metrics, "total_alloc_mb")
	assert.Contains(t, metrics, "sys_mb")
	assert.Contains(t, metrics, "num_gc")
}

func TestMetricsGet_GoRoutinesPositive(t *testing.T) {
	stats := Get()
	assert.Greater(t, stats.GoRoutines, 0)
}

func TestMetricsGet_UptimeNonEmpty(t *testing.T) {
	stats := Get()
	assert.NotEmpty(t, stats.Uptime)
}

func TestMetricsGet_AllocPositive(t *testing.T) {
	stats := Get()
	assert.Greater(t, stats.AllocMB, float64(0))
}
