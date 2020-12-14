package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	casbinsrv "venus/service/casbin"
)

func RBACMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		sub := fmt.Sprintf("%d", int64(role.(float64)))

		e, err := casbinsrv.NewCasbinAuth()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		enforcer := e.GetEnforcer()

		isGrant, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if !isGrant {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
