package cmd

import (
	"github.com/spf13/cobra"
)

type validateOpts struct {
	inputYAML string
}

func Validate() *cobra.Command { _ = "STUB: not implemented"; return nil }

func validateYAML(opts validateOpts) error { _ = "STUB: not implemented"; return nil }
