package python2

//#include "go-python.h"
import "C"

import (
	"github.com/sbinet/go-python/runtime"
	"unsafe"
)

func (py2Runtime) sysSetObject(name string, v *Object) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return int(C.PySys_SetObject(cname, v.toPy()))
}

func (py py2Runtime) SysSetObject(name string, v runtime.Object) int {
	p := fromPtr(v)
	return py.sysSetObject(name, p)
}

func (py2Runtime) sysGetObject(name string) *Object {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return toGo(C.PySys_GetObject(cname))
}

func (py py2Runtime) SysGetObject(name string) runtime.Object {
	p := py.sysGetObject(name)
	return toPtr(p)
}
