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
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/dustin/go-humanize"

	"github.com/go-corelibs/hrx"
	"github.com/go-corelibs/path"
)

var (
	gReporting *reporting
)

type reportEntry struct {
	src, pathname, note string
	argv                []interface{}
}

type reporting struct {
	entries []*reportEntry
	sync.RWMutex
}

func reporterFn(src, pathname, note string, argv ...interface{}) {
	safeInitReporting()
	gReporting.Lock()
	defer gReporting.Unlock()
	gReporting.entries = append(gReporting.entries, &reportEntry{src: src, pathname: pathname, note: note, argv: argv})
}

func safeInitReporting() {
	if gReporting == nil {
		gReporting = &reporting{}
	}
}

func safeResetReporting() {
	safeInitReporting()
	gReporting.Lock()
	defer gReporting.Unlock()
	gReporting.entries = make([]*reportEntry, 0)
}

func printSummary(a hrx.Archive, note, value string) {
	maxPathname, maxComment, _ := printSummaryReporting(a, note)
	if path.IsDir(value) {
		note = humanize.Bytes(path.DirSize(value))
	} else {
		note = humanize.Bytes(uint64(path.FileSize(value)))
	}
	if comment, ok := a.GetComment(); ok || maxComment > 0 {
		if comment != "" {
			comment = strings.ReplaceAll(strings.TrimSpace(comment), "\n", "\\n")
		} else {
			comment = "-"
		}
		Notifier.Info("%11s | %-"+strconv.Itoa(maxPathname)+"s | %s\n", note, filepath.Base(value), comment)
		return
	}
	Notifier.Info("%11s | %s\n", note, filepath.Base(value))
}

func printSummaryReporting(a hrx.Archive, note string) (maxPathname, maxComment int, ok bool) {
	if ok = gReporting != nil; ok {
		gReporting.Lock()
		defer gReporting.Unlock()

		for _, entry := range gReporting.entries {
			if size := len(entry.pathname); size > maxPathname {
				maxPathname = size
			}
			if _, comment, present := a.Get(entry.pathname); present {
				if size := len(comment); size > maxComment {
					maxComment = size
				}
			}
		}

		separator := "------------+-" + strings.Repeat("-", maxPathname)
		headingf := []string{"%11s", "%-" + strconv.Itoa(maxPathname) + "s"}
		headingv := []interface{}{strings.ToUpper(note), "PATHNAME"}
		if maxComment > 0 {
			headingf = append(headingf, "%-"+strconv.Itoa(maxComment)+"s")
			headingv = append(headingv, "COMMENT")
			separator += "-+-" + strings.Repeat("-", maxComment)
		}
		format := strings.Join(headingf, " | ") + "\n"

		Notifier.Info(format, headingv...)
		Notifier.Info(separator + "--\n")

		var comment string
		for _, entry := range gReporting.entries {
			if _, comment, _ = a.Get(entry.pathname); comment != "" {
				comment = strings.ReplaceAll(strings.TrimSpace(comment), "\n", "\\n")
			} else if maxComment > 0 {
				comment = "-"
			} else {
				comment = ""
			}
			printSummaryReport(entry, format, comment)
		}

		Notifier.Info(separator + "--\n")
	}
	return
}

func printSummaryReport(re *reportEntry, format, comment string) {
	var desc string
	switch re.note {
	case hrx.OpBoundary:
		return
	case OpListing:
		desc = humanize.Bytes(uint64(len(re.argv[0].(string))))
	case hrx.OpCreated:
		if arg := re.argv[0].(string); path.IsDir(arg) {
			desc = humanize.Bytes(0)
		} else {
			desc = humanize.Bytes(uint64(path.FileSize(arg)))
		}
	case hrx.OpAppended:
		desc = humanize.Bytes(uint64(len(re.argv[0].(string))))
	case hrx.OpExtracted:
		fullname := re.argv[0].(string)
		desc = humanize.Bytes(uint64(path.FileSize(fullname)))
	default:
		desc = re.note
	}
	if comment != "" {
		Notifier.Info(format, desc, re.pathname, comment)
		return
	}
	Notifier.Info(format, desc, re.pathname)
	return
}
