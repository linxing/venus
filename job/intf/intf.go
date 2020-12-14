package intf

import (
	"context"
)

type (
	ImportFromExcelReq struct {
		URL string `json:"url"`
	}

	Servlet interface {
		ImportFromExcel(context.Context, *ImportFromExcelReq) error
	}
)
