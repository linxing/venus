package venus

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	venuspb "venus/pkg/venus"
)

const ServiceName = "VenusService"

type Servlet struct {
	healthSrv *health.Server
}

func NewGrpcServlet(healthSrv *health.Server) *Servlet {

	servlet := Servlet{
		healthSrv: healthSrv,
	}

	servlet.setServingStatus(healthpb.HealthCheckResponse_SERVING)

	return &servlet
}

func (servlet *Servlet) RegisterGrpcServer(grpcSrv *grpc.Server) {
	venuspb.RegisterVenusServiceServer(grpcSrv, servlet)
}

func (servlet *Servlet) Close() error {
	servlet.setServingStatus(healthpb.HealthCheckResponse_NOT_SERVING)
	return nil
}

func (servlet *Servlet) setServingStatus(status healthpb.HealthCheckResponse_ServingStatus) {
	if healthSrv := servlet.healthSrv; healthSrv != nil {
		healthSrv.SetServingStatus(ServiceName, status)
	}
}
