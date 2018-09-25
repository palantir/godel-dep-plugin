// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package depplugin

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/kardianos/osext"
	"github.com/palantir/amalgomate/amalgomated"
	"github.com/pkg/errors"
)

func Run(args []string, stdout io.Writer) error {
	pathToSelf, err := osext.Executable()
	if err != nil {
		return errors.Wrapf(err, "failed to determine path to self")
	}

	cmd := exec.Command(pathToSelf, append([]string{amalgomated.ProxyCmdPrefix + "dep"}, args...)...)
	cmd.Stdout = stdout
	cmd.Stderr = stdout
	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			// if error is not an exit error, wrap it
			return errors.Wrapf(err, "failed to execute command %v", cmd.Args)
		}
		// otherwise, return empty error
		return fmt.Errorf("")
	}
	return nil
}

func Verify() error {
	pathToSelf, err := osext.Executable()
	if err != nil {
		return errors.Wrapf(err, "failed to determine path to self")
	}

	args := []string{
		"ensure",
		"-no-vendor",
		"-dry-run",
	}
	cmd := exec.Command(pathToSelf, append([]string{amalgomated.ProxyCmdPrefix + "dep"}, args...)...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			// if error is not an exit error, wrap it
			return errors.Wrapf(err, "failed to execute command %v", cmd.Args)
		}
		// otherwise, error with output
		return errors.Errorf(strings.TrimSuffix(string(output), "\n"))
	} else if strings.Contains(string(output), "Would have written") {
		// If dep ensure would make changes, return the output as an error
		return errors.Errorf(strings.TrimSuffix(string(output), "\n"))
	}
	return nil
}
