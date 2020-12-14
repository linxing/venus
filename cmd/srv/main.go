package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"venus/env"
	"venus/env/global"
	"venus/handler"
	"venus/router"
	"venus/setting"
)

var (
	BuildTime = "20200101"
	GitTag    = "v0.0.1"
)

func waitForTeardown(router *router.Router) {

	sigCh := make(chan os.Signal, 1)

	signal.Reset(os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)

	<-sigCh

	if err := router.Close(); err != nil {
		logrus.WithError(err).Println("router close")
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
		logrus.Infof("Init env fail %+v", err)
		return
	}
}

func main() {

	router := router.NewRouter(new(handler.Servlet))

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", global.Config.ServiceHTTPAddr, global.Config.ServicePort),
		Handler: router.GinEngine,
	}

	logrus.Infof("Start http serve on %s listening %d", global.Config.ServiceHTTPAddr, global.Config.ServicePort)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logrus.Errorf("Start fail %+v", err)
			panic(err)
		}
	}()

	waitForTeardown(router)
}
