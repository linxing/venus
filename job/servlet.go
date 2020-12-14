package job

import (
	jobintf "venus/job/intf"
)

type (
	servlet struct {
	}
)

func NewServlet() jobintf.Servlet {
	servlet := servlet{}
	return &servlet
}
