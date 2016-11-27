package frontend

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const NAME = "/creasty/gin-contrib/frontend.Wrap"

func Wrap(fn gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			c.Next()
			return
		}

		c.Next()

		if strings.Contains(c.HandlerName(), NAME) {
			fn(c)
		}
	}
}
