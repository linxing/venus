package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gocraft/work"
	"github.com/pkg/errors"

	prom "github.com/prometheus/client_golang/prometheus"
)

type ServerMetrics struct {
	serviceName string

	serverStartedCounter     *prom.CounterVec
	serverProcessedCounter   *prom.CounterVec
	serverProcessedHistogram *prom.HistogramVec
}

func NewServerMetrics(servlet interface{}) (*ServerMetrics, error) {
	metrics := ServerMetrics{
		serviceName: fmt.Sprintf("%T", servlet),
		serverStartedCounter: prom.NewCounterVec(
			prom.CounterOpts{
				Name: "server_started_counter",
				Help: "Total number of jobs started on the server.",
			},
			[]string{"job_service", "job_method"},
		),
		serverProcessedCounter: prom.NewCounterVec(
			prom.CounterOpts{
				Name: "server_processed_counter",
				Help: "Total number of jobs processed.",
			},
			[]string{"job_service", "job_method", "success"},
		),
		serverProcessedHistogram: prom.NewHistogramVec(
			prom.HistogramOpts{
				Name:    "server_processed_histogram",
				Help:    "Histogram of processing latency of job.",
				Buckets: prom.DefBuckets,
			},
			[]string{"job_service", "job_method"},
		),
	}
	for _, collector := range []prom.Collector{
		metrics.serverStartedCounter,
		metrics.serverProcessedCounter,
		metrics.serverProcessedHistogram,
	} {
		err := prom.Register(collector)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return &metrics, nil
}

func (metrics *ServerMetrics) Middleware(job *work.Job, next work.NextMiddlewareFunc) error {

	startTime := time.Now()

	metrics.serverStartedCounter.WithLabelValues(metrics.serviceName, job.Name).Inc()

	defer func() {
		metrics.serverProcessedCounter.WithLabelValues(metrics.serviceName, job.Name, strconv.FormatBool(job.Fails == 0)).Inc()
		metrics.serverProcessedHistogram.WithLabelValues(metrics.serviceName, job.Name).Observe(time.Since(startTime).Seconds())
	}()

	return next()
}
