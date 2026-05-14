package health

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// Status represents the health status of the application.
type Status struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Version   string            `json:"version"`
	GoVersion string            `json:"go_version"`
	Uptime    string            `json:"uptime"`
	Checks    map[string]string `json:"checks"`
}

var startTime = time.Now()

// AppVersion is set at build time via ldflags.
var AppVersion = "dev"

// Handler returns a Gin handler function for the health check endpoint.
func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		checks := map[string]string{
			"memory": checkMemory(),
		}

		s := Status{
			Status:    "ok",
			Timestamp: time.Now().UTC(),
			Version:   AppVersion,
			GoVersion: runtime.Version(),
			Uptime:    time.Since(startTime).Round(time.Second).String(),
			Checks:    checks,
		}

		for _, v := range checks {
			if v != "ok" {
				s.Status = "degraded"
				break
			}
		}

		httpStatus := http.StatusOK
		if s.Status != "ok" {
			httpStatus = http.StatusServiceUnavailable
		}

		c.JSON(httpStatus, s)
	}
}

// checkMemory returns "ok" if allocated memory is below 500 MB.
func checkMemory() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	const maxAllocMB = 500
	if m.Alloc/1024/1024 > maxAllocMB {
		return "high memory usage"
	}
	return "ok"
}
