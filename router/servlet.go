package router

import (
	"github.com/gin-gonic/gin"
)

type Servlet interface {
	User
	Grpc
}

type User interface {
	UpdateUserPhoneNumberAndNickname(c *gin.Context)
	EditUserPassword(c *gin.Context)
	SigninUserByName(c *gin.Context)
	SigninUserByPhone(c *gin.Context)
	CreateUser(c *gin.Context)
	GetRefreshToken(c *gin.Context)
	GetUserInfo(c *gin.Context)
}

type Grpc interface {
	GetVersion(c *gin.Context)
}
