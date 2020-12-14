package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"venus/env"
	"venus/env/global"
	venusgrpc "venus/grpc/venus"
	"venus/setting"
)

var (
	BuildTime = "20200101"
	GitTag    = "v0.0.1"
)

func setupForTeardown(grpc *grpc.Server) {
	go func() {
		sigCh := make(chan os.Signal, 1)

		signal.Reset(os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)

		<-sigCh

		grpc.GracefulStop()
	}()
}

func startPrometheus(grpcSrv *grpc.Server) {
	// Prometheus metrics
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(grpcSrv)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe("0.0.0.0:8882", nil)
	if err != nil {
		logrus.Fatalf("fail to serve prometheus: %s", err)
	}
}

func init() {

	if os.Args[1] == "-v" {
		fmt.Println("build at: " + BuildTime)
		fmt.Println("git tag: " + GitTag)
		os.Exit(0)
	}

	conf := flag.String("conf.ini", "setting/conf.test.ini", "config file")
	flag.Parse()

	if err := setting.Init(*conf); err != nil {
		panic(err)
	}

	global.Config = *setting.Config

	if err := env.Configure(); err != nil {
		logrus.Infof("Init env fail %+v", err)
		return
	}
}

func initTracer() *zipkin.Tracer {
	reporter := zkHttp.NewReporter(global.Config.ZipkinReporter)
	defer reporter.Close()

	endpoint, err := zipkin.NewEndpoint(global.Config.ZipkinEndpointName, global.Config.ZipkinEndPointHost)
	if err != nil {
		panic(err)
	}

	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		panic(err)
	}

	return nativeTracer
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", global.Config.ServiceGrpcAddr, global.Config.ServiceGrpcPort))
	if err != nil {
		panic(err)
	}

	logrusLogger := logrus.New()
	logrusEntry := logrus.NewEntry(logrusLogger)

	serverPayloadLoggingDecider := func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
		return fullMethodName != "/grpc.health.v1.Health/Check"
	}

	opts := []grpc.ServerOption{
		grpc_middleware.WithStreamServerChain(
			grpc_logrus.StreamServerInterceptor(logrusEntry),
			grpc_prometheus.StreamServerInterceptor,
			grpc_logrus.PayloadStreamServerInterceptor(logrusEntry, serverPayloadLoggingDecider),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.PayloadUnaryServerInterceptor(logrusEntry, serverPayloadLoggingDecider),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc.StatsHandler(zipkingrpc.NewServerHandler(initTracer())),
	}

	grpcSrv := grpc.NewServer(opts...)

	setupForTeardown(grpcSrv)

	healthSrv := health.NewServer()

	healthpb.RegisterHealthServer(grpcSrv, healthSrv)

	{
		venusGrpcServlet := venusgrpc.NewGrpcServlet(healthSrv)

		defer venusGrpcServlet.Close()

		venusGrpcServlet.RegisterGrpcServer(grpcSrv)
	}

	go startPrometheus(grpcSrv)

	if err := grpcSrv.Serve(lis); err != nil {
		logrus.Fatal(err)
	}

	logrus.Error("Exit grpc server")
}
