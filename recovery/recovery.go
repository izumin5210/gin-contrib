package recovery

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/creasty/gin-contrib/readbody"
)

type Config struct {
	HidePanic    bool
	CallbackFunc CallbackFunc
}

func Wrap() gin.HandlerFunc {
	return WrapWithCallback(noopFunc)
}

func WrapWithCallback(callback CallbackFunc) gin.HandlerFunc {
	return WrapWithConfig(&config{CallbackFunc: CallbackFunc})
}

func WrapWithConfig(c *Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := readbody.Read(c)

		defer func() {
			r := recover()
			if r == nil {
				return
			}

			config.callback(c, body, r)

			if config.HidePanic {
				fmt.Println(r)
			}
			printBacktrace(20, 3)

			c.AbortWithStatus(http.StatusInternalServerError)
		}()

		c.Next()
	}
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
