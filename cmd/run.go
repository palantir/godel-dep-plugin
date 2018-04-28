// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/palantir/godel-dep-plugin/depplugin"
)

var runCmd = &cobra.Command{
	Use:   "run [flags] [args]",
	Short: "Format specified files (if no files are specified, format all project Go files)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if verifyFlagVal {
			return depplugin.Verify()
		}
		return depplugin.Run(args, cmd.OutOrStdout())
	},
}

func init() {
	runCmd.Flags().BoolVar(&verifyFlagVal, "verify", false, "verify files match formatting without applying formatting")
	rootCmd.AddCommand(runCmd)
}
