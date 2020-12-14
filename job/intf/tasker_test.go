package intf

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const jobNameSpacePrefix = "test_job_name_space"

type (
	testServlet struct {
		importFromExcel func(context.Context, *ImportFromExcelReq) error
	}
)

func (ts *testServlet) ImportFromExcel(ctx context.Context, req *ImportFromExcelReq) error {
	return ts.importFromExcel(ctx, req)
}

func recvCheckedChan(checkedChan <-chan bool, workPool *work.WorkerPool) bool {

	const MaxWaitTimes = 10

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for cnt := 0; cnt < MaxWaitTimes; cnt++ {
		select {
		case checked := <-checkedChan:
			return checked
		case <-ticker.C:
			workPool.Drain()
		}
	}

	return false
}

func initTestEnvForJob() *redis.Pool {

	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,

		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://127.0.0.1:6379/5")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func TestImportFromExcel(t *testing.T) {

	redisPool := initTestEnvForJob()
	defer redisPool.Close()

	now := time.Now()
	r := rand.New(rand.NewSource(now.UnixNano()))

	ctx := context.Background()

	req := ImportFromExcelReq{
		URL: fmt.Sprintf("url_%d", r.Int()),
	}

	nameSpace := fmt.Sprintf("%s_%d", jobNameSpacePrefix, r.Int())

	checkedChan := make(chan bool, 1)
	servlet := testServlet{
		importFromExcel: func(ctx context.Context, gotReq *ImportFromExcelReq) error {
			checkedChan <- assert.Equal(t, &req, gotReq)
			return nil
		},
	}

	tasker := NewTasker(redisPool, nameSpace)
	workPool := NewWorkerPool(redisPool, nameSpace, &servlet)

	err := tasker.ImportFromExcel(ctx, &req)
	require.NoError(t, err)

	workPool.Start()
	defer workPool.Stop()

	checked := recvCheckedChan(checkedChan, workPool)
	assert.True(t, checked)
}
