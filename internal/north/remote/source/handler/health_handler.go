package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return new(HealthHandler)
}

func (h *HealthHandler) Check(c *gin.Context) {
	c.Status(http.StatusOK)
}
