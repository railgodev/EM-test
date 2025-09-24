package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/railgodev/EM-test/backend/internal/model"
)

func (h *Handle) Total(c *gin.Context) {
	// gin has troubles with binding custom types, so we bind to string and parse manually
	// https://github.com/gin-gonic/gin/issues/2423
	var req model.SubscriptionsTotal
	var err error

	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Warn("invalid request parameters", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		})
		return
	}
	req.UserID, err = uuid.Parse(req.UserIDString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid UUID format",
		})
		return
	}
	if err := req.Start.UnmarshalText([]byte(req.StartString)); err != nil {
		h.log.Warn("invalid start date", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid start date",
		})
		return
	}
	if err := req.End.UnmarshalText([]byte(req.EndString)); err != nil {
		h.log.Warn("invalid end date", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid end date",
		})
		return
	}
	h.log.Debug("Total request", slog.Any("req", []any{
		"ServiceName", req.ServiceName,
		"UserID", req.UserID,
		"Start", req.Start.Time,
		"End", req.End.Time,
	}))
	if req.End.Time.Equal(req.Start.Time) || req.End.Time.Before(req.Start.Time) {
		h.log.Warn("end_date is before start_date")
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "end_date cannot be before or equal start_date",
		})
		return
	}
	total, err := h.uc.Total(c.Request.Context(), req)
	if err != nil {
		h.log.Error("failed to calculate total subscription cost", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "failed to calculate total subscription cost",
		})
		return
	}

	h.log.Info("total subscription cost calculated", slog.Int("total", total))
	c.JSON(http.StatusOK, gin.H{"total": total})
}
