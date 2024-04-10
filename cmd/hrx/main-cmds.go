// Copyright (c) 2024  The Go-CoreUtils Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"

	hrxutil "github.com/go-coreutils/hrx"

	"github.com/go-corelibs/notify"
	clPath "github.com/go-corelibs/path"
)

type opMode uint8

const (
	opError  opMode = 0
	opCreate opMode = iota + 1
	opExtract
	opList
)

func orCheck(argv ...bool) (ok bool) {
	for _, arg := range argv {
		if ok = arg; ok {
			return
		}
	}
	return
}

func withCheck(present bool, argv ...bool) (ok bool) {
	ok = present && orCheck(argv...)
	return
}

func prepareOptions(ctx *cli.Context) (opt *hrxutil.Options) {
	return &hrxutil.Options{
		All:        ctx.Bool(gAllFlag.Name),
		Recurse:    ctx.Bool(gRecurseFlag.Name),
		Boundary:   ctx.Int(gBoundaryFlag.Name),
		PruneDir:   ctx.Bool(gPruneDirFlag.Name),
		KeepEmpty:  ctx.Bool(gKeepEmptyFlag.Name),
		TrimPrefix: ctx.String(gTrimPrefixFlag.Name),
	}
}

func prepareOpMode(ctx *cli.Context) (op opMode, err error) {
	fCreate, fExtract, fList := ctx.Bool(gCreateFlag.Name),
		ctx.Bool(gExtractFlag.Name),
		ctx.Bool(gListFlag.Name)
	if present := fCreate; withCheck(present, fExtract, fList) {
		err = ErrMustOpMode
	} else if present {
		op = opCreate
	} else if present = fExtract; withCheck(present, fCreate, fList) {
		err = ErrMustOpMode
	} else if present {
		op = opExtract
	} else if present = fList; withCheck(present, fCreate, fExtract) {
		err = ErrMustOpMode
	} else if present {
		op = opList
	} else {
		err = ErrNeedOpMode
	}
	return
}

func action(ctx *cli.Context) (err error) {
	if ctx.Bool(gUsageFlag.Name) {
		//clcli.ShowUsageAndExit(ctx, 0)
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", AppUsageBrief)
		os.Exit(0)
	}
	if ctx.Bool(gVerboseFlag.Name) || ctx.Bool(gListFlag.Name) {
		hrxutil.Notifier = notify.New(notify.Info).Make()
	}

	var op opMode
	if op, err = prepareOpMode(ctx); err != nil {
		if op == opError && ctx.NumFlags() == 0 {
			cli.ShowAppHelpAndExit(ctx, 0)
		}
		return
	}

	argv := ctx.Args().Slice()

	switch op {
	case opCreate:
		return actionCreate(ctx, argv)
	case opExtract:
		return actionExtract(ctx, argv)
	case opList:
		return actionList(ctx, argv)
	case opError:
	}

	return
}

func actionList(ctx *cli.Context, argv []string) (err error) {
	if !ctx.IsSet(gFileFlag.Name) {
		err = ErrNeedArchive
		return
	}
	err = hrxutil.List(ctx.String(gFileFlag.Name), argv...)
	return
}

func actionCreate(ctx *cli.Context, argv []string) (err error) {
	argc := len(argv)
	var dst string
	if ctx.IsSet(gFileFlag.Name) {
		dst = ctx.String(gFileFlag.Name)
	} else if argc > 1 {
		dst = "archive.hrx"
	} else if argc == 1 && clPath.IsDir(argv[0]) {
		dst = filepath.Base(argv[0]) + ".hrx"
	} else {
		dst = clPath.Base(argv[0]) + ".hrx"
	}
	_, err = hrxutil.Create(prepareOptions(ctx), dst, argv...)
	return
}

func actionExtract(ctx *cli.Context, argv []string) (err error) {
	var src, dst string
	if ctx.IsSet(gFileFlag.Name) {
		if src = ctx.String(gFileFlag.Name); !clPath.IsFile(src) {
			err = ErrFileNotFound
			return
		}
	} else {
		err = ErrNeedArchive
		return
	}

	if ctx.IsSet(gDirFlag.Name) {
		if dst = ctx.String(gDirFlag.Name); !clPath.IsDir(dst) {
			err = ErrDirNotFound
			return
		}
	} else {
		dst = "."
	}

	err = hrxutil.Extract(prepareOptions(ctx), ctx.String(gFileFlag.Name), dst, argv...)
	return
}
