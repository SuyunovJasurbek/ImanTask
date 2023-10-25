package middlewares

import (
	"sum_days/helper"
	"sum_days/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetToken(c *gin.Context) *string {

	// get token uuid example
	token := helper.GenerateUUID()

	//setting the token to the cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", token, 36000, "/", "", false, true)

	// return token
	return &token
}

func CheckToken(c *gin.Context) {
	// get token from header
	token := c.GetHeader("token")
	//1.empty check
	if token == "" {
		c.AbortWithStatusJSON(403, models.Error{
			Error: "token not fount",
		})
		return
	}

	//2.length check
	if len(token) != 36 {
		c.AbortWithStatusJSON(403, models.Error{
			Error: "token length did not match, length token 36 characters",
		})
		return
	}

	//3.cookie check
	token_string, err :=c.Cookie("token")
	if err!= nil {
		c.AbortWithStatusJSON(401, models.Error{
			Error: "token expired, pleace gettoken",
		})
		return
	}
	
	//4. token equalizing
	if token != token_string {
		c.AbortWithStatusJSON(403, models.Error{
			Error: "unauthorized token",
		})
		return
	}
}