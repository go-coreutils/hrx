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
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/go-corelibs/chdirs"
	"github.com/go-corelibs/hrx"
	"github.com/go-corelibs/mock-stdio"
	"github.com/go-corelibs/notify"
	clPath "github.com/go-corelibs/path"
	"github.com/go-corelibs/tdata"
)

var gOriginalNotifier notify.Notifier

func backupNotifier() {
	if gOriginalNotifier == nil {
		gOriginalNotifier = Notifier
	}
}

func restoreNotifier() {
	if gOriginalNotifier != nil {
		Notifier = gOriginalNotifier
	}
}

func Test(t *testing.T) {

	td := tdata.New()
	_ = chdirs.Push(td.Path())

	Convey("List", t, func() {
		backupNotifier()
		defer restoreNotifier()

		so, se := stdio.NewStdout(), stdio.NewStderr()
		So(so.Capture(), ShouldBeNil)
		So(se.Capture(), ShouldBeNil)
		Notifier = notify.New(notify.Info).Make()
		defer func() {
			so.Restore()
			se.Restore()
		}()

		err := List(td.Join("simple.hrx"))
		So(err, ShouldBeNil)
		sod, sed := string(so.Data()), string(se.Data())
		So(sod, ShouldContainSubstring, "65 B | input.scss\n")
		So(sod, ShouldContainSubstring, "62 B | output.css\n")
		So(sed, ShouldEqual, "")

		So(so.Reset(), ShouldBeNil)
		So(se.Reset(), ShouldBeNil)

		err = List("/dev/null")
		So(err, ShouldNotBeNil)

		So(so.Reset(), ShouldBeNil)
		So(se.Reset(), ShouldBeNil)

		err = List(td.Join("simple.hrx"), "input.scss")
		So(err, ShouldBeNil)
		sod, sed = string(so.Data()), string(se.Data())
		So(sod, ShouldContainSubstring, "65 B | input.scss\n")
		So(sed, ShouldEqual, "")
	})

	Convey("Create", t, func() {

		var a hrx.Archive
		tempdir, err := tdata.NewTempData("", "hrx.create.*")
		So(err, ShouldBeNil)
		So(tempdir, ShouldNotBeNil)
		defer tempdir.Destroy()

		// basic archive creation

		Convey("basic archive creation", func() {

			a, err = Create(
				nil,
				tempdir.Join("created.hrx"),
				"files-in-directories",
			)
			So(err, ShouldBeNil)
			So(a, ShouldNotBeNil)

			//err = a.WriteFile(tempdir.Join("created.hrx"))
			//So(err, ShouldBeNil)

			a, err = Create(nil, "")
			So(err, ShouldEqual, ErrPathRequired)
			So(a, ShouldBeNil)

			a, err = Create(nil, "/dev/null", "nope")
			So(err, ShouldNotBeNil)
			So(a, ShouldBeNil)

			_ = os.WriteFile(tempdir.Join("already-created.hrx"), []byte("<==>\nonly a comment"), 0440)

			a, err = Create(
				nil,
				tempdir.Join("already-created.hrx"),
				"files-in-directories",
			)
			So(err, ShouldNotBeNil)
			So(a, ShouldBeNil)

		})

		Convey("keep empty directories", func() {

			a, err = Create(
				&Options{Recurse: true, KeepEmpty: true},
				tempdir.Join("keep-empty.hrx"),
				"empty-dir",
			)
			So(err, ShouldBeNil)
			So(a, ShouldNotBeNil)
			So(a.List(), ShouldEqual, []string{"empty-dir/"})

			a, err = Create(
				&Options{Recurse: true, KeepEmpty: true, TrimPrefix: "files-in-directories"},
				tempdir.Join("keep-empty.hrx"),
				"empty-dir",
				"files-in-directories/dir/file1",
				"files-in-directories/path/to/file2",
			)
			So(err, ShouldBeNil)
			So(a, ShouldNotBeNil)
			So(a.List(), ShouldEqual, []string{"empty-dir/", "dir/file1", "path/to/file2"})

		})

		Convey("prune top directories", func() {

			a, err = Create(
				&Options{Recurse: true, PruneDir: true},
				tempdir.Join("pruned.hrx"),
				"files-in-directories",
			)
			So(err, ShouldBeNil)
			So(a, ShouldNotBeNil)
			So(a.List(), ShouldEqual, []string{"dir/file1", "path/to/file2"})

			a, err = Create(
				&Options{Recurse: true, PruneDir: true},
				tempdir.Join("pruned.hrx"),
				"files-in-directories/dir/file1",
				"files-in-directories/path/to/file2",
			)
			So(err, ShouldBeNil)
			So(a, ShouldNotBeNil)
			So(a.List(), ShouldEqual, []string{"dir/file1", "path/to/file2"})

			a, err = Create(
				&Options{Recurse: true, PruneDir: true},
				tempdir.Join("broken.hrx"),
				"/dev/null",
				"files-in-directories/path/to/file2",
			)
			So(err, ShouldNotBeNil)
			So(a, ShouldBeNil)

		})

		Convey("trim path prefix", func() {

			a, err = Create(
				&Options{Recurse: true, PruneDir: true, TrimPrefix: "path"},
				tempdir.Join("not-broken.hrx"),
				"files-in-directories",
			)
			So(err, ShouldBeNil)
			So(a, ShouldNotBeNil)
			So(a.List(), ShouldEqual, []string{"dir/file1", "to/file2"})

		})

		Convey("path listing cases", func() {

			a, err = Create(
				nil,
				tempdir.Join("listing.hrx"),
				"empty-dir",
			)
			So(err, ShouldBeNil)
			So(a, ShouldNotBeNil)
			So(a.List(), ShouldEqual, []string(nil))

			_ = os.Mkdir(tempdir.Join("no-perms-dir"), 0000)

			a, err = Create(
				nil,
				tempdir.Join("listing.hrx"),
				tempdir.Join("no-perms-dir"),
			)
			So(err, ShouldNotBeNil)
			So(a, ShouldBeNil)

			a, err = Create(
				&Options{},
				tempdir.Join("listing.hrx"),
				tempdir.Join("no-perms-dir"),
			)
			So(err, ShouldNotBeNil)
			So(a, ShouldBeNil)

			_ = os.Mkdir(tempdir.Join("bin-files-dir"), 0750)
			_ = os.WriteFile(tempdir.Join("bin-files-dir", "binary"), []byte{0xff, 0xfe, 0xfd}, 0640)
			a, err = Create(
				&Options{Recurse: true},
				tempdir.Join("listing.hrx"),
				tempdir.Join("bin-files-dir"),
			)
			So(err, ShouldBeNil) // binary is invalid unicode - malformed input
			So(a, ShouldNotBeNil)

			_ = os.Mkdir(tempdir.Join("bad-files-dir"), 0750)
			_ = os.Symlink("/dev/null", tempdir.Join("bad-files-dir", "dev-null"))
			a, err = Create(
				&Options{Recurse: true},
				tempdir.Join("listing.hrx"),
				tempdir.Join("bad-files-dir"),
			)
			So(err, ShouldBeNil)
			So(a, ShouldNotBeNil)
			So(a.List(), ShouldBeEmpty)

			_ = os.WriteFile(tempdir.Join("bad-files-dir", "write-only"), []byte{}, 0220)
			a, err = Create(
				&Options{Recurse: true},
				tempdir.Join("listing.hrx"),
				tempdir.Join("bad-files-dir"),
			)
			So(err, ShouldNotBeNil)
			So(a, ShouldBeNil)
		})

	})

	Convey("Extract", t, func() {

		tempdir, err := tdata.NewTempData("", "hrx.extract.*")
		So(err, ShouldBeNil)
		So(tempdir, ShouldNotBeNil)
		defer tempdir.Destroy()

		err = Extract(nil, td.Join("simple.hrx"), tempdir.Join("simple.d"))
		So(err, ShouldBeNil)
		var found []string
		found, err = clPath.ListAllFiles(tempdir.Join("simple.d"), true)
		So(err, ShouldBeNil)
		So(found, ShouldEqual, []string{
			tempdir.Join("simple.d", "input.scss"),
			tempdir.Join("simple.d", "output.css"),
		})

		_ = chdirs.Push(tempdir.Path())
		err = Extract(nil, td.Join("simple.hrx"), "")
		_ = chdirs.Pop()
		So(err, ShouldBeNil)
		found, err = clPath.ListAllFiles(tempdir.Join("simple"), true)
		So(err, ShouldBeNil)
		So(found, ShouldEqual, []string{
			tempdir.Join("simple", "input.scss"),
			tempdir.Join("simple", "output.css"),
		})

		err = Extract(nil, "/dev/null", "")
		So(err, ShouldNotBeNil)

		_ = os.Chmod(tempdir.Join("simple.d", "input.scss"), 0440)
		err = Extract(nil, td.Join("simple.hrx"), tempdir.Join("simple.d"))
		So(err, ShouldNotBeNil)

		_ = os.Mkdir(tempdir.Join("read-only.d"), 0550)
		err = Extract(nil, td.Join("simple.hrx"), tempdir.Join("read-only.d", "out.d"))
		So(err, ShouldNotBeNil)

		Convey("prune dir", func() {

			err = Extract(
				&Options{PruneDir: true},
				td.Join("files-in-directories.hrx"),
				tempdir.Join("fid.d"),
			)
			So(err, ShouldBeNil)
			found, err = clPath.ListAllFiles(tempdir.Join("fid.d"), true)
			So(err, ShouldBeNil)
			So(found, ShouldContain, tempdir.Join("fid.d", "file1"))
			So(found, ShouldContain, tempdir.Join("fid.d", "to/file2"))

		})

		Convey("trim prefix present", func() {

			err = Extract(
				&Options{TrimPrefix: "path"},
				td.Join("files-in-directories.hrx"),
				tempdir.Join("fid2.d"),
			)
			So(err, ShouldBeNil)
			found, err = clPath.ListAllFiles(tempdir.Join("fid2.d"), true)
			So(err, ShouldBeNil)
			So(found, ShouldContain, tempdir.Join("fid2.d", "dir/file1"))
			So(found, ShouldContain, tempdir.Join("fid2.d", "to/file2"))

		})

		Convey("trim prefix not present", func() {

			err = Extract(
				&Options{TrimPrefix: "path"},
				td.Join("files-in-directories.hrx"),
				tempdir.Join("fid3.d"),
				"dir/file1",
			)
			So(err, ShouldBeNil)
			found, err = clPath.ListAllFiles(tempdir.Join("fid3.d"), true)
			So(err, ShouldBeNil)
			So(found, ShouldEqual, []string{tempdir.Join("fid3.d", "dir/file1")})

		})

		Convey("trim prefix and mkdir error", func() {

			err = Extract(
				&Options{TrimPrefix: "path"},
				td.Join("directory.hrx"),
				tempdir.Join("fid4.d"),
			)
			So(err, ShouldBeNil)
			found, err = clPath.ListAllFiles(tempdir.Join("fid4.d"), true)
			So(err, ShouldBeNil)
			So(found, ShouldEqual, []string(nil))
			found, err = clPath.ListAllDirs(tempdir.Join("fid4.d"), true)
			So(err, ShouldBeNil)
			So(found, ShouldContain, tempdir.Join("fid4.d", "dir"))
			So(found, ShouldContain, tempdir.Join("fid4.d", "dir", "subdir"))
			So(found, ShouldContain, tempdir.Join("fid4.d", "other"))
			So(found, ShouldContain, tempdir.Join("fid4.d", "other", "subdir"))
			_ = os.Remove(tempdir.Join("fid4.d", "dir", "subdir"))
			_ = os.Chmod(tempdir.Join("fid4.d", "dir"), 0550)
			err = Extract(
				&Options{TrimPrefix: "path"},
				td.Join("directory.hrx"),
				tempdir.Join("fid4.d"),
			)
			So(err, ShouldNotBeNil)

		})

		Convey("mkdir and write file errors", func() {

			err = Extract(
				&Options{Recurse: true},
				td.Join("files-in-directories.hrx"),
				tempdir.Join("fid5.d"),
			)
			So(err, ShouldBeNil)
			_ = os.RemoveAll(tempdir.Join("fid5.d", "path", "to"))
			_ = os.Chmod(tempdir.Join("fid5.d", "path"), 0220)
			err = Extract(
				&Options{Recurse: true, TrimPrefix: "nope"},
				td.Join("files-in-directories.hrx"),
				tempdir.Join("fid5.d"),
			)
			So(err, ShouldNotBeNil)
			_ = os.Chmod(tempdir.Join("fid5.d", "path"), 0770)
			_ = os.Mkdir(tempdir.Join("fid5.d", "path", "to"), 0220)
			err = Extract(
				&Options{Recurse: true, TrimPrefix: "nope"},
				td.Join("files-in-directories.hrx"),
				tempdir.Join("fid5.d"),
			)
			So(err, ShouldNotBeNil)

		})

	})

}
