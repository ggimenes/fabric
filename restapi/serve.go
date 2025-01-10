package restapi

import (
	"github.com/danielmiessler/fabric/core"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func Serve(registry *core.PluginRegistry, address string) (err error) {
	r := gin.New()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Register API routes
	fabricDb := registry.Db
	NewPatternsHandler(r, fabricDb.Patterns)
	NewContextsHandler(r, fabricDb.Contexts)
	NewSessionsHandler(r, fabricDb.Sessions)
	NewChatHandler(r, registry, fabricDb)
	NewConfigHandler(r, fabricDb)
	NewModelsHandler(r, registry.VendorManager)

	// Serve static files from the web UI build directory
	webDir := "./dist/web"
	if os.Getenv("GO_ENV") == "production" {
		// In production, serve from the absolute path
		execPath, _ := os.Executable()
		webDir = filepath.Join(filepath.Dir(execPath), "dist/web")
	}

	// Serve static files
	r.Static("/", webDir)

	// SPA fallback - serve index.html for any unmatched routes
	r.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(webDir, "index.html"))
	})

	// Start server
	err = r.Run(address)
	if err != nil {
		return err
	}

	return
}
