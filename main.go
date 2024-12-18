package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/westford14/advent_of_code_2024/internal/cmd"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cmd.Execute(ctx)
}
