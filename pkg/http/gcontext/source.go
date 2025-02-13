package gcontext

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Source string

const (
	SourceApp     Source = "app"
	SourceWeb     Source = "web"
	SourceService Source = "srv"
)

func ParseSource(c *gin.Context) Source {
	switch {
	case strings.Contains(c.Request.URL.Path, "/app/"):
		return SourceApp
	case strings.Contains(c.Request.URL.Path, "/web/"):
		return SourceWeb
	case strings.Contains(c.Request.URL.Path, "/srv/"):
		return SourceService
	default:
		return SourceService
	}
}

func GetSource(c *gin.Context) Source {
	source, _ := c.Get(ContextKeySource)
	return source.(Source)
}

func SetSource(c *gin.Context, source Source) {
	c.Set(ContextKeySource, source)
}
