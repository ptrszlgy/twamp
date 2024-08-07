package main

import (
	"flag"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/ptrszlgy/twamp/pkg/server"
	"go.uber.org/zap"
)

var log logr.Logger

func main() {
	listenPtr := flag.String("listen", "localhost:862", "listen address")
	flag.Parse()

	zapLog, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	log = zapr.NewLogger(zapLog)

	server.SetLogger(log)

	server.SetupSignals()

	log.Info("Start")

	err = server.ServeTwamp(*listenPtr)
	if err != nil {
		log.Error(err, "could not start TWAMP server")
	}

}
