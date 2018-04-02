package python2

/*
#include "go-python.h"
*/
import "C"

import (
	"github.com/sbinet/go-python/runtime"
	"unsafe"
)

func (py2Runtime) None() runtime.Object {
	return toGo(C.Py_None)
}
func (py2Runtime) False() runtime.Object {
	return toGo(C.Py_False)
}
func (py2Runtime) True() runtime.Object {
	return toGo(C.Py_True)
}
func (py2Runtime) FromString(v string) runtime.Object {
	cv := C.CString(v)
	defer C.free(unsafe.Pointer(cv))
	return toGo(C.PyString_FromString(cv))
}
func (py2Runtime) FromInt64(v int64) runtime.Object {
	return toGo(C.PyLong_FromLong(C.long(v)))
}
func (py2Runtime) FromFloat64(v float64) runtime.Object {
	return toGo(C.PyFloat_FromDouble(C.double(v)))
}
func (py py2Runtime) FromBool(v bool) runtime.Object {
	if v {
		return py.True()
	}
	return py.False()
}
