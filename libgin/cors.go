package libgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Header("Access-Control-Allow-Methods", "GET, PATCH, OPTIONS, POST, PUT, DELETE")
		if c.Request.Method == http.MethodOptions {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}
