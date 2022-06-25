package util

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"k8s.io/klog/v2"
)

var FuncList []func()
var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{syscall.SIGTERM, syscall.SIGINT}

// SignalHandler registers for SIGTERM and SIGINT
func SignalHandler(funcList []func()) context.Context {
	close(onlyOneSignalHandler)
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		klog.Infoln("stoping server")
		cancel()
		// custom closing logic
		if len(funcList) > 0 {
			for _, f := range funcList {
				f()
			}
		}
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()
	return ctx
}

func Exit() {
	if len(FuncList) > 0 {
		for _, f := range FuncList {
			f()
		}
	}
	time.Sleep(1 * time.Second)
	os.Exit(0)
}
