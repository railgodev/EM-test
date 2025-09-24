package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/railgodev/EM-test/backend/internal/model"
)

func (h *Handle) Update(c *gin.Context) {

	id := c.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid UUID format",
		})
		return
	}
	var req model.SubscriptionUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Warn("invalid request body", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}
	if req.EndDate != nil &&
		req.StartDate != nil &&
		(req.EndDate.Time.Before(req.StartDate.Time) || req.EndDate.Time.Equal(req.StartDate.Time)) {
		h.log.Warn("end_date is before start_date")
		c.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "end_date cannot be before or equal start_date",
		})
		return
	}

	resp, err := h.uc.Update(c.Request.Context(), id, &req)

	if err != nil {
		if err.Error() == "no fields to update" {
			h.log.Warn("no fields to update", slog.String("id", id))
			c.JSON(http.StatusBadRequest, model.Error{
				Code:    http.StatusBadRequest,
				Message: "no fields to update",
			})
			return
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.ConstraintName == "end_date_valid" {
				h.log.Warn("end_date is before start_date after update")
				c.JSON(http.StatusBadRequest, model.Error{
					Code:    http.StatusBadRequest,
					Message: "end_date cannot be before or equal start_date after update",
				})
				return
			}
		}
		h.log.Error("failed to create subscription", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "failed to create subscription",
		})
		return
	}
	if resp == nil {
		h.log.Warn("subscription not found", slog.String("id", id))
		c.JSON(http.StatusNotFound, model.Error{
			Code:    http.StatusNotFound,
			Message: "subscription not found",
		})
		return
	}

	h.log.Info("subscription created", slog.String("id", resp.ID.String()))
	c.JSON(http.StatusCreated, resp)
}
