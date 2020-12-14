package intf

import (
	"github.com/gocraft/work"
)

type WorkerPoolOption func(*work.WorkerPool)

func GenericMwToWorkerContextMw(genericMw func(*work.Job, work.NextMiddlewareFunc) error) func(*WorkerContext, *work.Job, work.NextMiddlewareFunc) error {
	return func(ctx *WorkerContext, job *work.Job, next work.NextMiddlewareFunc) error {
		return genericMw(job, next)
	}
}

func WithMiddleware(mw func(*WorkerContext, *work.Job, work.NextMiddlewareFunc) error) WorkerPoolOption {
	return func(pool *work.WorkerPool) {
		pool.Middleware(mw)
	}
}
