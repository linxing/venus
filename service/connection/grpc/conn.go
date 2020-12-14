package grpc

import (
	"context"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"venus/env/global"
)

func initTracer() (*zipkin.Tracer, error) {
	reporter := zkHttp.NewReporter(global.Config.ZipkinReporter)
	defer reporter.Close()

	endpoint, err := zipkin.NewEndpoint(global.Config.ZipkinEndpointName, global.Config.ZipkinEndPointHost)
	if err != nil {
		return nil, err
	}

	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return nil, err
	}

	return nativeTracer, nil
}

func NewConn(host string, seconds int, keepAliveSec int) (*grpc.ClientConn, error) {

	secs := time.Second * time.Duration(seconds)
	ctx := context.Background()

	ctx, cancelFunc := context.WithTimeout(ctx, secs)
	defer cancelFunc()

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	logrusEntry := logrus.NewEntry(log)

	clientPayloadLoggingDecider := func(ctx context.Context, fullMethodName string) bool {
		return fullMethodName != "/grpc.health.v1.Health/Check"
	}

	grpcRetryOpts := []grpc_retry.CallOption{
		grpc_retry.WithMax(3),
		grpc_retry.WithPerRetryTimeout(2 * time.Second),
	}

	unaryMW := []grpc.UnaryClientInterceptor{
		grpc_logrus.PayloadUnaryClientInterceptor(logrusEntry, clientPayloadLoggingDecider),
		grpc_retry.UnaryClientInterceptor(grpcRetryOpts...),
	}
	streamMW := []grpc.StreamClientInterceptor{
		grpc_logrus.PayloadStreamClientInterceptor(logrusEntry, clientPayloadLoggingDecider),
		grpc_retry.StreamClientInterceptor(grpcRetryOpts...),
	}

	tracer, err := initTracer()
	if err != nil {
		return nil, err
	}

	conn, err := grpc.DialContext(
		ctx,
		host,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithConnectParams(
			grpc.ConnectParams{
				MinConnectTimeout: secs,
			},
		),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(unaryMW...)),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(streamMW...)),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Duration(keepAliveSec) * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.WithStatsHandler(zipkingrpc.NewClientHandler(tracer)),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return conn, nil
}
