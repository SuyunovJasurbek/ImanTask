package handler

import (
	"days_limit/middlewares"
	"days_limit/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SetToken(c *gin.Context) {
	token := middlewares.SetGetToken(c)
	if token != "" {
		c.JSON(404, models.Error{
			Error: "Token's get error",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		Token: token,
	})
	return
}