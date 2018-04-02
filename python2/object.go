package python2

/*
#include "go-python.h"

void _gopy2_decref(PyObject *ptr) { Py_DECREF(ptr); }
int _gopy2_PyString_Check(PyObject *ptr) { PyString_Check(ptr); }
int _gopy2_PyString_CheckExact(PyObject *ptr) { PyString_CheckExact(ptr); }
*/
import "C"

import (
	"github.com/sbinet/go-python/runtime"
)

type Object C.PyObject

func (o *Object) Valid() bool {
	return o != nil
}
func (o *Object) IsNone() bool {
	return o.toPy() == C.Py_None
}
func (o *Object) DecRef() {
	if !o.Valid() {
		return
	}
	C._gopy2_decref(o.toPy())
}

func (o *Object) toPy() *C.PyObject {
	return (*C.PyObject)(o)
}

func toGo(obj *C.PyObject) *Object {
	return (*Object)(obj)
}

func fromPtr(obj runtime.Object) *Object {
	if obj == nil {
		return nil
	}
	return obj.(*Object)
}

func toPtr(obj *Object) runtime.Object {
	if obj == nil {
		return nil
	}
	return obj
}

func (obj *Object) HasAttr(name runtime.Object) bool {
	return C.PyObject_HasAttr(obj.toPy(), fromPtr(name).toPy()) != 0
}

func (obj *Object) AsString() string {
	return C.GoString(C.PyString_AsString(obj.toPy()))
}

func (obj *Object) StringCheck() bool {
	return C._gopy2_PyString_Check(obj.toPy()) != 0
}

func (obj *Object) StringCheckExact() bool {
	return C._gopy2_PyString_CheckExact(obj.toPy()) != 0
}
