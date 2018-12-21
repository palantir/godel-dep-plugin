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
	Short: "Run dep with the provided arguments",
	Long: `Executes "dep" using the bundled version of dep with the provided flags and arguments. The "--" separator must be used 
before specifying any flags for the "dep" program. For example, "./godelw run-dep -- -h" executes "dep -h".`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return depplugin.Run(args, cmd.OutOrStdout())
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
