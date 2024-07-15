package settings

import (
	"errors"

	"github.com/gin-gonic/gin"

	errDomain "gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/domain/error"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *errDomain.Error:
				if errors.Is(err, errDomain.NotFoundErr) {
					ReturnNotFound(c, e)
				}
				ReturnStatusBadRequest(c, e)
			default:
				ReturnStatusInternalServerError(c, e)
			}
		}
	}
}
