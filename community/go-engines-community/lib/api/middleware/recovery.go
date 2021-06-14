package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"strings"

	"git.canopsis.net/canopsis/go-engines/lib/api/auth"
	"git.canopsis.net/canopsis/go-engines/lib/api/common"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var (
	MissingUserKeyPanicMsg = "Key \"" + auth.UserKey + "\" does not exist"
)

// Recovery recovers from any panics and if there was one it logs error and writes InternalServerError response.
func Recovery(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var err, errToLog error
				var ok bool
				if err, ok = r.(error); ok {
					errToLog = err
				} else {
					errToLog = fmt.Errorf("%v", r)
				}

				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var se *os.SyscallError
				if err != nil && errors.As(err, &se) {
					if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
						logger.Err(err).Msgf("panic recovered")
						_ = c.Error(err)
						c.Abort()
						return
					}
				}

				if errMsg, ok := r.(string); ok {
					// Handle unauth response
					if errMsg == MissingUserKeyPanicMsg {
						c.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
						return
					}
				}

				logger.Err(errToLog).Msgf("panic recovered\n%s\n", debug.Stack())
				c.AbortWithStatusJSON(http.StatusInternalServerError, common.InternalServerErrorResponse)
			}
		}()
		c.Next()
	}
}
