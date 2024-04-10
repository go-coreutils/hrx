// Copyright (c) 2024  The Go-CoreLibs Authors
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
	"sort"

	"github.com/urfave/cli/v2"

	clcli "github.com/go-corelibs/cli"
)

var (
	AppName       = "hrx"
	AppVersion    = "v0.5.x"
	AppUsageText  = `hrx [global options] <-l|-c|-x> -f <archive> [pathnames...]`
	AppUsageBrief = "usage: " + AppUsageText + "\n" + `       hrx --help
       hrx --list -f existing.hrx
       hrx --create -f new.hrx <path> [paths...]
       hrx --extract -f existing.hrx [pathnames...]`
	AppDescription = `hrx is like the tar command except that the archives are human readable.

These archives are in a plain-text, human-friendly format for defining multiple
virtual text files in a single physical file, for situations when creating many
physical files is undesirable, such as defining datasets for unit test cases.

For details on the HRX file format, see: https://github.com/google/hrx

Note: this project has no affiliation with Google and is entirely independent.

OPERATIONS:

  There are currently three operational modes that can be performed:

    --list     (-l)
    --create   (-c)
    --extract  (-x)

  All modes require the --archive (-f) flag.

EXAMPLES:

  # list the contents of an archive named "custom-name.hrx"
  hrx -lf custom-name.hrx

  # create a new archive named "custom-name.hrx" with all files named "files"
  # with any extension
  hrx -cf custom-name.hrx files.*

  # extract an archive named "custom-name.hrx" into a sub-directory named
  # "custom-name"
  hrx -xf custom-name.hrx
`
)

var (
	gApp = cli.App{
		Name:                   AppName,
		Version:                AppVersion,
		Usage:                  "human-readable archive (.hrx) utility",
		UsageText:              AppUsageText,
		Description:            AppDescription,
		HideVersion:            false,
		HideHelpCommand:        true,
		UseShortOptionHandling: true,
		Action:                 action,
		Flags: []cli.Flag{
			gAllFlag,
			gDirFlag,
			gListFlag,
			gFileFlag,
			gUsageFlag,
			gCreateFlag,
			//gUpdateFlag,
			//gDeleteFlag,
			gExtractFlag,
			gRecurseFlag,
			gVerboseFlag,
			gPruneDirFlag,
			gBoundaryFlag,
			gKeepEmptyFlag,
			gTrimPrefixFlag,
		},
	}
)

func init() {
	cli.FlagStringer = clcli.NewFlagStringer().
		PruneRepeats(true).
		PruneDefaults(true).
		DetailsOnNewLines(true).
		Make()
}

func main() {
	sort.Sort(cli.FlagsByName(gApp.Flags))
	if err := gApp.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf("error: %v\n", err))
		os.Exit(1)
	}
}
