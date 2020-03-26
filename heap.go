package python

// #include "go-python.h"
import "C"

import (
	"errors"
	"unsafe"
)

var (
	gModules = make(map[string]unsafe.Pointer) // register of [modules - method-definitions]
)

// cpyMethodDefs creates a C-array of C.PyMethodDef, with a {0} sentinel
// FIXME(sbinet) convert 'methods' to a *C.PyMethodDef
func cpyMethodDefs(name string, methods []PyMethodDef) *C.PyMethodDef {
	if len(methods) <= 0 {
		return nil
	}

	n := C.size_t(len(methods) + 1)
	cmeths := C._gopy_malloc_PyMethodDefArray(n)
	for i, meth := range methods {
		cmeth := C.PyMethodDef{
			ml_name:  C.CString(meth.Name),
			ml_meth:  C.PyCFunction(meth.Meth),
			ml_flags: C.int(meth.Flags),
			ml_doc:   C.CString(meth.Doc),
		}
		C._gopy_set_PyMethodDef(cmeths, C.int(i), &cmeth)
	}

	gModules[name] = unsafe.Pointer(cmeths)
	return cmeths
}

func Py_InitModule(name string, methods []PyMethodDef) (*PyObject, error) {
	c_mname := C.CString(name)
	defer C.free(unsafe.Pointer(c_mname))

	cmeths := cpyMethodDefs(name, methods)

	obj := togo(C._gopy_InitModule(c_mname, cmeths))
	if obj == nil {
		PyErr_Print()
		return nil, errors.New("python: internal error; module creation failed.")
	}
	return obj, nil
}

func Py_InitModule3(name string, methods []PyMethodDef, doc string) (*PyObject, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cdoc := C.CString(doc)
	defer C.free(unsafe.Pointer(cdoc))

	cmeths := cpyMethodDefs(name, methods)

	obj := togo(C._gopy_InitModule3(cname, cmeths, cdoc))
	if obj == nil {
		PyErr_Print()
		return nil, errors.New("python: internal error; module creation failed.")
	}
	return obj, nil
}

// EOF
