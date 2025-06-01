package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Muhammed19m/qbook/internal/app"
)

func main() {
	slog.Info("main: Starting")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err := app.Run(ctx)
		if !errors.Is(err, context.Canceled) && err != nil {
			slog.Error("app.Run:" + err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	slog.Info("main: Received signal " + (<-interrupt).String())
	cancel()
	slog.Info("main: Context canceled")
	time.Sleep(3 * time.Second)
}
