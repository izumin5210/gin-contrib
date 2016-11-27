package app_error

import (
	"github.com/gin-gonic/gin"
)

type CallbackFunc func(*gin.Context, []byte, error)

var noopFunc = func(c *gin.Context, body []byte, err error) {}
