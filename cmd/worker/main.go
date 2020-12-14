package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	"venus/env"
	"venus/env/global"
	"venus/job"
	jobintf "venus/job/intf"
	jobmw "venus/job/middleware"
	"venus/setting"
)

var (
	BuildTime = "20200101"
	GitTag    = "v0.0.1"
)

func waitForTeardown() {

	sigCh := make(chan os.Signal, 1)

	signal.Reset(os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)

	<-sigCh
}

func startPrometheus() {
	// Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe("0.0.0.0:8883", nil)
	if err != nil {
		logrus.Fatalf("fail to serve prometheus: %s", err)
	}
}

func init() {

	if os.Args[1] == "-v" {
		fmt.Println("build at: " + BuildTime)
		fmt.Println("git tag: " + GitTag)
		os.Exit(0)
	}

	conf := flag.String("conf.ini", "setting/conf.test.ini", "config file")
	flag.Parse()

	if err := setting.Init(*conf); err != nil {
		panic(err)
	}

	global.Config = *setting.Config

	if err := env.Configure(); err != nil {
		log.Printf("[info] init env fail %+v", err)
		return
	}
}

func main() {

	servlet := job.NewServlet()

	metrics, err := jobmw.NewServerMetrics(servlet)
	if err != nil {
		panic(err)
	}

	workerPool := jobintf.NewWorkerPool(global.Redis.Pool, global.Config.JobNamespace, servlet,
		jobintf.WithMiddleware(jobintf.GenericMwToWorkerContextMw(jobmw.LogMiddleware)),
		jobintf.WithMiddleware(jobintf.GenericMwToWorkerContextMw(metrics.Middleware)))

	workerPool.Start()

	go startPrometheus()

	waitForTeardown()

	workerPool.Stop()
}
