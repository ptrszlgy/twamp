package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/ptrszlgy/twamp/pkg/server"
)

// const maxuint64 = ^uint(0)

func main() {
	listenPtr := flag.String("listen", "localhost:862", "listen address")
	flag.Parse()

	server.SetupSignals()

	err := server.ServeTwamp(*listenPtr, *udpStart)
	if err != nil {
		glog.Error(err)
	}
	server.Cleanup()
}
