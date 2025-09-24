package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/railgodev/EM-test/backend/internal/model"
)

func (h *Handle) List(c *gin.Context) {
	var req model.SubscriptionsListRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Warn("invalid request parameters", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		})
		return
	}

	subs, err := h.uc.List(c.Request.Context(), req.Page, req.PageSize)
	if err != nil {
		h.log.Error("failed to list subscriptions", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "failed to list subscriptions",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":      req.Page,
		"page_size": req.PageSize,
		"data":      subs,
	})
}
