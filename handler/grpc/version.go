package grpc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*Servlet) GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "v0.0.1",
	})
}
