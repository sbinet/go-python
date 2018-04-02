package python2

/*
#include <stdio.h>
#include "go-python.h"

PyObject* _gopy2_PyFile_FromFile(int fd, char *name, char *mode) {
    FILE *f = fdopen(fd, mode);
    PyObject *py = PyFile_FromFile(f, name, mode, NULL);
    PyFile_SetBufSize(py, 0);
    return py;
}

*/
import "C"

import (
	"github.com/sbinet/go-python/runtime"
	"os"
	"unsafe"
)

// FromFile converts a Go file to Python file object.
// Calling close from Python will not close a file descriptor.
func (py2Runtime) fromFile(f *os.File, mode string) *Object {
	cname := C.CString(f.Name())
	cmode := C.CString(mode)
	defer func() {
		C.free(unsafe.Pointer(cname))
		C.free(unsafe.Pointer(cmode))
	}()

	p := C._gopy2_PyFile_FromFile(C.int(f.Fd()), cname, cmode)
	return toGo(p)
}

func (py py2Runtime) FromFile(f *os.File, mode string) runtime.Object {
	p := py.fromFile(f, mode)
	return toPtr(p)
}
