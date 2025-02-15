// nolint: lll
package router

import (
	"github.com/gin-gonic/gin"

	"github.com/andys920605/meme-coin/internal/north/remote/source/handler"
	"github.com/andys920605/meme-coin/pkg/errors"
)

type Router struct {
	// middleware
	interceptorHandler gin.HandlerFunc

	// handler
	healthHandler   *handler.HealthHandler
	memeCoinHandler *handler.MemeCoinHandler
}

func NewRouter(
	interceptorHandler gin.HandlerFunc,
	healthHandler *handler.HealthHandler,
	memeCoinHandler *handler.MemeCoinHandler,
) *Router {
	return &Router{
		interceptorHandler: interceptorHandler,
		healthHandler:      healthHandler,
		memeCoinHandler:    memeCoinHandler,
	}
}

func (r *Router) Register(engine *gin.Engine) {
	engine.NoRoute(func(c *gin.Context) {
		_ = c.Error(errors.RouteNotFound)
	})
	engine.GET("/healthz", r.healthHandler.Check)

	// middleware
	normal := engine.Group("/srv", r.interceptorHandler)

	normal.POST("/meme-coins", r.memeCoinHandler.Create)
	normal.GET("/meme-coins/:id", r.memeCoinHandler.Get)
	normal.PUT("/meme-coins/:id", r.memeCoinHandler.Update)
	normal.DELETE("/meme-coins/:id", r.memeCoinHandler.Delete)
	normal.POST("/meme-coins/:id/poke", r.memeCoinHandler.Poke)
}
