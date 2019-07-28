package router

import (
	"net/http"

	"./middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(mw...)
	// 404 handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The sp handlers
	sp := g.Group("/sp")

	return g
}
