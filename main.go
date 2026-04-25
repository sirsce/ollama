package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ollama/ollama/cmd"
)

func main() {
	// Initialize the root command and execute it.
	// All subcommands (serve, run, pull, push, list, etc.) are registered
	// within the cmd package.
	if err := cmd.NewCLI().ExecuteContext(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
