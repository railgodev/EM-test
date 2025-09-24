package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/railgodev/EM-test/backend/internal/model"
)

func (h *Handle) Create(c *gin.Context) {
	var req model.SubscriptionCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Warn("invalid request body", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}
	if req.EndDate != nil && (req.EndDate.Time.Before(req.StartDate.Time) || req.EndDate.Time.Equal(req.StartDate.Time)) {
		h.log.Warn("end_date is before start_date")
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "end_date cannot be before or equal start_date",
		})
		return
	}
	resp, err := h.uc.Create(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("failed to create subscription", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "failed to create subscription",
		})
		return
	}

	h.log.Info("subscription created", slog.String("id", resp.ID.String()))
	c.JSON(http.StatusCreated, resp)
}
