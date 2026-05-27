package cmd

import (
	"github.com/K-Phoen/grabana"
	"github.com/spf13/cobra"
)

type applyOpts struct {
	inputYAML         string
	destinationFolder string
	grafanaHost       string
	grafanaToken      string
}

func Apply() *cobra.Command { _ = "STUB: not implemented"; return nil }

func applyYAML(opts applyOpts) error { _ = "STUB: not implemented"; return nil }

func grabanaClient(opts applyOpts) *grabana.Client { _ = "STUB: not implemented"; return nil }
