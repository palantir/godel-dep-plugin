// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package amalgomated

import (
	"github.com/palantir/godel-dep-plugin/generated_src/internal/github.com/golang/dep/amalgomated_flag"

	"github.com/palantir/godel-dep-plugin/generated_src/internal/github.com/golang/dep"
	"github.com/palantir/godel-dep-plugin/generated_src/internal/github.com/golang/dep/gps"
	"github.com/palantir/godel-dep-plugin/generated_src/internal/github.com/golang/dep/gps/pkgtree"
	"github.com/pkg/errors"
)

func (cmd *hashinCommand) Name() string		{ return "hash-inputs" }
func (cmd *hashinCommand) Args() string		{ return "" }
func (cmd *hashinCommand) ShortHelp() string	{ return "" }
func (cmd *hashinCommand) LongHelp() string	{ return "" }
func (cmd *hashinCommand) Hidden() bool		{ return true }

func (cmd *hashinCommand) Register(fs *flag.FlagSet)	{}

type hashinCommand struct{}

func (hashinCommand) Run(ctx *dep.Ctx, args []string) error {
	p, err := ctx.LoadProject()
	if err != nil {
		return err
	}

	sm, err := ctx.SourceManager()
	if err != nil {
		return err
	}
	sm.UseDefaultSignalHandling()
	defer sm.Release()

	params := p.MakeParams()
	params.RootPackageTree, err = pkgtree.ListPackages(p.ResolvedAbsRoot, string(p.ImportRoot))
	if err != nil {
		return errors.Wrap(err, "gps.ListPackages")
	}

	s, err := gps.Prepare(params, sm)
	if err != nil {
		return errors.Wrap(err, "prepare solver")
	}
	ctx.Out.Println(gps.HashingInputsAsString(s))
	return nil
}
