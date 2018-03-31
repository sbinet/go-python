package python3

//#include "go-python.h"
import "C"

import (
	"github.com/sbinet/go-python"
	"os"
	"unsafe"
)

// FromFile converts a Go file to Python file object.
// Calling close from Python will not close a file descriptor.
func (py3Runtime) fromFile(f *os.File, mode string) *Object {
	cname := C.CString(f.Name())
	cmode := C.CString(mode)
	defer func() {
		C.free(unsafe.Pointer(cname))
		C.free(unsafe.Pointer(cmode))
	}()

	p := C.PyFile_FromFd(C.int(f.Fd()), cname, cmode, 0, nil, nil, nil, 0)
	return toGo(p)
}

func (py py3Runtime) FromFile(f *os.File, mode string) python.ObjectPtr {
	p := py.fromFile(f, mode)
	return toPtr(p)
}
