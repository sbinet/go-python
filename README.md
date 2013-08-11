go-python
=========

Naive 'go' bindings towards the C-API of CPython

this package provides a ``go`` package named "python" under which most of the ``PyXYZ`` functions and macros of the public C-API of CPython have been exposed.

theoretically, you should be able to just look at:

  http://docs.python.org/c-api/index.html

and know what to type in your ``go`` program.


this package also provides an executable "go-python" which just loads "python" and then call ``python.Py_Main(os.Args)``.
the rational being that under such an executable, ``go`` based extensions for C-Python would be easier to implement (as this usually means calling into ``go`` from ``C`` through some rather convoluted functions hops)


## Install

With `Go 1` and the ``go`` tool, ``cgo`` packages can't pass anymore additional ``CGO_CFLAGS`` from external programs (except `pkg-config`) to the "fake" ``#cgo`` preprocessor directive.
So one has to do instead:

```
$ CGO_CFLAGS="-I/usr/include/python2.7" CGO_LDFLAGS="-lpython2.7 -L/usr/lib" go get github.com/sbinet/go-python
```

or 

```sh
$ mkdir -p $GOPATH/src/github.com/sbinet
$ cd $GOPATH/src/github.com/sbinet
$ git clone http://github.com/sbinet/go-python
$ cd go-python && make
```

or (if you are into one-liners):

```sh
$ CGO_CFLAGS="-I/usr/include/python2.7" \
  CGO_LDFLAGS="-lpython2.7 -L/usr/lib" \
  go get github.com/sbinet/go-python
```

*Note*: you'll need the proper header and `python` development environment. On Debian, you'll need to install the ``python-all-dev`` package

Documentation
-------------

Is available on ``godoc``:

 http://godoc.org/github.com/sbinet/go-python


Example:
--------

```go
package main

import "fmt"
import "github.com/sbinet/go-python"

func init() {
   err := python.Initialize()
   if err != nil {
          panic(err.Error())
   } 
}

func main() {
 	 gostr := "foo" 
	 pystr := python.PyString_FromString(gostr)
	 str := python.PyString_AsString(pystr)
	 fmt.Println("hello [", str, "]")
}
```

```sh
$ go run ./main.go
hello [ foo ]
```

TODO:
-----

 - fix handling of integers (I did a poor job at making sure everything was ok)

 - add CPython unit-tests

 - do not expose ``C.FILE`` pointer and replace it with ``os.File`` in "go-python" API

 - provide an easy way to extend go-python with ``go`` based extensions

 - think about the need (or not) to translate CPython exceptions into go panic/recover mechanism

 - use SWIG to automatically wrap the whole CPython api ?

