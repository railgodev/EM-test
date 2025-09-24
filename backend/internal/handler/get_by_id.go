package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/railgodev/EM-test/backend/internal/model"
)

func (h *Handle) GetByID(c *gin.Context) {

	id := c.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid UUID format",
		})
		return
	}
	subscription, err := h.uc.GetByID(c.Request.Context(), id)
	if err != nil {
		h.log.Error("failed to get subscription", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "failed to get subscription",
		})
		return
	}

	if subscription == nil {
		h.log.Warn("subscription not found", slog.String("id", id))
		c.JSON(http.StatusNotFound, model.Error{
			Code:    http.StatusNotFound,
			Message: "subscription not found",
		})
		return
	}

	h.log.Info("subscription found", slog.String("id", subscription.ID.String()))
	c.JSON(http.StatusOK, subscription)
}
