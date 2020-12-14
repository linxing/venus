package venus

import (
	pb "venus/pkg/kdb/venus"
	"venus/service/connection/grpc"
)

func NewVenusService(host string, seconds int, keepAliveSec int) (pb.VenusServiceClient, error) {

	conn, err := grpc.NewConn(host, seconds, keepAliveSec)
	if err != nil {
		return nil, err
	}

	return pb.NewVenusServiceClient(conn), nil
}
