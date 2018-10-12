// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package cmd

import (
	"github.com/palantir/godel/framework/pluginapi"
	"github.com/palantir/pkg/cobracli"
	"github.com/spf13/cobra"
)

var (
	debugFlagVal  bool
	verifyFlagVal bool
)

var rootCmd = &cobra.Command{
	Use:   "dep-plugin [flags] [files]",
	Short: "Run dep",
}

func Execute() int {
	return cobracli.ExecuteWithDebugVarAndDefaultParams(rootCmd, &debugFlagVal)
}

func init() {
	pluginapi.AddDebugPFlagPtr(rootCmd.PersistentFlags(), &debugFlagVal)
}
