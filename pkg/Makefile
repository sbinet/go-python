# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=python

PYTHON:=python2
PYTHON_VERSION:=`$(PYTHON) -c "import distutils.sysconfig, sys; sys.stdout.write(''.join(distutils.sysconfig.get_config_vars('VERSION'))+'\n')"`
PYTHON_CFLAGS :=`$(PYTHON) -c "import distutils.sysconfig, sys; sys.stdout.write(''.join(distutils.sysconfig.get_config_vars('CFLAGS'))+'\n')"`
PYTHON_CPPFLAGS:=`$(PYTHON) -c "import distutils.sysconfig, sys; sys.stdout.write(distutils.sysconfig.get_python_inc()+'\n')"`
PYTHON_LDFLAGS:=`$(PYTHON) -c "import distutils.sysconfig, sys; sys.stdout.write('-L'+distutils.sysconfig.get_python_lib(0,1)+' -lpython')"`$(PYTHON_VERSION)
CGO_LDFLAGS=$(PYTHON_LDFLAGS)
CGO_CFLAGS:=-I$(PYTHON_CPPFLAGS) $(PYTHON_CFLAGS)

CGOFILES=\
	veryhigh.go\
	object.go\
	type.go\
	none.go\
	numeric.go\
	float.go\
	sequence.go\
	dict.go\
	buffer.go\
	python.go


#CLEANFILES+=python_test

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
