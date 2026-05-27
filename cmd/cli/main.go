package main

import (
	"fmt"
	"os"

	"github.com/K-Phoen/grabana/cmd/cli/cmd"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var version = "SNAPSHOT"

func main() {
	logger, err := createLogger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create logger: %s", err)
		os.Exit(1)
	}

	root := &cobra.Command{Use: "grabana"}
	root.Version = version
	root.SilenceUsage = true

	root.AddCommand(cmd.Apply())
	root.AddCommand(cmd.Validate())
	root.AddCommand(cmd.SelfUpdate(version))
	root.AddCommand(cmd.Render())
	root.AddCommand(cmd.ConvertGo(logger))

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func createLogger() (*zap.Logger, error) { _ = "STUB: not implemented"; return nil, nil }
