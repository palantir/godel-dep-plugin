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
	Long: `Executes "dep ensure" using the bundled version of dep with the provided flags and arguments. The "--" separator must 
be used before specifying any flags for the "dep" program. For example, "./godelw dep -- -v" executes "dep ensure -v".`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if verifyFlagVal {
			return depplugin.Verify()
		}
		return depplugin.Run(append([]string{"ensure"}, args...), cmd.OutOrStdout())
	},
}

func init() {
	depCmd.Flags().BoolVar(&verifyFlagVal, "verify", false, "verify that `dep ensure` would make no changes")
	rootCmd.AddCommand(depCmd)
}
