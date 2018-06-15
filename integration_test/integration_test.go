// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package integration_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/nmiyake/pkg/dirs"
	"github.com/nmiyake/pkg/gofiles"
	"github.com/palantir/godel/framework/pluginapitester"
	"github.com/palantir/godel/pkg/products/v2/products"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const godelYML = `exclude:
  names:
    - "\\..+"
    - "vendor"
  paths:
    - "godel"
`

func TestRunDep(t *testing.T) {
	pluginPath, err := products.Bin("dep-plugin")
	require.NoError(t, err)

	projectDir, cleanup, err := dirs.TempDir(".", "")
	require.NoError(t, err)
	defer cleanup()

	origWd, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		err = os.Chdir(origWd)
		require.NoError(t, err)
	}()
	err = os.Chdir(projectDir)
	require.NoError(t, err)

	err = os.MkdirAll(path.Join(projectDir, "godel", "config"), 0755)
	require.NoError(t, err)
	err = ioutil.WriteFile(path.Join(projectDir, "godel", "config", "godel.yml"), []byte(godelYML), 0644)
	require.NoError(t, err)

	_, err = os.Stat("vendor")
	require.True(t, os.IsNotExist(err))

	outputBuf := &bytes.Buffer{}
	runPluginCleanup, err := pluginapitester.RunPlugin(pluginapitester.NewPluginProvider(pluginPath), nil, "run-dep", []string{"init"}, projectDir, false, outputBuf)
	defer runPluginCleanup()
	require.NoError(t, err, "Output: %s", outputBuf.String())

	_, err = os.Stat("vendor")
	assert.NoError(t, err, "Output: %s", outputBuf.String())
	_, err = os.Stat("Gopkg.lock")
	assert.NoError(t, err, "Output: %s", outputBuf.String())
	_, err = os.Stat("Gopkg.toml")
	assert.NoError(t, err, "Output: %s", outputBuf.String())
}

func TestDep(t *testing.T) {
	pluginPath, err := products.Bin("dep-plugin")
	require.NoError(t, err)

	projectDir, cleanup, err := dirs.TempDir(".", "")
	require.NoError(t, err)
	defer cleanup()

	origWd, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		err = os.Chdir(origWd)
		require.NoError(t, err)
	}()
	err = os.Chdir(projectDir)
	require.NoError(t, err)

	err = os.MkdirAll(path.Join(projectDir, "godel", "config"), 0755)
	require.NoError(t, err)
	err = ioutil.WriteFile(path.Join(projectDir, "godel", "config", "godel.yml"), []byte(godelYML), 0644)
	require.NoError(t, err)

	_, err = os.Stat("vendor")
	require.True(t, os.IsNotExist(err))

	outputBuf := &bytes.Buffer{}
	runPluginCleanup, err := pluginapitester.RunPlugin(pluginapitester.NewPluginProvider(pluginPath), nil, "run-dep", []string{"init"}, projectDir, false, outputBuf)
	defer runPluginCleanup()
	require.NoError(t, err, "Output: %s", outputBuf.String())

	specs := []gofiles.GoFileSpec{
		{
			RelPath: "foo.go",
			Src:     `package foo; import _ "github.com/pkg/errors";`,
		},
	}

	_, err = gofiles.Write(projectDir, specs)
	require.NoError(t, err)

	outputBuf = &bytes.Buffer{}
	runPluginCleanup, err = pluginapitester.RunPlugin(pluginapitester.NewPluginProvider(pluginPath), nil, "dep", nil, projectDir, false, outputBuf)
	defer runPluginCleanup()
	require.NoError(t, err, "Output: %s", outputBuf.String())

	_, err = os.Stat("vendor/github.com/pkg/errors")
	assert.NoError(t, err, "Output: %s", outputBuf.String())
}

func TestDepVerifyApplyFalseFails(t *testing.T) {
	pluginPath, err := products.Bin("dep-plugin")
	require.NoError(t, err)

	projectDir, cleanup, err := dirs.TempDir(".", "")
	require.NoError(t, err)
	defer cleanup()

	origWd, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		err = os.Chdir(origWd)
		require.NoError(t, err)
	}()
	err = os.Chdir(projectDir)
	require.NoError(t, err)

	err = os.MkdirAll(path.Join(projectDir, "godel", "config"), 0755)
	require.NoError(t, err)
	err = ioutil.WriteFile(path.Join(projectDir, "godel", "config", "godel.yml"), []byte(godelYML), 0644)
	require.NoError(t, err)

	_, err = os.Stat("vendor")
	require.True(t, os.IsNotExist(err))

	outputBuf := &bytes.Buffer{}
	runPluginCleanup, err := pluginapitester.RunPlugin(pluginapitester.NewPluginProvider(pluginPath), nil, "run-dep", []string{"init"}, projectDir, false, outputBuf)
	defer runPluginCleanup()
	require.NoError(t, err, "Output: %s", outputBuf.String())

	specs := []gofiles.GoFileSpec{
		{
			RelPath: "foo.go",
			Src:     `package foo; import _ "github.com/pkg/errors";`,
		},
	}

	_, err = gofiles.Write(projectDir, specs)
	require.NoError(t, err)

	outputBuf = &bytes.Buffer{}
	runPluginCleanup, err = pluginapitester.RunPlugin(pluginapitester.NewPluginProvider(pluginPath), nil, "dep", []string{"--verify"}, projectDir, false, outputBuf)
	defer runPluginCleanup()
	require.Error(t, err)
	// if verification is due to vendor state not matching expected state, there should not be error output
	assert.Equal(t, "", outputBuf.String())
}

func TestDepVerifyApplyFalseExecErrorFails(t *testing.T) {
	pluginPath, err := products.Bin("dep-plugin")
	require.NoError(t, err)

	projectDir, cleanup, err := dirs.TempDir(".", "")
	require.NoError(t, err)
	defer cleanup()

	origWd, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		err = os.Chdir(origWd)
		require.NoError(t, err)
	}()
	err = os.Chdir(projectDir)
	require.NoError(t, err)

	err = os.MkdirAll(path.Join(projectDir, "godel", "config"), 0755)
	require.NoError(t, err)
	err = ioutil.WriteFile(path.Join(projectDir, "godel", "config", "godel.yml"), []byte(godelYML), 0644)
	require.NoError(t, err)

	_, err = os.Stat("vendor")
	require.True(t, os.IsNotExist(err))

	outputBuf := &bytes.Buffer{}
	runPluginCleanup, err := pluginapitester.RunPlugin(pluginapitester.NewPluginProvider(pluginPath), nil, "run-dep", []string{"init"}, projectDir, false, outputBuf)
	defer runPluginCleanup()
	require.NoError(t, err, "Output: %s", outputBuf.String())

	// note: import path is not quoted, so this does not compile
	specs := []gofiles.GoFileSpec{
		{
			RelPath: "foo.go",
			Src:     `package foo; import _ github.com/pkg/errors;`,
		},
	}

	files, err := gofiles.Write(projectDir, specs)
	require.NoError(t, err)

	outputBuf = &bytes.Buffer{}
	runPluginCleanup, err = pluginapitester.RunPlugin(pluginapitester.NewPluginProvider(pluginPath), nil, "dep", []string{"--verify"}, projectDir, false, outputBuf)
	defer runPluginCleanup()
	require.Error(t, err)
	// if verification is due to error besides vendor state not matching expected state, verification output should
	// include error output
	assert.Equal(t, fmt.Sprintf(`Error: found 1 errors in the package tree:
%s:1:23: expected 'STRING', found 'IDENT' github
`, files["foo.go"].Path), outputBuf.String())
}
