package panic_handler

import (
	"github.com/creasty/panicsync"
	"github.com/gin-gonic/gin"

	"github.com/creasty/gin-contrib/readbody"
)

const CONTEXT_NAME = "PanicHandler"

func Wrap() gin.HandlerFunc {
	return WrapWithCallback(defaultFunc)
}

func WrapWithCallback(callback CallbackFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := readbody.Read(c)

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
