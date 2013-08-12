# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

PYTHON := python2
PYTHON_CFLAGS   := $(shell $(PYTHON) -c "from distutils import sysconfig; print(sysconfig.get_config_var('CFLAGS'))")
PYTHON_CPPFLAGS := $(shell $(PYTHON) -c "from distutils import sysconfig; print(sysconfig.get_python_inc())")
PYTHON_LDFLAGS  := $(shell $(PYTHON) -c "from distutils import sysconfig; print('-L' + sysconfig.get_python_lib(0,1) + ' -lpython' + sysconfig.get_python_version())")

CGO_LDFLAGS := "$(PYTHON_LDFLAGS)"
CGO_CFLAGS  := "-I$(PYTHON_CPPFLAGS) $(PYTHON_CFLAGS)"

# default to gc, but allow caller to override on command line
GO_COMPILER:=$(GC)
ifeq ($(GO_COMPILER),)
	GO_COMPILER:="gc"
endif

build_cwd = CGO_LDFLAGS=$(CGO_LDFLAGS) CGO_CFLAGS=$(CGO_CFLAGS) go build -compiler=$(GO_COMPILER) .
install_cwd = CGO_LDFLAGS=$(CGO_LDFLAGS) CGO_CFLAGS=$(CGO_CFLAGS) go install -compiler=$(GO_COMPILER) .

all: install

install:
	$(install_cwd)
	(cd ./cmd/go-python && $(install_cwd))

build:
	$(build_cwd)
	(cd ./cmd/go-python && $(build_cwd))
