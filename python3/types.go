package python3

/*
#include "go-python.h"
*/
import "C"

import (
	"github.com/sbinet/go-python/runtime"
)

func (py3Runtime) None() runtime.Object {
	return toGo(C.Py_None)
}
func (py3Runtime) False() runtime.Object {
	return toGo(C.Py_False)
}
func (py3Runtime) True() runtime.Object {
	return toGo(C.Py_True)
}
func (py3Runtime) FromString(v string) runtime.Object {
	data, size := stringToWcharT(v)
	size-- // \0
	return toGo(C.PyUnicode_FromWideChar(data, C.Py_ssize_t(size)))
}
func (py3Runtime) FromInt64(v int64) runtime.Object {
	return toGo(C.PyLong_FromLong(C.long(v)))
}
func (py3Runtime) FromFloat64(v float64) runtime.Object {
	return toGo(C.PyFloat_FromDouble(C.double(v)))
}
func (py py3Runtime) FromBool(v bool) runtime.Object {
	if v {
		return py.True()
	}
	return py.False()
}
