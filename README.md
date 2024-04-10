[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-coreutils/hrx)
[![codecov](https://codecov.io/gh/go-coreutils/hrx/graph/badge.svg?token=174Q05XI65)](https://codecov.io/gh/go-coreutils/hrx)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-coreutils/hrx)](https://goreportcard.com/report/github.com/go-coreutils/hrx)

# hrx - human readable archiver

hrx is a command line interface for working with `.hrx` files, which are human
readable archives, typically used in writing unit tests.

For details on the HRX file format, see: https://github.com/google/hrx

Note: this project has no affiliation with Google and is entirely independent.

# Installation

``` shell
> go install github.com/go-coreutils/hrx/cmd/hrx@latest
```

# Usage

``` shell
$ hrx -h
usage: hrx [global options] <-l|-c|-x> -f <archive> [pathnames...]
       hrx --help
       hrx --list -f existing.hrx
       hrx --create -f new.hrx <path> [paths...]
       hrx --extract -f existing.hrx [pathnames...]
```

# Help

``` shell
$ hrx --help
NAME:
   hrx - human-readable archive (.hrx) utility

USAGE:
   hrx [global options] <-l|-c|-x> -f <archive> [pathnames...]

VERSION:
   v0.5.x

DESCRIPTION:
   hrx is like the tar command except that the archives are human readable.

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


GLOBAL OPTIONS:
   --help         show detailed help
   --usage, -h    show brief usage 
   --verbose, -v  output progress 
   --version, -V  display the version

   OPERATIONS

   --create, -c   create a new archive 
   --extract, -x  extract an existing archive 
   --list, -l     list all archive entries 

   SETTINGS

   --all, -a                      include hidden files and directories 
   --archive value, -f value      specify the archive file
   --boundary value, -b value     specify the entry boundary size 
   --directory value, -o value    specify the output directory
   --keep-empty, -k               include empty files and directories 
   --prune-dir, -P                remove the top directory from all pathnames 
   --recurse, -r                  recurse into directories (default) 
   --trim-prefix value, -T value  trim given prefix from all pathnames
```

# HRX Go Module

For pragmatically interacting with HRX archives, please use the
[go-corelibs/hrx] module which this project
is built with. The functions available within
[this](https://github.com/go-coreutils/hrx) project are intended for use within
[this](https://github.com/go-coreutils/hrx/tree/trunk/cmd/hrx) specific
command-line application.

# Go-CoreUtils

[Go-CoreUtils] is a collection of command line utilities, mostly related to the
development of the [Go-Curses] and [Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreUtils Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[go-corelibs/hrx]: https://github.com/go-corelibs/hrx
[Go-CoreUtils]: https://github.com/go-coreutils
[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
