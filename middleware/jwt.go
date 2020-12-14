package middleware

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"venus/env/global"
)

func validateToken(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(global.Config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invaid token")
	}

	return token.Claims, nil
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("x-token")
		if authorization == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := validateToken(authorization)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claimsData := claims.(jwt.MapClaims)

		exp := int64(claimsData["exp"].(float64))
		if time.Now().After(time.Unix(exp, 0)) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userID", claimsData["user_id"].(float64))
		c.Set("role", claimsData["role"].(float64))
		c.Set("nickname", claimsData["nickname"].(string))
		c.Next()
	}
}
