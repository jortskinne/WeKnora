package metrics

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// Stats holds runtime metrics for the application.
type Stats struct {
	Uptime       string `json:"uptime"`
	GoRoutines   int    `json:"goroutines"`
	AllocMB      float64 `json:"alloc_mb"`
	TotalAllocMB float64 `json:"total_alloc_mb"`
	SysMB        float64 `json:"sys_mb"`
	NumGC        uint32  `json:"num_gc"`
}

var startTime = time.Now()

// Get returns the current runtime metrics.
func Get() Stats {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	return Stats{
		Uptime:       time.Since(startTime).Round(time.Second).String(),
		GoRoutines:   runtime.NumGoroutine(),
		AllocMB:      toMB(mem.Alloc),
		TotalAllocMB: toMB(mem.TotalAlloc),
		SysMB:        toMB(mem.Sys),
		NumGC:        mem.NumGC,
	}
}

func toMB(b uint64) float64 {
	return float64(b) / 1024 / 1024
}

// Handler returns an HTTP handler that exposes runtime metrics as JSON.
func Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"metrics": Get(),
	})
}
