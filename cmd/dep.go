// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/palantir/godel-dep-plugin/depplugin"
)

var depCmd = &cobra.Command{
	Use:   "dep [flags] [args]",
	Short: "Runs dep ensure for the project",
	RunE: func(cmd *cobra.Command, args []string) error {
		if verifyFlagVal {
			return depplugin.Verify()
		}
		return depplugin.Run([]string{"ensure"}, cmd.OutOrStdout())
	},
}

func init() {
	depCmd.Flags().BoolVar(&verifyFlagVal, "verify", false, "verify files match formatting without applying formatting")
	rootCmd.AddCommand(depCmd)
}
