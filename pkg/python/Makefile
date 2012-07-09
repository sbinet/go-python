# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

TARG=github.com/sbinet/go-python/pkg/python

PYTHON:=python2
PYTHON_VERSION:=`$(PYTHON) -c "import distutils.sysconfig, sys; sys.stdout.write(''.join(distutils.sysconfig.get_config_vars('VERSION'))+'\n')"`
PYTHON_CFLAGS :=`$(PYTHON) -c "import distutils.sysconfig, sys; sys.stdout.write(''.join(distutils.sysconfig.get_config_vars('CFLAGS'))+'\n')"`
PYTHON_CPPFLAGS:=`$(PYTHON) -c "import distutils.sysconfig, sys; sys.stdout.write(distutils.sysconfig.get_python_inc()+'\n')"`
PYTHON_LDFLAGS:=`$(PYTHON) -c "import distutils.sysconfig, sys; sys.stdout.write('-L'+distutils.sysconfig.get_python_lib(0,1)+' -lpython')"`$(PYTHON_VERSION)
CGO_LDFLAGS:=$(PYTHON_LDFLAGS)
CGO_CFLAGS:=-I$(PYTHON_CPPFLAGS) $(PYTHON_CFLAGS)

GO_COMPILER:=$(GC)
ifeq ($(GO_COMPILER),)
	GO_COMPILER:="gc"
endif


build:
	CGO_LDFLAGS="$(CGO_LDFLAGS)" CGO_CFLAGS="$(CGO_CFLAGS)" go build -compiler=$(GO_COMPILER) .

install:
	CGO_LDFLAGS="$(CGO_LDFLAGS)" CGO_CFLAGS="$(CGO_CFLAGS)" go install -compiler=$(GO_COMPILER) .

test:
	CGO_LDFLAGS="$(CGO_LDFLAGS)" CGO_CFLAGS="$(CGO_CFLAGS)" go test -compiler=$(GO_COMPILER) .

all: install

PHONY: build install
