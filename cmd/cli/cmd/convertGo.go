package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type convertGoOpts struct {
	inputJSON string
}

func ConvertGo(logger *zap.Logger) *cobra.Command { _ = "STUB: not implemented"; return nil }

func convertGo(logger *zap.Logger, opts convertGoOpts) error { _ = "STUB: not implemented"; return nil }
