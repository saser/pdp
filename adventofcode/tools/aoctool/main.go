package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Saser/pdp/adventofcode/tools/aoctool/cmd"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := cmd.RootCmd().ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
