package handler

import (
	"sum_days/middlewares"
	"sum_days/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetToken(c *gin.Context) {
	token := middlewares.SetGetToken(c)
	if token!= nil {
		c.JSON(http.StatusOK, models.Message{
			Token: *token,
		})
	}
}