package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetGetToken(c *gin.Context) string {
	// set token
	token := uuid.NewString()
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", token, 36000, "/", "localhost", false, true)

	// get token
	any, ok := c.Get("token")
	t := any.(string)
	if ok {
		return t
	}
	return ""
}