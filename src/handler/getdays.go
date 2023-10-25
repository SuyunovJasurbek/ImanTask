package handler

import "github.com/gin-gonic/gin"

func (h *Handler) GetDays(c *gin.Context) {
	h.service.Time()
}