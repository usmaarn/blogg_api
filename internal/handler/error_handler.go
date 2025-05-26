package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmaarn/blogg_api/package/utils/response"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			c.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error(), nil))
		}
	}
}

func PanicHandler() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		switch any(err).(type) {
		case response.Response:
			c.AbortWithStatusJSON(http.StatusOK, err)
		case string, int, nil:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
	})
}
