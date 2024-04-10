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
	"errors"
)

var (
	ErrNeedOpMode   = errors.New("missing one of -l, -c or -x")
	ErrMustOpMode   = errors.New("only one of -l, -c or -x are allowed")
	ErrFileNotFound = errors.New("-f is not found or not an archive")
	ErrNeedArchive  = errors.New("missing -f archive")
	ErrDirNotFound  = errors.New("-o is not found or not a directory")
)
