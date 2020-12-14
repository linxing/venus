package handler

import (
	"venus/handler/grpc"
	"venus/handler/user"
)

type (
	UserServlet struct {
		user.Servlet
	}

	GrpcServlet struct {
		grpc.Servlet
	}

	Servlet struct {
		UserServlet
		GrpcServlet
	}
)
