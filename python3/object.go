package python3

/*
#include "go-python.h"

void _gopy3_decref(PyObject *ptr) { Py_DECREF(ptr); }
int _gopy3_PyUnicode_Check(PyObject *ptr) { PyUnicode_Check(ptr); }
int _gopy3_PyUnicode_CheckExact(PyObject *ptr) { PyUnicode_CheckExact(ptr); }
*/
import "C"

import (
	"github.com/sbinet/go-python/runtime"
	"unsafe"
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
	C._gopy3_decref(o.toPy())
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
	var size C.Py_ssize_t
	data := C.PyUnicode_AsWideCharString(obj.toPy(), &size)
	defer C.PyMem_Free(unsafe.Pointer(data))
	s, err := wcharTNToString(data, C.size_t(size))
	if err != nil {
		panic(err)
	}
	return s
}

func (obj *Object) StringCheck() bool {
	return C._gopy3_PyUnicode_Check(obj.toPy()) != 0
}

func (obj *Object) StringCheckExact() bool {
	return C._gopy3_PyUnicode_CheckExact(obj.toPy()) != 0
}
