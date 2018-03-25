package python

/*
#include <stdio.h>
#include "go-python.h"

PyObject*
_gopy_PyFile_FromFile(int fd, char *name, char *mode) {
    FILE *f = fdopen(fd, mode);
    PyObject *py = PyFile_FromFile(f, name, mode, NULL);
    PyFile_SetBufSize(py, 0);
    return py;
}

*/
import "C"

import (
	"os"
)

// FromFile converts a Go file to Python file object.
// Calling close from Python will not close a file descriptor.
func FromFile(f *os.File, mode string) *PyObject {
	p := C._gopy_PyFile_FromFile(C.int(f.Fd()), C.CString(f.Name()), C.CString(mode))
	return togo(p)
}

// SetStdin sets a sys.stdin to a specified file descriptor.
func SetStdin(f *os.File) error {
	return PySys_SetObject("stdin", FromFile(f, "r"))
}

// SetStdout sets a sys.stdout to a specified file descriptor.
func SetStdout(f *os.File) error {
	return PySys_SetObject("stdout", FromFile(f, "w"))
}

// SetStderr sets a sys.stderr to a specified file descriptor.
func SetStderr(f *os.File) error {
	return PySys_SetObject("stderr", FromFile(f, "w"))
}
