package util

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type StartFunc func()

type StopFunc func(context.Context)

type ShutdownFunc func(stopFuncs ...StopFunc)

func NewShutdownContext() (context.Context, ShutdownFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	quite := make(chan os.Signal, 1)
	signal.Notify(quite, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-quite
		log.Print("shutting down application...")

		cancel()
	}()

	shutdownFunc := func(stopFuncs ...StopFunc) {
		<-ctx.Done()
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()

		for _, f := range stopFuncs {
			f(shutdownCtx)
		}

		log.Println("application shut down...")
	}

	return ctx, shutdownFunc
}