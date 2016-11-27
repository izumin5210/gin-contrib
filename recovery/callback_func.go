package recovery

import (
	"github.com/gin-gonic/gin"
)

type CallbackFunc func(*gin.Context, []byte, interface{})

var noopFunc = func(c *gin.Context, body []byte, err interface{}) {}
