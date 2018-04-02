# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY: all install test build

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

build-py2: install
	go build -buildmode=plugin ./python2/plugin/python2.go

build-py3: install
	go build -buildmode=plugin ./python3/plugin/python3.go

build: build-py2 build-py3
	go build ./cmd/go-python
