package intf

import (
	"encoding/json"
	"reflect"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"

	"venus/job/middleware"
)

type WorkerContext struct {
	middleware.MiddlewareContext
}

func NewWorkerPool(redisPool *redis.Pool, namespace string, servlet Servlet, opts ...WorkerPoolOption) *work.WorkerPool {

	pool := work.NewWorkerPool(WorkerContext{}, 64, namespace, redisPool)

	for jobName, jobInfo := range map[string]struct {
		opts    work.JobOptions
		handler interface{}
	}{
		ImportFromExcelJobName: {
			opts: work.JobOptions{
				MaxFails:       3,
				MaxConcurrency: 50,
			},
			handler: servlet.ImportFromExcel,
		},
	} {
		wrappedHandler, _ := WrapJobHandler(jobInfo.handler)
		pool.JobWithOptions(jobName, jobInfo.opts, wrappedHandler)
	}

	for _, opt := range opts {
		opt(pool)
	}

	return pool
}

func WrapJobHandler(f interface{}) (interface{}, error) {

	fv := reflect.ValueOf(f)
	ft := fv.Type()

	err := validateJobHandlerType(ft)
	if err != nil {
		return nil, err
	}

	reqType := ft.In(1).Elem()

	return func(workerContext *WorkerContext, job *work.Job) error {

		buf, err := json.Marshal(job.Args)
		if err != nil {
			return errors.WithStack(err)
		}

		req := reflect.New(reqType)

		err = json.Unmarshal(buf, req.Interface())
		if err != nil {
			return err
		}

		ctxVal := reflect.ValueOf(&workerContext.MiddlewareContext.Ctx).Elem()

		results := fv.Call([]reflect.Value{ctxVal, req})
		if err, ok := results[0].Interface().(error); ok && err != nil {
			return err
		}

		return nil
	}, nil
}

func validateJobHandlerType(ft reflect.Type) error {
	if ft.Kind() == reflect.Func &&
		ft.NumIn() == 2 &&
		ft.In(1).Kind() == reflect.Ptr &&
		ft.In(1).Elem().Kind() == reflect.Struct &&
		ft.NumOut() == 1 &&
		ft.Out(0).String() == "error" {
		return nil
	}
	return errors.Errorf("handler type is %s, `func(*RequestStruct) error` expected", ft.String())
}
