package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/andys920605/meme-coin/internal/north/local/appservice"
	"github.com/andys920605/meme-coin/internal/north/message"
	"github.com/andys920605/meme-coin/internal/north/remote/source/handler/request"
	"github.com/andys920605/meme-coin/pkg/errors"
	"github.com/andys920605/meme-coin/pkg/http/template_response"
)

type MemeCoinHandler struct {
	memeCoinAppService *appservice.MemeCoinAppService
}

func NewMemeCoinHandler(memeCoinAppService *appservice.MemeCoinAppService) *MemeCoinHandler {
	return &MemeCoinHandler{
		memeCoinAppService: memeCoinAppService,
	}
}

func (h *MemeCoinHandler) Create(c *gin.Context) {
	var req request.CreateMemeCoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(errors.InvalidRequest.Wrap(err, "invalid request"))
		return
	}
	if err := req.Valid(); err != nil {
		_ = c.Error(errors.InvalidRequest.Wrap(err, "invalid request"))
		return
	}
	cmd := message.CreateMemeCoinCommand{}
	err := h.memeCoinAppService.Create(c.Request.Context(), cmd)
	if err != nil {
		_ = c.Error(errors.InternalServerError.Wrap(err, "internal server error"))
		return
	}

	template_response.OK(nil).To(c, http.StatusOK)
}
