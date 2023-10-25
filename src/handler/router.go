package handler

import (
	"sum_days/middlewares"
	"sum_days/src/service"

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