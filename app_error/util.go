package app_error

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

// {status}.{package}.{struct/domain}...
var _APP_ERROR_PATTERN = regexp.MustCompile(`^([1-5]\d\d)((?:\.[a-z0-9_]+)+)$`)

func extractError(c *gin.Context) (err error) {
	if l := c.Errors.Last(); l != nil {
		err = l.Err
	}

	return
}

func parseAppError(err error) (ok bool, statusCode int, errId string) {
	m := _APP_ERROR_PATTERN.FindStringSubmatch(err.Error())
	if len(m) != 3 {
		return
	}

	ok = true
	statusCode, _ = strconv.Atoi(m[1])
	errId = m[2][1:]

	return
}

func ensureNonSuccessfulStatus(c *gin.Context) {
	statusCode := c.Writer.Status()
	if statusCode < 400 {
		statusCode = http.StatusInternalServerError
	}
	c.Status(statusCode)
}
