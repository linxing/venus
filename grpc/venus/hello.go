package venus

import (
	"context"
	venuspb "venus/pkg/venus"
)

func (servlet *Servlet) Hello(ctx context.Context, req *venuspb.HelloRequest) (*venuspb.HelloResponse, error) {
	return &venuspb.HelloResponse{
		Result: "Welcode" + req.GetName(),
	}, nil
}
