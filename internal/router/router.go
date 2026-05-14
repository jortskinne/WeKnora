package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/WeKnora/WeKnora/internal/health"
)

// Config holds configuration for the router.
type Config struct {
	// Mode is the Gin mode: "debug", "release", or "test".
	Mode string
}

// New creates and configures a new Gin engine with all application routes.
func New(cfg Config) *gin.Engine {
	if cfg.Mode != "" {
		gin.SetMode(cfg.Mode)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	registerHealthRoutes(r)

	// Fallback for undefined routes.
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
		})
	})

	return r
}

// registerHealthRoutes attaches health-check endpoints to the router.
func registerHealthRoutes(r *gin.Engine) {
	r.GET("/health", health.Handler())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
