package readbody

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) (body []byte) {
	body, _ = ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}
