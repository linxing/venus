package mock

import (
	"venus/router"
)

type Servlet struct {
	router.User
	router.Grpc
}
