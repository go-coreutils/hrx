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
	"github.com/urfave/cli/v2"
)

func init() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:               "version",
		Aliases:            []string{"V"},
		Usage:              "display the version",
		DisableDefaultText: true,
	}
	cli.HelpFlag = &cli.BoolFlag{
		Name:               "help",
		Usage:              "show detailed help",
		DisableDefaultText: true,
	}
}

var (
	gUsageFlag = &cli.BoolFlag{
		Name:    "usage",
		Usage:   "show brief usage",
		Aliases: []string{"h"},
	}
	gVerboseFlag = &cli.BoolFlag{
		Name:    "verbose",
		Usage:   "output progress",
		Aliases: []string{"v"},
	}

	gAllFlag = &cli.BoolFlag{
		Category: "SETTINGS",
		Name:     "all",
		Usage:    "include hidden files and directories",
		Aliases:  []string{"a"},
	}
	gFileFlag = &cli.StringFlag{
		Category: "SETTINGS",
		Name:     "archive",
		Usage:    "specify the archive file",
		Aliases:  []string{"f"},
	}
	gDirFlag = &cli.StringFlag{
		Category: "SETTINGS",
		Name:     "directory",
		Usage:    "specify the output directory",
		Aliases:  []string{"o"},
	}
	gTrimPrefixFlag = &cli.StringFlag{
		Category: "SETTINGS",
		Name:     "trim-prefix",
		Usage:    "trim given prefix from all pathnames",
		Aliases:  []string{"T"},
	}
	gPruneDirFlag = &cli.BoolFlag{
		Category: "SETTINGS",
		Name:     "prune-dir",
		Usage:    "remove the top directory from all pathnames",
		Aliases:  []string{"P"},
	}
	gRecurseFlag = &cli.BoolFlag{
		Category: "SETTINGS",
		Name:     "recurse",
		Usage:    "recurse into directories (default)",
		Aliases:  []string{"r"},
		Value:    true,
	}
	gKeepEmptyFlag = &cli.BoolFlag{
		Category: "SETTINGS",
		Name:     "keep-empty",
		Usage:    "include empty files and directories",
		Aliases:  []string{"k"},
	}
	gBoundaryFlag = &cli.IntFlag{
		Category: "SETTINGS",
		Name:     "boundary",
		Usage:    "specify the entry boundary size",
		Aliases:  []string{"b"},
	}

	gListFlag = &cli.BoolFlag{
		Category: "OPERATIONS",
		Name:     "list",
		Usage:    "list all archive entries",
		Aliases:  []string{"l"},
	}
	gCreateFlag = &cli.BoolFlag{
		Category: "OPERATIONS",
		Name:     "create",
		Usage:    "create a new archive",
		Aliases:  []string{"c"},
	}
	gUpdateFlag = &cli.BoolFlag{
		Category: "OPERATIONS",
		Name:     "update",
		Usage:    "add to an existing archive",
		Aliases:  []string{"u"},
	}
	gDeleteFlag = &cli.BoolFlag{
		Category: "OPERATIONS",
		Name:     "delete",
		Usage:    "remove from an existing archive",
		Aliases:  []string{"d"},
	}
	gExtractFlag = &cli.BoolFlag{
		Category: "OPERATIONS",
		Name:     "extract",
		Usage:    "extract an existing archive",
		Aliases:  []string{"x"},
	}
)
