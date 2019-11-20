// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/palantir/amalgomate/amalgomated"
	"github.com/palantir/godel/framework/pluginapi/v2/pluginapi"

	"github.com/palantir/godel-dep-plugin/cmd"
	amalgomateddep "github.com/palantir/godel-dep-plugin/generated_src"
)

func main() {
	if len(os.Args) >= 2 && os.Args[1] == amalgomated.ProxyCmdPrefix+"dep" {
		os.Args = append(os.Args[:1], os.Args[2:]...)
		amalgomateddep.Instance().Run("dep")
		os.Exit(0)
	}
	if ok := pluginapi.InfoCmd(os.Args, os.Stdout, cmd.PluginInfo); ok {
		return
	}
	os.Exit(cmd.Execute())
}
