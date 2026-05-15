package info

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

// Info holds combined system and application information.
type Info struct {
	GoVersion  string `json:"go_version"`
	OS         string `json:"os"`
	Arch       string `json:"arch"`
	CPUs       int    `json:"cpus"`
	GoRoutines int    `json:"goroutines"`
	Uptime     string `json:"uptime"`
	StartedAt  string `json:"started_at"`
}

// Get returns the current system and runtime information.
func Get() Info {
	return Info{
		GoVersion:  runtime.Version(),
		OS:         runtime.GOOS,
		Arch:       runtime.GOARCH,
		CPUs:       runtime.NumCPU(),
		GoRoutines: runtime.NumGoroutine(),
		Uptime:     time.Since(startTime).Round(time.Second).String(),
		StartedAt:  startTime.UTC().Format(time.RFC3339),
	}
}

// Handler returns a Gin handler that responds with system info as JSON.
func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"info":   Get(),
		})
	}
}
