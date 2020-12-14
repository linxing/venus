package intf

import (
	"context"
	"encoding/json"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

// job names
const (
	ImportFromExcelJobName = "import_from_excel"
)

type (
	Tasker struct {
		enqueuer *Enqueuer
	}

	TaskerInterface interface {
		Servlet
	}

	Enqueuer struct {
		*work.Enqueuer
	}
)

func (e *Enqueuer) Enqueue(ctx context.Context, jobName string, args map[string]interface{}) (*work.Job, error) {
	return e.Enqueuer.Enqueue(jobName, args)
}

func (e *Enqueuer) EnqueueIn(ctx context.Context, jobName string, secondsFromNow int64, args map[string]interface{}) (*work.ScheduledJob, error) {
	return e.Enqueuer.EnqueueIn(jobName, secondsFromNow, args)
}

func (e *Enqueuer) EnqueueUnique(ctx context.Context, jobName string, args map[string]interface{}) (*work.Job, error) {
	return e.Enqueuer.EnqueueUnique(jobName, args)
}

func (e *Enqueuer) EnqueueUniqueIn(ctx context.Context, jobName string, secondsFromNow int64, args map[string]interface{}) (*work.ScheduledJob, error) {
	return e.Enqueuer.EnqueueUniqueIn(jobName, secondsFromNow, args)
}

func NewTasker(redisPool *redis.Pool, namespace string) TaskerInterface {
	return &Tasker{
		enqueuer: &Enqueuer{work.NewEnqueuer(namespace, redisPool)},
	}
}

func convertToWorkQ(req interface{}) (work.Q, error) {

	buf, err := json.Marshal(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var q work.Q
	err = json.Unmarshal(buf, &q)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return q, nil
}

func (tasker *Tasker) ImportFromExcel(ctx context.Context, req *ImportFromExcelReq) error {

	q, err := convertToWorkQ(req)
	if err != nil {
		return err
	}

	_, err = tasker.enqueuer.Enqueue(ctx, ImportFromExcelJobName, q)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
