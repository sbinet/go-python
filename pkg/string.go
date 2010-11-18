package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
//int _gopy_PyString_Check(PyObject *o) { return PyString_Check(o); }
import "C"
import "unsafe"

type PyString PyObject

/*
int PyString_Check(PyObject *o)
Return true if the object o is a string object or an instance of a subtype of the string type.

Changed in version 2.2: Allowed subtypes to be accepted.
*/
func PyString_Check(self *PyObject) int {
	return int(C._gopy_PyString_Check(self.ptr))
}
/*
func (self *PyString) Check() int {
	return int(C.PyString_Check(self.topy()))
}
*/

/*
PyObject* PyString_FromString(const char *v)
Return value: New reference.
Return a new string object with a copy of the string v as value on success, and NULL on failure. The parameter v must not be NULL; it will not be checked.
*/
func PyString_FromString(v string) *PyObject {
	cstr := C.CString(v)
	defer C.free(unsafe.Pointer(cstr))
	return togo(C.PyString_FromString(cstr))
}

/*
char* PyString_AsString(PyObject *string)
Return a NUL-terminated representation of the contents of string. The pointer refers to the internal buffer of string, not a copy. The data must not be modified in any way, unless the string was just created using PyString_FromStringAndSize(NULL, size). It must not be deallocated. If string is a Unicode object, this function computes the default encoding of string and operates on that. If string is not a string object at all, PyString_AsString() returns NULL and raises TypeError.
*/
func PyString_AsString(self *PyObject) string {
	c_str := C.PyString_AsString(self.ptr)
	// we dont own c_str...
	//defer C.free(unsafe.Pointer(c_str))
	return C.GoString(c_str)
}
// EOF
