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
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/go-corelibs/hrx"
	"github.com/go-corelibs/tdata"
)

func TestInternals(t *testing.T) {

	td := tdata.New()

	Convey("prepare source", t, func() {

		a, err := prepareNewSrc("/dev/null")
		So(err, ShouldNotBeNil)
		So(a, ShouldBeNil)

		a, err = prepareNewSrc(td.Join("simple.hrx"))
		So(err, ShouldBeNil)
		So(a, ShouldNotBeNil)
		So(a.Len(), ShouldEqual, 0)

	})

	Convey("prepare existing source", t, func() {

		a, err := prepareExistingSrc("/.not/a/thing.nope")
		So(err, ShouldNotBeNil)
		So(a, ShouldBeNil)

		a, err = prepareExistingSrc("/dev/null")
		So(err, ShouldNotBeNil)
		So(a, ShouldBeNil)

		a, err = prepareExistingSrc(td.Join("simple.hrx"))
		So(err, ShouldBeNil)
		So(a, ShouldNotBeNil)
		So(a.Len(), ShouldEqual, 2)

	})

	Convey("has path prefix", t, func() {

		So(hasPathPrefix(""), ShouldBeFalse)
		So(hasPathPrefix("nope"), ShouldBeFalse)
		So(hasPathPrefix("/"), ShouldBeTrue)
		So(hasPathPrefix("./"), ShouldBeTrue)
		So(hasPathPrefix("../"), ShouldBeTrue)

	})

	Convey("trim path prefix", t, func() {

		So(trimPathPrefix(""), ShouldEqual, "")
		So(trimPathPrefix("nope"), ShouldEqual, "nope")
		So(trimPathPrefix("/"), ShouldEqual, "")
		So(trimPathPrefix("./"), ShouldEqual, "")
		So(trimPathPrefix("../"), ShouldEqual, "")

	})

	Convey("prune top dir", t, func() {

		So(pruneTopDir(""), ShouldEqual, "")
		So(pruneTopDir("nope"), ShouldEqual, "nope")
		So(pruneTopDir("one/two"), ShouldEqual, "two")
		So(pruneTopDir("/one/two"), ShouldEqual, "two")

	})

	Convey("prepare options", t, func() {

		So(prepareOptions(nil), ShouldEqual, &Options{Recurse: true, Boundary: hrx.DefaultBoundary})
		So(prepareOptions(&Options{Boundary: 2}), ShouldEqual, &Options{Boundary: 2})

	})

	Convey("prune name", t, func() {

		So(pruneName("", "", false), ShouldEqual, "")
		So(pruneName("path/name.txt", "", true), ShouldEqual, "name.txt")
		So(pruneName("path/name.txt", "path", false), ShouldEqual, "name.txt")
		So(pruneName("path/name.txt", "name", true), ShouldEqual, "name.txt")

	})

}
