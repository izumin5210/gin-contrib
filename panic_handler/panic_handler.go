package panic_handler

import (
	"bytes"
	"io/ioutil"

	"github.com/creasty/panicsync"
	"github.com/gin-gonic/gin"
)

const CONTEXT_NAME = "PanicHandler"

func Wrap() gin.HandlerFunc {
	return WrapWithCallback(defaultFunc)
}

func WrapWithCallback(callback CallbackFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := []byte{}
		{
			body, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}

		ph := panicsync.NewHandler(func(info *panicsync.Info) {
			callback(c, body, info)
		})
		defer ph.Done()

		c.Set(CONTEXT_NAME, ph)

		c.Next()
	}
}

func Get(c *gin.Context) *panicsync.Handler {
	v := c.MustGet(CONTEXT_NAME)

	ph, ok := v.(*panicsync.Handler)
	if !ok {
		panic("Cannot retrive `" + CONTEXT_NAME + "` from context")
	}

	return ph
}
