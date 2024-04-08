#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

# Copyright (c) 2024  The Go-CoreLibs Authors
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

MAKEFILE_KEYS += GOLANG_LIBS
GOLANG_LIBS_MK_FILE := Golang.libs.mk
GOLANG_LIBS_MK_VERSION := v0.1.1
GOLANG_LIBS_MK_DESCRIPTION := go-corelibs presets

#
#: Core Library Package Presets
#

# github.com/go-corelibs/CoreLibs.mk
CL_CORE_LIBS_MK_GO_PACKAGE ?= ${CORELIBS_BASE}/CoreLibs.mk
CL_CORE_LIBS_MK_LOCAL_PATH ?= ${CORELIBS_PATH}/CoreLibs.mk
ifeq (${AUTO_CORELIBS},true)
ifeq (CoreLibs.mk,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^CoreLibs.mk$$'))
GOPKG_KEYS += CL_CORE_LIBS_MK
endif
endif

# github.com/go-corelibs/bcss
CL_BCSS_GO_PACKAGE ?= ${CORELIBS_BASE}/bcss
CL_BCSS_LOCAL_PATH ?= ${CORELIBS_PATH}/bcss
ifeq (${AUTO_CORELIBS},true)
ifeq (bcss,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^bcss$$'))
GOPKG_KEYS += CL_BCSS
endif
endif

# github.com/go-corelibs/chdirs
CL_CHDIRS_GO_PACKAGE ?= ${CORELIBS_BASE}/chdirs
CL_CHDIRS_LOCAL_PATH ?= ${CORELIBS_PATH}/chdirs
ifeq (${AUTO_CORELIBS},true)
ifeq (chdirs,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^chdirs$$'))
GOPKG_KEYS += CL_CHDIRS
endif
endif

# github.com/go-corelibs/cli
CL_CLI_GO_PACKAGE ?= ${CORELIBS_BASE}/cli
CL_CLI_LOCAL_PATH ?= ${CORELIBS_PATH}/cli
ifeq (${AUTO_CORELIBS},true)
ifeq (cli,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^cli$$'))
GOPKG_KEYS += CL_CLI
endif
endif

# github.com/go-corelibs/convert
CL_CONVERT_GO_PACKAGE ?= ${CORELIBS_BASE}/convert
CL_CONVERT_LOCAL_PATH ?= ${CORELIBS_PATH}/convert
ifeq (${AUTO_CORELIBS},true)
ifeq (convert,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^convert$$'))
GOPKG_KEYS += CL_CONVERT
endif
endif

# github.com/go-corelibs/diff
CL_DIFF_GO_PACKAGE ?= ${CORELIBS_BASE}/diff
CL_DIFF_LOCAL_PATH ?= ${CORELIBS_PATH}/diff
ifeq (${AUTO_CORELIBS},true)
ifeq (diff,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^diff$$'))
GOPKG_KEYS += CL_DIFF
endif
endif

# github.com/go-corelibs/env
CL_ENV_GO_PACKAGE ?= ${CORELIBS_BASE}/env
CL_ENV_LOCAL_PATH ?= ${CORELIBS_PATH}/env
ifeq (${AUTO_CORELIBS},true)
ifeq (env,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^env$$'))
GOPKG_KEYS += CL_ENV
endif
endif

# github.com/go-corelibs/filewriter
CL_FILEWRITER_GO_PACKAGE ?= ${CORELIBS_BASE}/filewriter
CL_FILEWRITER_LOCAL_PATH ?= ${CORELIBS_PATH}/filewriter
ifeq (${AUTO_CORELIBS},true)
ifeq (filewriter,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^filewriter$$'))
GOPKG_KEYS += CL_FILEWRITER
endif
endif

# github.com/go-corelibs/fmtstr
CL_FMTSTR_GO_PACKAGE ?= ${CORELIBS_BASE}/fmtstr
CL_FMTSTR_LOCAL_PATH ?= ${CORELIBS_PATH}/fmtstr
ifeq (${AUTO_CORELIBS},true)
ifeq (fmtstr,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^fmtstr$$'))
GOPKG_KEYS += CL_FMTSTR
endif
endif

# github.com/go-corelibs/globs
CL_GLOBS_GO_PACKAGE ?= ${CORELIBS_BASE}/globs
CL_GLOBS_LOCAL_PATH ?= ${CORELIBS_PATH}/globs
ifeq (${AUTO_CORELIBS},true)
ifeq (globs,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^globs$$'))
GOPKG_KEYS += CL_GLOBS
endif
endif

# github.com/go-corelibs/hrx
CL_HRX_GO_PACKAGE ?= ${CORELIBS_BASE}/hrx
CL_HRX_LOCAL_PATH ?= ${CORELIBS_PATH}/hrx
ifeq (${AUTO_CORELIBS},true)
ifeq (hrx,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^hrx$$'))
GOPKG_KEYS += CL_HRX
endif
endif

# github.com/go-corelibs/htmlcss
CL_HTMLCSS_GO_PACKAGE ?= ${CORELIBS_BASE}/htmlcss
CL_HTMLCSS_LOCAL_PATH ?= ${CORELIBS_PATH}/htmlcss
ifeq (${AUTO_CORELIBS},true)
ifeq (htmlcss,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^htmlcss$$'))
GOPKG_KEYS += CL_HTMLCSS
endif
endif

# github.com/go-corelibs/lang
CL_LANG_GO_PACKAGE ?= ${CORELIBS_BASE}/lang
CL_LANG_LOCAL_PATH ?= ${CORELIBS_PATH}/lang
ifeq (${AUTO_CORELIBS},true)
ifeq (lang,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^lang$$'))
GOPKG_KEYS += CL_LANG
endif
endif

# github.com/go-corelibs/maps
CL_MAPS_GO_PACKAGE ?= ${CORELIBS_BASE}/maps
CL_MAPS_LOCAL_PATH ?= ${CORELIBS_PATH}/maps
ifeq (${AUTO_CORELIBS},true)
ifeq (maps,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^maps$$'))
GOPKG_KEYS += CL_MAPS
endif
endif

# github.com/go-corelibs/maths
CL_MATHS_GO_PACKAGE ?= ${CORELIBS_BASE}/maths
CL_MATHS_LOCAL_PATH ?= ${CORELIBS_PATH}/maths
ifeq (${AUTO_CORELIBS},true)
ifeq (maths,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^maths$$'))
GOPKG_KEYS += CL_MATHS
endif
endif

# github.com/go-corelibs/mime
CL_MIME_GO_PACKAGE ?= ${CORELIBS_BASE}/mime
CL_MIME_LOCAL_PATH ?= ${CORELIBS_PATH}/mime
ifeq (${AUTO_CORELIBS},true)
ifeq (mime,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^mime$$'))
GOPKG_KEYS += CL_MIME
endif
endif

# github.com/go-corelibs/mock-stdio
CL_MOCK_STDIO_GO_PACKAGE ?= ${CORELIBS_BASE}/mock-stdio
CL_MOCK_STDIO_LOCAL_PATH ?= ${CORELIBS_PATH}/mock-stdio
ifeq (${AUTO_CORELIBS},true)
ifeq (mock-stdio,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^mock-stdio$$'))
GOPKG_KEYS += CL_MOCK_STDIO
endif
endif

# github.com/go-corelibs/notify
CL_NOTIFY_GO_PACKAGE ?= ${CORELIBS_BASE}/notify
CL_NOTIFY_LOCAL_PATH ?= ${CORELIBS_PATH}/notify
ifeq (${AUTO_CORELIBS},true)
ifeq (notify,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^notify$$'))
GOPKG_KEYS += CL_NOTIFY
endif
endif

# github.com/go-corelibs/path
CL_PATH_GO_PACKAGE ?= ${CORELIBS_BASE}/path
CL_PATH_LOCAL_PATH ?= ${CORELIBS_PATH}/path
ifeq (${AUTO_CORELIBS},true)
ifeq (path,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^path$$'))
GOPKG_KEYS += CL_PATH
endif
endif

# github.com/go-corelibs/regexps
CL_REGEXPS_GO_PACKAGE ?= ${CORELIBS_BASE}/regexps
CL_REGEXPS_LOCAL_PATH ?= ${CORELIBS_PATH}/regexps
ifeq (${AUTO_CORELIBS},true)
ifeq (regexps,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^regexps$$'))
GOPKG_KEYS += CL_REGEXPS
endif
endif

# github.com/go-corelibs/replace
CL_REPLACE_GO_PACKAGE ?= ${CORELIBS_BASE}/replace
CL_REPLACE_LOCAL_PATH ?= ${CORELIBS_PATH}/replace
ifeq (${AUTO_CORELIBS},true)
ifeq (replace,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^replace$$'))
GOPKG_KEYS += CL_REPLACE
endif
endif

# github.com/go-corelibs/run
CL_RUN_GO_PACKAGE ?= ${CORELIBS_BASE}/run
CL_RUN_LOCAL_PATH ?= ${CORELIBS_PATH}/run
ifeq (${AUTO_CORELIBS},true)
ifeq (run,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^run$$'))
GOPKG_KEYS += CL_RUN
endif
endif

# github.com/go-corelibs/scanners
CL_SCANNERS_GO_PACKAGE ?= ${CORELIBS_BASE}/scanners
CL_SCANNERS_LOCAL_PATH ?= ${CORELIBS_PATH}/scanners
ifeq (${AUTO_CORELIBS},true)
ifeq (scanners,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^scanners$$'))
GOPKG_KEYS += CL_SCANNERS
endif
endif

# github.com/go-corelibs/shasum
CL_SHASUM_GO_PACKAGE ?= ${CORELIBS_BASE}/shasum
CL_SHASUM_LOCAL_PATH ?= ${CORELIBS_PATH}/shasum
ifeq (${AUTO_CORELIBS},true)
ifeq (shasum,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^shasum$$'))
GOPKG_KEYS += CL_SHASUM
endif
endif

# github.com/go-corelibs/slices
CL_SLICES_GO_PACKAGE ?= ${CORELIBS_BASE}/slices
CL_SLICES_LOCAL_PATH ?= ${CORELIBS_PATH}/slices
ifeq (${AUTO_CORELIBS},true)
ifeq (slices,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^slices$$'))
GOPKG_KEYS += CL_SLICES
endif
endif

# github.com/go-corelibs/spinner
CL_SPINNER_GO_PACKAGE ?= ${CORELIBS_BASE}/spinner
CL_SPINNER_LOCAL_PATH ?= ${CORELIBS_PATH}/spinner
ifeq (${AUTO_CORELIBS},true)
ifeq (spinner,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^spinner$$'))
GOPKG_KEYS += CL_SPINNER
endif
endif

# github.com/go-corelibs/strcases
CL_STRCASES_GO_PACKAGE ?= ${CORELIBS_BASE}/strcases
CL_STRCASES_LOCAL_PATH ?= ${CORELIBS_PATH}/strcases
ifeq (${AUTO_CORELIBS},true)
ifeq (strcases,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^strcases$$'))
GOPKG_KEYS += CL_STRCASES
endif
endif

# github.com/go-corelibs/strings
CL_STRINGS_GO_PACKAGE ?= ${CORELIBS_BASE}/strings
CL_STRINGS_LOCAL_PATH ?= ${CORELIBS_PATH}/strings
ifeq (${AUTO_CORELIBS},true)
ifeq (strings,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^strings$$'))
GOPKG_KEYS += CL_STRINGS
endif
endif

# github.com/go-corelibs/tdata
CL_TDATA_GO_PACKAGE ?= ${CORELIBS_BASE}/tdata
CL_TDATA_LOCAL_PATH ?= ${CORELIBS_PATH}/tdata
ifeq (${AUTO_CORELIBS},true)
ifeq (tdata,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^tdata$$'))
GOPKG_KEYS += CL_TDATA
endif
endif

# github.com/go-corelibs/templates
CL_TEMPLATES_GO_PACKAGE ?= ${CORELIBS_BASE}/templates
CL_TEMPLATES_LOCAL_PATH ?= ${CORELIBS_PATH}/templates
ifeq (${AUTO_CORELIBS},true)
ifeq (templates,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^templates$$'))
GOPKG_KEYS += CL_TEMPLATES
endif
endif

# github.com/go-corelibs/tmplstr
CL_TMPLSTR_GO_PACKAGE ?= ${CORELIBS_BASE}/tmplstr
CL_TMPLSTR_LOCAL_PATH ?= ${CORELIBS_PATH}/tmplstr
ifeq (${AUTO_CORELIBS},true)
ifeq (tmplstr,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^tmplstr$$'))
GOPKG_KEYS += CL_TMPLSTR
endif
endif

# github.com/go-corelibs/values
CL_VALUES_GO_PACKAGE ?= ${CORELIBS_BASE}/values
CL_VALUES_LOCAL_PATH ?= ${CORELIBS_PATH}/values
ifeq (${AUTO_CORELIBS},true)
ifeq (values,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^values$$'))
GOPKG_KEYS += CL_VALUES
endif
endif

# github.com/go-corelibs/words
CL_WORDS_GO_PACKAGE ?= ${CORELIBS_BASE}/words
CL_WORDS_LOCAL_PATH ?= ${CORELIBS_PATH}/words
ifeq (${AUTO_CORELIBS},true)
ifeq (words,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^words$$'))
GOPKG_KEYS += CL_WORDS
endif
endif

# github.com/go-corelibs/x-text
CL_X_TEXT_GO_PACKAGE ?= ${CORELIBS_BASE}/x-text
CL_X_TEXT_LOCAL_PATH ?= ${CORELIBS_PATH}/x-text
ifeq (${AUTO_CORELIBS},true)
ifeq (x-text,$(shell echo "${FOUND_CORELIBS}" | sed -e 's/ /\n/g' | grep '^x-text$$'))
GOPKG_KEYS += CL_X_TEXT
endif
endif
