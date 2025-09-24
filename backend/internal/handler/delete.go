package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/railgodev/EM-test/backend/internal/model"
)

func (h *Handle) Delete(c *gin.Context) {
	id := c.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid UUID format",
		})
		return
	}

	subscription, err := h.uc.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusInternalServerError, model.Error{
				Code:    http.StatusInternalServerError,
				Message: "failed to delete subscription",
			})
		}
		h.log.Error("failed to delete subscription", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "failed to delete subscription",
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
	h.log.Info("subscription deleted", slog.String("id", id))
	c.JSON(http.StatusAccepted, subscription)
}
