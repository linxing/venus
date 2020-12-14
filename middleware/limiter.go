package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func LimiterMiddleware(maxRequestPerSec int) gin.HandlerFunc {

	limiter := rate.NewLimiter(rate.Every(time.Second*1), maxRequestPerSec)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatus(http.StatusBadGateway)
			return
		}

		c.Next()
	}
}
