package python

//#include "go-python.h"
import "C"

type PyTypeObject struct {
	ptr *C.PyTypeObject
}

// int PyType_Check(PyObject *o)
// Return true if the object o is a type object, including instances of types derived from the standard type object. Return false in all other cases.
func PyType_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyType_Check(topy(self)))
}

// int PyType_CheckExact(PyObject *o)
// Return true if the object o is a type object, but not a subtype of the standard type object. Return false in all other cases.
//
// New in version 2.2.
func PyType_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyType_CheckExact(topy(self)))
}

// unsigned int PyType_ClearCache()
// Clear the internal lookup cache. Return the current version tag.
//
// New in version 2.6.
func PyType_ClearCache() uint {
	return uint(C.PyType_ClearCache())
}

// void PyType_Modified(PyTypeObject *type)
// Invalidate the internal lookup cache for the type and all of its subtypes. This function must be called after any manual modification of the attributes or base classes of the type.
//
// New in version 2.6.
func PyType_Modified(self *PyTypeObject) {
	C.PyType_Modified(self.ptr)
}

// int PyType_HasFeature(PyObject *o, int feature)
// Return true if the type object o sets the feature feature. Type features are denoted by single bit flags.
func PyType_HasFeature(self *PyObject, feature int) bool {
	//err := C._gopy_PyType_HasFeature(topy(self), C.int(feature))
	//return int2bool(err)
	//FIXME
	panic("not implemented")
}

// int PyType_IS_GC(PyObject *o)
// Return true if the type object includes support for the cycle detector; this tests the type flag Py_TPFLAGS_HAVE_GC.
//
// New in version 2.0.
func PyType_IS_GC(self *PyObject) bool {
	//return int2bool(C._gopy_PyType_IS_GC(topy(self)))
	//FIXME
	panic("not implemented")
}

// int PyType_IsSubtype(PyTypeObject *a, PyTypeObject *b)
// Return true if a is a subtype of b.
//
// New in version 2.2.
func PyType_IsSubtype(a, b *PyTypeObject) bool {
	return int2bool(C.PyType_IsSubtype(a.ptr, b.ptr))
}

// PyObject* PyType_GenericAlloc(PyTypeObject *type, Py_ssize_t nitems)
// Return value: New reference.
// New in version 2.2.
//
// Changed in version 2.5: This function used an int type for nitems. This might require changes in your code for properly supporting 64-bit systems.
func PyType_GenericAlloc(self *PyTypeObject, nitems int) *PyObject {
	return togo(C.PyType_GenericAlloc(self.ptr, C.Py_ssize_t(nitems)))
}

// PyObject* PyType_GenericNew(PyTypeObject *type, PyObject *args, PyObject *kwds)
// Return value: New reference.
// New in version 2.2.
func PyType_GenericNew(self *PyTypeObject, args, kwds *PyObject) *PyObject {
	return togo(C.PyType_GenericNew(self.ptr, topy(args), topy(kwds)))
}

// int PyType_Ready(PyTypeObject *type)
// Finalize a type object. This should be called on all type objects to finish their initialization. This function is responsible for adding inherited slots from a type’s base class. Return 0 on success, or return -1 and sets an exception on error.
//
// New in version 2.2.
func PyType_Ready(self *PyTypeObject) error {
	return int2err(C.PyType_Ready(self.ptr))
}

//EOF
