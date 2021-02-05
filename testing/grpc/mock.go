package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	venuspb "venus/pkg/venus"
)

type Server struct {
	venuspb.UnimplementedVenusServiceServer
	HelloResult func() (*venuspb.HelloResponse, error)
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Hello(ctx context.Context, req *venuspb.HelloRequest) (*venuspb.HelloResponse, error) {
	return s.HelloResult()
}

func GetSunServiceClient(ctx context.Context, srv *Server) (venuspb.VenusServiceClient, func() error) {

	buffer := 1024 * 1024
	listener := bufconn.Listen(buffer)

	s := grpc.NewServer()

	venuspb.RegisterVenusServiceServer(s, srv)
	go func() {
		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}()

	conn, _ := grpc.DialContext(ctx, "", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithInsecure())

	closer := func() error {
		s.Stop()
		return listener.Close()
	}

	client := venuspb.NewVenusServiceClient(conn)

	return client, closer
}
