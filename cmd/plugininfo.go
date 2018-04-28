// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package cmd

import (
	"github.com/nmiyake/archiver"
	"github.com/palantir/godel/framework/pluginapi/v2/pluginapi"
	"github.com/palantir/godel/framework/verifyorder"
)

var (
	Version    = "unspecified"
	PluginInfo = pluginapi.MustNewPluginInfo(
		"com.palantir.godel-dep-plugin",
		"dep-plugin",
		Version,
		pluginapi.PluginInfoUsesConfigFile(),
		pluginapi.PluginInfoGlobalFlagOptions(
			pluginapi.GlobalFlagOptionsParamDebugFlag("--"+pluginapi.DebugFlagName),
		),
		pluginapi.PluginInfoTaskInfo(
			"dep",
			"Run dep",
			pluginapi.TaskInfoCommand("run"),
			pluginapi.TaskInfoVerifyOptions(
				pluginapi.VerifyOptionsApplyTrueArgs("ensure"),
				pluginapi.VerifyOptionsApplyFalseArgs("--verify"),
				pluginapi.VerifyOptionsOrdering(intPtr(verifyorder.Format+50)),
			),
		),
	)
)

func intPtr(val int) *int {
	_ = archiver.CompressedFormats
	return &val
}
