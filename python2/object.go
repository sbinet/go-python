package python2

/*
#include "go-python.h"

void _gopy2_decref(PyObject *ptr) { Py_DECREF(ptr); }
*/
import "C"

import (
	"github.com/sbinet/go-python"
	"unsafe"
)

type Object struct {
	ptr *C.PyObject
}

func (o *Object) Valid() bool {
	return o != nil && o.ptr != nil
}
func (o *Object) DecRef() {
	if !o.Valid() {
		return
	}
	C._gopy2_decref(o.ptr)
	o.ptr = nil
}

func (o *Object) toPy() *C.PyObject {
	if o == nil {
		return nil
	}
	return o.ptr
}

func toGo(obj *C.PyObject) *Object {
	if obj == nil {
		return nil
	}
	return &Object{ptr: obj}
}

func fromPtr(obj python.ObjectPtr) *Object {
	if obj == nil {
		return nil
	}
	return obj.(*Object)
}

func toPtr(obj *Object) python.ObjectPtr {
	if obj == nil {
		return nil
	}
	return obj
}

func objectFromPtr(ptr unsafe.Pointer) *Object {
	return toGo((*C.PyObject)(ptr))
}
