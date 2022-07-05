package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var headerRequestID string = "X-Request-Id"

//RequestID Get X-Request-Id With Request Header
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(headerRequestID)

		if requestID == "" {
			requestID = uuid.NewV4().String()
		}
		c.Set(headerRequestID, requestID)

		c.Writer.Header().Set(headerRequestID, requestID)
		c.Next()
	}
}
