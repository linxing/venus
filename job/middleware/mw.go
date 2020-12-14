package middleware

import (
	"context"
	"encoding/json"

	"github.com/gocraft/work"
	"github.com/sirupsen/logrus"
)

type MiddlewareContext struct {
	Ctx context.Context
}

func (mc *MiddlewareContext) ContextMiddleware(job *work.Job, next work.NextMiddlewareFunc) error {
	mc.Ctx = context.Background()
	return next()
}

func LogMiddleware(job *work.Job, next work.NextMiddlewareFunc) error {
	defer func() {
		jobInfo, err := json.Marshal(job)
		if err != nil {
			return
		}

		logrus.Infof("End job [`%s`]", jobInfo)
	}()

	jobInfo, err := json.Marshal(job)
	if err != nil {
		return err
	}

	logrus.Infof("Start job [`%s`]", jobInfo)

	return next()
}
