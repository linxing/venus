package venus

import (
	"context"
	"log"
	venuspb "venus/pkg/kdb/venus"
)

func (servlet *Servlet) Hello(ctx context.Context, req *venuspb.HelloRequest) (*venuspb.HelloResponse, error) {
	log.Println(req.Name)
	return nil, nil
}

func (servlet *Servlet) MakePhoneCall(ctx context.Context, req *venuspb.MakePhoneCallRequest) (*venuspb.MakePhoneCallResponse, error) {
	return nil, nil
}

func (servlet *Servlet) NotifyPhoneCallStatus(ctx context.Context, req *venuspb.NotifyPhoneCallStatusRequest) (*venuspb.NotifyPhoneCallStatusResponse, error) {
	return nil, nil
}
