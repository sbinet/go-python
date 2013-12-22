# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

PYTHON_CONFIG := $(PYTHON_CONFIG)
ifeq ($(PYTHON_CONFIG),)
	PYTHON_CONFIG:="python-config"
endif

PYTHON_CFLAGS := $(shell $(PYTHON_CONFIG) --cflags)
PYTHON_LDFLAGS := $(shell $(PYTHON_CONFIG) --ldflags)

CGO_LDFLAGS := "$(PYTHON_LDFLAGS)"
CGO_CFLAGS  := "-I$(PYTHON_CPPFLAGS) $(PYTHON_CFLAGS)"

# default to gc, but allow caller to override on command line
GO_COMPILER:=$(GC)
ifeq ($(GO_COMPILER),)
	GO_COMPILER:="gc"
endif

GO_VERBOSE := $(VERBOSE)
ifneq ($(GO_VERBOSE),)
	GO_VERBOSE:= -v -x
endif

build_cwd = CGO_LDFLAGS=$(CGO_LDFLAGS) CGO_CFLAGS=$(CGO_CFLAGS) go build $(GO_VERBOSE) -compiler=$(GO_COMPILER) .
install_cwd = CGO_LDFLAGS=$(CGO_LDFLAGS) CGO_CFLAGS=$(CGO_CFLAGS) go install $(GO_VERBOSE) -compiler=$(GO_COMPILER) .

all: install

install:
	$(install_cwd)
	(cd ./cmd/go-python && $(install_cwd))

build:
	$(build_cwd)
	(cd ./cmd/go-python && $(build_cwd))
