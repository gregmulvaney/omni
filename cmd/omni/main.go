package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gregmulvaney/omni/internal/cli"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	err := cli.Execute(ctx)
	defer stop()

	if err != nil {
		os.Exit(1)
	}
}
