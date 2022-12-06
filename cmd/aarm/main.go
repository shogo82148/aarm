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

	exitCode, err := aarm.CLI(ctx, os.Args[1:])
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitCode)
}
