package recovery

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func Wrap() gin.HandlerFunc {
	return WrapWithCallback(noopFunc)
}

func WrapWithCallback(callback CallbackFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := extractBody(c)

		defer func() {
			r := recover()
			if r == nil {
				return
			}

			callback(c, body, r)

			fmt.Println(r)
			printBacktrace(20, 3)

			c.AbortWithStatus(http.StatusInternalServerError)
		}()

		c.Next()
	}
}

func extractBody(c *gin.Context) (body []byte) {
	body, _ = ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}

func printBacktrace(maxStacks, offset int) {
	stack := make([]uintptr, maxStacks)
	length := runtime.Callers(offset, stack[:])

	record := false

	for _, pc := range stack[:length] {
		f := runtime.FuncForPC(pc)
		if f == nil {
			continue
		}

		if !record {
			if f.Name() == "runtime.gopanic" {
				record = true
			}
			continue
		}

		file, line := f.FileLine(pc)
		fmt.Printf("\t%s [%s:%d]\n", f.Name(), file, line)
	}
}
