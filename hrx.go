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

// Package hrx provides the primary functionality of the go-coreutils hrx
// command.
package hrx

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/go-corelibs/hrx"
	clPath "github.com/go-corelibs/path"
	"github.com/go-corelibs/tdata"
)

const (
	OpWrote    = "wrote"
	OpListing  = "listing"
	OpArchived = "archived"
)

// Options are the complete configurable options for Create and Extract
type Options struct {
	// All specifies to include hidden files and directories
	All bool
	// Recurse specifies to traverse directories recursively
	Recurse bool
	// PruneDir specifies to prune the top directory from files added to the
	// Archive
	PruneDir bool
	// Boundary specifies the Archive boundary size to use
	Boundary int
	// TrimPrefix specifies an arbitrary string prefix to trim from files
	// added to the Archive
	TrimPrefix string
	// KeepEmpty specifies to include empty directories when added to the
	// Archive
	KeepEmpty bool
}

// List displays a list of pathnames within an existing `src` archive file.
// If any `pathnames` are given, List will only display those pathnames given
// that exist within the `src` archive file
func List(src string, pathnames ...string) (err error) {
	var a hrx.Archive
	if a, err = prepareExistingSrc(src); err != nil {
		return
	}
	safeResetReporting()
	tc := tdata.NewTestCheck(len(pathnames) > 0, pathnames...)
	for _, entry := range a.Entries() {
		pathname := entry.GetPathname()
		if tc.NotPresent(pathname) {
			continue
		}
		reporterFn(src, pathname, OpListing, entry.GetBody())
	}
	printSummary(a, OpListing, src)
	return
}

// Create produces an archive with the given `pathnames`, according to the
// Options given and writes the archive to the `dst` path
func Create(opt *Options, dst string, pathnames ...string) (a hrx.Archive, err error) {
	if len(pathnames) == 0 {
		err = ErrPathRequired
		return
	} else if a, err = prepareNewSrc(dst); err != nil {
		a = nil
		return
	}
	safeResetReporting()
	a.SetReporter(reporterFn)
	opt = prepareOptions(opt)
	_ = a.SetBoundary(opt.Boundary)

	for _, arg := range pathnames {

		if clPath.IsFile(arg) {
			if err = readFileAndSet(a, arg, preparePath(opt, arg)); err != nil {
				a = nil
				return
			}
			continue
		}

		// is a directory
		var files []string
		if opt.Recurse {
			if files, err = clPath.ListAllFiles(arg, opt.All); err != nil {
				a = nil
				return
			}
		} else if files, err = clPath.ListFiles(arg, opt.All); err != nil {
			a = nil
			return
		}
		if len(files) == 0 {
			// no files found, empty directory or not recursive
			if opt.KeepEmpty {
				_ = a.Set(preparePath(opt, arg)+"/", "", "")
			}
			continue
		}
		for _, name := range files {
			if err = readFileAndSet(a, name, preparePath(opt, name)); err != nil {
				if isCreateFileErrIgnored(err) {
					continue // skip
				}
				a = nil
				return
			}
		}

	}

	if err = a.WriteFile(dst); err != nil {
		a = nil
	} else {
		printSummary(a, OpArchived, dst)
	}
	return
}

// Extract takes an existing `src` archive and extracts it to the `dst`
// directory, according to the Options given. If `pathnames` are provided,
// only those pathnames that exist within the `src` archive are extracted
func Extract(opt *Options, src, dst string, pathnames ...string) (err error) {
	var a hrx.Archive
	if a, err = prepareExistingSrc(src); err != nil {
		return
	}
	safeResetReporting()
	a.SetReporter(reporterFn)
	opt = prepareOptions(opt)

	if dst == "" {
		dst = "./" + strings.TrimSuffix(filepath.Base(src), ".hrx")
	}

	if err = clPath.MkdirAll(dst); err != nil {
		return
	}

	if opt.PruneDir || opt.TrimPrefix != "" {
		tc := tdata.NewTestCheck(len(pathnames) > 0, pathnames...)

		for _, pathname := range a.List() {
			if tc.NotPresent(pathname) {
				continue
			}

			entry := a.Entry(pathname)
			pruned := pruneName(pathname, opt.TrimPrefix, opt.PruneDir)
			destination := filepath.Join(dst, pruned)
			if entry.IsDir() {
				if err = os.MkdirAll(destination, 0770); err != nil {
					return
				}
				reporterFn(src, destination, hrx.OpCreated, destination)
			} else if entry.IsFile() {
				dirname := filepath.Dir(destination)
				if err = os.MkdirAll(dirname, 0770); err != nil {
					return
				} else if err = os.WriteFile(destination, []byte(entry.GetBody()), 0660); err != nil {
					return
				}
				reporterFn(src, destination, OpWrote)
			}
		}

	} else if err = a.ExtractTo(dst, pathnames...); err != nil {
		return
	}
	printSummary(a, hrx.OpExtracted, dst)
	return
}
