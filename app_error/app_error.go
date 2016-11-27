package app_error

import (
	"github.com/gin-gonic/gin"

	"github.com/creasty/gin-contrib/readbody"
)

const (
	HEADER_KEY = "X-App-Error"
	JSON_KEY   = "error"
)

func Wrap() gin.HandlerFunc {
	return WrapWithCallback(noopFunc)
}

func WrapWithCallback(callback CallbackFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := readbody.Read(c)

		c.Next()

		err := extractError(c)
		if err == nil {
			return
		}

		if ok, statusCode, errId := parseAppError(err); ok {
			if c.IsAborted() {
				c.JSON(statusCode, gin.H{JSON_KEY: errId})
			} else {
				c.Status(statusCode)
				c.Header(HEADER_KEY, errId)
			}

			return
		}

		ensureNonSuccessfulStatus(c)

		callback(c, body, err)
	}
}
