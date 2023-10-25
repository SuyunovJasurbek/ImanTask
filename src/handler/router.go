package handler

import (
	"days_limit/src/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	s := service.NewService()
	h := NewHandler(s)
	
	// get token
	r.POST("/gettoken",h.SetToken)

	//get days
	r.POST("getdays",h.GetDays)
	return r
}