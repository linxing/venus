package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"

	"venus/env/global"
)

func ZipkinTrace() gin.HandlerFunc {
	return func(c *gin.Context) {
		reporter := zkHttp.NewReporter(global.Config.ZipkinReporter)
		defer reporter.Close()

		endpoint, err := zipkin.NewEndpoint(global.Config.ZipkinEndpointName, global.Config.ZipkinEndPointHost)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		tracer := zipkinot.Wrap(nativeTracer)
		opentracing.SetGlobalTracer(tracer)

		span := tracer.StartSpan(c.FullPath())

		defer span.Finish()

		c.Next()
	}
}
