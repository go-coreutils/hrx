#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

# Copyright (c) 2023  The Go-Enjin Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

MAKEFILE_KEYS += GOLANG_LIB
GOLANG_LIB_MK_FILE := Golang.lib.mk
GOLANG_LIB_MK_VERSION := v0.3.0
GOLANG_LIB_MK_DESCRIPTION := go-corelibs support

#
#: Core Library Packages
#

AUTO_CORELIBS ?= false

CORELIBS_BASE ?= github.com/go-corelibs
CORELIBS_PATH ?= ../../go-corelibs

define _grep_corelibs
$(shell find * \
	-name "*.go" -exec grep '"github.com/go-corelibs/' \{\} \; \
	| perl -pe 's!^[^"]*!!;s![\s"]!!g;s!github\.com/go-corelibs/!!;s!$$!\n!;' \
	| sort -u )
endef

FOUND_CORELIBS := $(call _grep_corelibs)

include Golang.libs.mk

corelibs:
	@if [ -n "${FOUND_CORELIBS}" ]; then \
		for CL in ${FOUND_CORELIBS}; do \
			echo "# github.com/go-corelibs/$${CL}"; \
		done; \
	else \
		echo "# no go-corelibs detected"; \
	fi
