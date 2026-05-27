package cmd

import (
	"github.com/spf13/cobra"
)

type renderOpts struct {
	inputYAML string
}

func Render() *cobra.Command { _ = "STUB: not implemented"; return nil }

func renderYAML(opts renderOpts) error { _ = "STUB: not implemented"; return nil }
