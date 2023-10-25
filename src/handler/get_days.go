package handler

import (
	"days_limit/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDays(c *gin.Context) {

	days := h.service.Time()

	// error algoretm sum of days
	if days == 0 {
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "days algoretm error",
		})
		return
	}

	// sum of days
	c.JSON(http.StatusOK, models.DaysMessage{
		Days: days,
	})
}