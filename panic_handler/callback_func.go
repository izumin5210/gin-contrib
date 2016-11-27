package panic_handler

import (
	"github.com/creasty/panicsync"
	"github.com/gin-gonic/gin"
)

type CallbackFunc func(*gin.Context, []byte, *panicsync.Info)

var defaultFunc = func(c *gin.Context, body []byte, info *panicsync.Info) {
	info.Print()
}
