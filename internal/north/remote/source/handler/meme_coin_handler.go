package handler

import (
	"net/http"
	"strconv"

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
	var req request.CreateMemeCoin
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(errors.InvalidRequest.Wrap(err, "invalid request"))
		return
	}
	if err := req.Valid(); err != nil {
		_ = c.Error(errors.InvalidRequest.Wrap(err, "invalid request"))
		return
	}
	cmd := message.CreateMemeCoinCommand{
		Name:        req.Name,
		Description: req.Description,
	}
	rsp, err := h.memeCoinAppService.CreateMemeCoin(c.Request.Context(), cmd)
	if err != nil {
		_ = c.Error(errors.InternalServerError.Wrap(err, "internal server error"))
		return
	}

	template_response.OK(rsp).To(c, http.StatusOK)
}

func (h *MemeCoinHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if err := validID(id); err != nil {
		_ = c.Error(errors.InvalidRequest.Wrap(err, "invalid request"))
		return
	}

	query := message.GetMemeCoinQuery{
		ID: id,
	}

	rsp, err := h.memeCoinAppService.GetMemeCoin(c.Request.Context(), query)
	if err != nil {
		_ = c.Error(errors.InternalServerError.Wrap(err, "internal server error"))
		return
	}

	template_response.OK(rsp).To(c, http.StatusOK)
}

func (h *MemeCoinHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if err := validID(id); err != nil {
		_ = c.Error(errors.InvalidRequest.Wrap(err, "invalid request"))
		return
	}

	var req request.UpdateMemeCoin
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(errors.InvalidRequest.Wrap(err, "invalid request"))
		return
	}
	if err := req.Valid(); err != nil {
		_ = c.Error(errors.InvalidRequest.Wrap(err, "invalid request"))
		return
	}

	cmd := message.UpdateMemeCoinCommand{
		ID:          id,
		Description: req.Description,
	}
	err := h.memeCoinAppService.UpdateMemeCoin(c.Request.Context(), cmd)
	if err != nil {
		_ = c.Error(errors.InternalServerError.Wrap(err, "internal server error"))
		return
	}

	template_response.OK(nil).To(c, http.StatusOK)
}

func (h *MemeCoinHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := validID(id); err != nil {
		_ = c.Error(errors.InvalidRequest.Wrap(err, "invalid request"))
		return
	}
	cmd := message.DeleteMemeCoinCommand{
		ID: id,
	}
	err := h.memeCoinAppService.DeleteMemeCoin(c.Request.Context(), cmd)
	if err != nil {
		_ = c.Error(errors.InternalServerError.Wrap(err, "internal server error"))
		return
	}

	template_response.OK(nil).To(c, http.StatusOK)
}

func (h *MemeCoinHandler) Poke(c *gin.Context) {
	id := c.Param("id")
	if err := validID(id); err != nil {
		_ = c.Error(errors.InvalidRequest.Wrap(err, "invalid request"))
		return
	}
	cmd := message.PokeMemeCoinCommand{
		ID: id,
	}
	err := h.memeCoinAppService.PokeMemeCoin(c.Request.Context(), cmd)
	if err != nil {
		_ = c.Error(errors.InternalServerError.Wrap(err, "internal server error"))
		return
	}

	template_response.OK(nil).To(c, http.StatusOK)
}

func validID(id string) error {
	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		return errors.New("invalid id")
	}
	return nil
}
