# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY: all install test

# default to gc, but allow caller to override on command line
GO_COMPILER:=$(GC)
ifeq ($(GO_COMPILER),)
	GO_COMPILER:="gc"
endif

GO_VERBOSE := $(VERBOSE)
ifneq ($(GO_VERBOSE),)
	GO_VERBOSE:= -v -x
endif

install_cwd = go get $(GO_VERBOSE) -compiler=$(GO_COMPILER) .
test_cwd = go test $(GO_VERBOSE) -compiler=$(GO_COMPILER) .

all: install test

install:
	$(install_cwd)
	(cd ./cmd/go-python && $(install_cwd))

test: install
	$(test_cwd)
