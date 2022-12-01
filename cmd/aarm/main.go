package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/shogo82148/aarm"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), trapSignals...)
	defer stop()

	exitCode, err := aarm.CLI(ctx)
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitCode)
}
