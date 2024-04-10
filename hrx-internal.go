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

package hrx

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-corelibs/hrx"
	"github.com/go-corelibs/notify"
	"github.com/go-corelibs/path"
)

var (
	// Notifier is the user notice output handler
	Notifier notify.Notifier
)

func init() {
	if Notifier == nil {
		Notifier = notify.New(notify.Error).Make()
	}
}

func validateNewSrc(src string) (err error) {
	if path.Exists(src) && !path.IsRegularFile(src) {
		err = fmt.Errorf("%v: %q", ErrNotRegular, src)
	}
	return
}

func validateExistingFile(src string) (err error) {
	if !path.Exists(src) {
		err = fmt.Errorf("%w: %q", ErrFileNotFound, src)
		return
	} else if !path.IsRegularFile(src) {
		err = fmt.Errorf("%w: %q", ErrNotRegular, src)
		return
	}
	return
}

func prepareNewSrc(src string) (a hrx.Archive, err error) {
	if err = validateNewSrc(src); err == nil {
		a = hrx.New(src, "")
	}
	return
}

func prepareExistingSrc(src string) (a hrx.Archive, err error) {
	if err = validateExistingFile(src); err == nil {
		a, err = hrx.ParseFile(src)
	}
	return
}

func readFileAndSet(a hrx.Archive, src, name string) (err error) {
	if err = validateExistingFile(src); err == nil {
		var data []byte
		if data, err = os.ReadFile(src); err == nil {
			err = a.Set(name, string(data), "")
		}
	}
	return
}

func hasPathPrefix(input string) bool {
	if size := len(input); size > 0 && input[0] == '/' {
		return true
	} else if size > 1 && input[0:2] == "./" {
		return true
	} else if size > 2 && input[0:3] == "../" {
		return true
	}
	return false
}

func trimPathPrefix(input string) string {
	if size := len(input); size > 0 && input[0] == '/' {
		return input[1:]
	} else if size > 1 && input[0:2] == "./" {
		return input[2:]
	} else if size > 2 && input[0:3] == "../" {
		return input[3:]
	}
	return input
}

func trimPathPrefixes(input string) string {
	tmp := input
	for ; hasPathPrefix(tmp); tmp = trimPathPrefix(tmp) {
	}
	return tmp
}

func pruneTopDir(input string) string {
	if parts := strings.Split(trimPathPrefixes(input), "/"); len(parts) > 1 {
		return filepath.Join(parts[1:]...)
	}
	return input
}

func prepareOptions(opt *Options) (prepared *Options) {
	if opt == nil {
		opt = &Options{Recurse: true}
	}
	if opt.Boundary <= 0 {
		opt.Boundary = hrx.DefaultBoundary
	}
	return opt
}

func pruneName(name, trimPrefix string, pruneDir bool) (pruned string) {
	parts := strings.Split(strings.TrimPrefix(name, "/"), "/")
	if pruneDir && len(parts) > 1 {
		parts = parts[1:]
		pruned = strings.Join(parts, "/")
	}
	if trimPrefix != "" && len(parts) > 1 {
		pruned = strings.TrimPrefix(strings.Join(parts, "/"), trimPrefix)
		pruned = strings.TrimPrefix(pruned, "/")
	}
	return
}

func preparePath(opt *Options, arg string) (pathname string) {
	pathname = arg
	if opt.PruneDir {
		pathname = pruneTopDir(pathname)
	}
	if opt.TrimPrefix != "" {
		pathname = strings.TrimPrefix(pathname, opt.TrimPrefix)
	}
	pathname = trimPathPrefixes(pathname)
	return
}

func isCreateFileErrIgnored(err error) (ignored bool) {
	return errors.Is(err, hrx.ErrInvalidUnicode) || errors.Is(err, ErrNotRegular)
}
