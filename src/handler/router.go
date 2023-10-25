package handler

import (
	"days_limit/middlewares"
	"days_limit/src/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	w := gin.Default()
	s := service.NewService()
	h := NewHandler(s)
	
	// get token
	w.POST("/gettoken", h.GetToken)

	//middlewares
	w.Use(middlewares.CheckToken)

	//get days
	w.POST("/getdays", h.GetDays)
	return w
}