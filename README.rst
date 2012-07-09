go-python
=========

Naive 'go' bindings towards the C-API of CPython

this package provides a ``go`` package named "python" under which most of the ``PyXYZ`` functions and macros of the public C-API of CPython have been exposed.

theoretically, you should be able to just look at:

  http://docs.python.org/c-api/index.html

and know what to type in your ``go`` program.


this package also provides an executable "go-python" which just loads "python" and then call ``python.Py_Main(os.Args)``.
the rational being that under such an executable, ``go`` based extensions for C-Python would be easier to implement (as this usually means calling into ``go`` from ``C`` through some rather convoluted functions hops)


Install:
--------

With `Go 1` and the `go` tool, `cgo` packages can't pass anymore additional `CGO_CFLAGS` from external programs (except `pkg-config`) to the "fake" `#cgo` preprocessor directive.
So one has to do instead:

 $ mkdir -p $GOPATH/pkg/github.com/sbinet
 $ cd $GOPATH/pkg/github.com/sbinet
 $ hg clone http://github.com/sbinet/go-python
 $ cd go-python && make



Example:
--------

 $ cat main.go
 package main
 
 import "fmt"
 import "github.com/sbinet/go-python/pkg/python"

 func main() {
  	 gostr := "foo" 
	 pystr := python.PyString_FromString(gostr)
	 str := python.PyString_AsString(pystr)
	 fmt.Println("hello [", str, "]")
 }

$ gorun ./main.go
hello [ foo ]


TODO:
-----

 - fix handling of integers (I did a poor job at making sure everything was ok)

 - add CPython unit-tests

 - do not expose ``C.FILE`` pointer and replace it with ``os.File`` in "go-python" API

 - provide an easy way to extend go-python with ``go`` based extensions

 - think about the need (or not) to translate CPython exceptions into go panic/recover mechanism

 - use SWIG to automatically wrap the whole CPython api ?
