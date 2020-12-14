package router

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	multierror "github.com/hashicorp/go-multierror"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	"venus/env/global"
	"venus/middleware"
	"venus/model"
)

type Router struct {
	GinEngine *gin.Engine
}

func NewRouter(s Servlet) *Router {

	//if global.Config.ServiceEnv == "Prod" {
	gin.SetMode(gin.ReleaseMode)
	//}

	engine := gin.New()

	{
		// gzip middleware
		engine.Use(gzip.Gzip(gzip.DefaultCompression))
		// gin revocer
		engine.Use(gin.Recovery())
	}

	if global.Config.ServiceEnv == "Prod" || global.Config.ServiceEnv == "Dev" {
		// prometheus middleware
		p := ginprometheus.NewPrometheus("http")
		p.ReqCntURLLabelMappingFn = (*gin.Context).HandlerName
		p.SetListenAddress(":23333")
		p.Use(engine)
		// log to file middleware
		engine.Use(middleware.LogToFile())
		// request limiter middleware
		engine.Use(middleware.LimiterMiddleware(global.Config.ServiceMaxRequestsPerSec))
		// zipkin trace middleware
		engine.Use(middleware.ZipkinTrace())
	}

	{
		apiGroup := engine.Group("/venus/")

		apiGroup.Use(cors.New(cors.Config{
			//TODO add domain
			AllowAllOrigins: true,
			AllowOrigins:    []string{},
			AllowMethods:    []string{"POST", "GET"},
			AllowHeaders:    []string{"Origin", "x-token"},
			ExposeHeaders:   []string{"Content-Length", "Content-Type"},
		}))

		if global.Config.ServiceEnv != "Prod" {
			apiGroup.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
				return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
					param.ClientIP,
					param.TimeStamp.Format("2006-01-02 15:04:05"),
					param.Method,
					param.Path,
					param.Request.Proto,
					param.StatusCode,
					param.Latency,
					param.Request.UserAgent(),
					param.ErrorMessage,
				)
			}))
		}

		{
			authGroup := apiGroup.Group("/auth")

			authGroup.POST("/user/signin_with_name", s.SigninUserByName)
			authGroup.POST("/user/signin_with_phone", s.SigninUserByPhone)
			authGroup.POST("/user/create", s.CreateUser)
		}

		// User
		{
			userGroup := apiGroup.Group("/user")
			userGroup.Use(middleware.JWTMiddleware())

			userGroup.POST("/info", s.UpdateUserPhoneNumberAndNickname)
			userGroup.POST("/password", s.EditUserPassword)
			userGroup.GET("/token/refresh", s.GetRefreshToken)

			userGroup.GET("/info", s.GetUserInfo)
		}
	}

	router := Router{
		GinEngine: engine,
	}

	return &router
}

func (*Router) Close() error {

	var redisErr error
	var dbError error

	if global.Redis.Pool != nil {
		redisErr = global.Redis.Pool.Close()
	}

	dbEngine := model.GetEngine()
	if dbEngine != nil {
		dbError = dbEngine.Close()
	}

	return multierror.Append(redisErr, dbError).ErrorOrNil()
}
