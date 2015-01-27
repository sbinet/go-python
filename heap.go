package python

/*
#include <Python.h>

static PyObject* _gopy_InitModule(const char* nm, PyMethodDef* methods) {
	return Py_InitModule(nm, methods);
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

func Py_InitModule(name string, methods ...PyMethodDef) (*PyObject, error) {
	c_mname := C.CString(name)
	defer C.free(unsafe.Pointer(c_mname))
	obj := togo(C._gopy_InitModule(c_mname, nil))
	if obj == nil {
		return nil, errors.New("python: internal error; module creation failed.")
	}
	return obj, nil
}

// EOF
