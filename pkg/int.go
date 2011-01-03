package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
//int _gopy_PyInt_Check(PyObject *o) { return PyInt_Check(o); }
//int _gopy_PyInt_CheckExact(PyObject *o) { return PyInt_CheckExact(o); }
//long _gopy_PyInt_AS_LONG(PyObject *io) { return PyInt_AS_LONG(io); }
import "C"
import "unsafe"

/*
int PyInt_Check(PyObject *o)
Return true if o is of type PyInt_Type or a subtype of PyInt_Type.

Changed in version 2.2: Allowed subtypes to be accepted.
*/
func PyInt_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyInt_Check(topy(self)))
}

/*
int PyInt_CheckExact(PyObject *o)
Return true if o is of type PyInt_Type, but not a subtype of PyInt_Type.

New in version 2.2.
*/
func PyInt_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyInt_CheckExact(topy(self)))
}

/*
PyObject* PyInt_FromString(char *str, char **pend, int base)
Return value: New reference.
Return a new PyIntObject or PyLongObject based on the string value in str, which is interpreted according to the radix in base. If pend is non-NULL, *pend will point to the first character in str which follows the representation of the number. If base is 0, the radix will be determined based on the leading characters of str: if str starts with '0x' or '0X', radix 16 will be used; if str starts with '0', radix 8 will be used; otherwise radix 10 will be used. If base is not 0, it must be between 2 and 36, inclusive. Leading spaces are ignored. If there are no digits, ValueError will be raised. If the string represents a number too large to be contained within the machine’s long int type and overflow warnings are being suppressed, a PyLongObject will be returned. If overflow warnings are not being suppressed, NULL will be returned in this case.
*/
func PyInt_FromString(str string, pend, base int) *PyObject {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	var c_end *C.char = nil
	if pend > 0 {
		//FIXME: pointer arithmetic...
		//c_end = c_str[pend]
	}
	return togo(C.PyInt_FromString(c_str, &c_end, C.int(base)))
}

/*
PyObject* PyInt_FromLong(long ival)
Return value: New reference.
Create a new integer object with a value of ival.

The current implementation keeps an array of integer objects for all integers between -5 and 256, when you create an int in that range you actually just get back a reference to the existing object. So it should be possible to change the value of 1. I suspect the behaviour of Python in this case is undefined. :-)
*/
func PyInt_FromLong(val int) *PyObject {
	return togo(C.PyInt_FromLong(C.long(val)))
}

/*
PyObject* PyInt_FromSsize_t(Py_ssize_t ival)
Return value: New reference.
Create a new integer object with a value of ival. If the value is larger than LONG_MAX or smaller than LONG_MIN, a long integer object is returned.

New in version 2.5.
*/
func PyInt_FromSsize_t(val int) *PyObject {
	return togo(C.PyInt_FromSsize_t(C.Py_ssize_t(val)))
}

/*
PyObject* PyInt_FromSize_t(size_t ival)
Create a new integer object with a value of ival. If the value exceeds LONG_MAX, a long integer object is returned.

New in version 2.5.
*/
func PyInt_FromSize_t(val int) *PyObject {
	return togo(C.PyInt_FromSize_t(C.size_t(val)))
}

/*
long PyInt_AsLong(PyObject *io)
Will first attempt to cast the object to a PyIntObject, if it is not already one, and then return its value. If there is an error, -1 is returned, and the caller should check PyErr_Occurred() to find out whether there was an error, or whether the value just happened to be -1.
*/
func PyInt_AsLong(self *PyObject) int {
	return int(C.PyInt_AsLong(topy(self)))
}

/*
long PyInt_AS_LONG(PyObject *io)
Return the value of the object io. No error checking is performed.
*/
func PyInt_AS_LONG(self *PyObject) int {
	return int(C._gopy_PyInt_AS_LONG(topy(self)))
}

/*
unsigned long PyInt_AsUnsignedLongMask(PyObject *io)
Will first attempt to cast the object to a PyIntObject or PyLongObject, if it is not already one, and then return its value as unsigned long. This function does not check for overflow.

New in version 2.3.
*/
func PyInt_AsUnsignedLongMask(self *PyObject) uint {
	return uint(C.PyInt_AsUnsignedLongMask(topy(self)))
}

/*
unsigned PY_LONG_LONG PyInt_AsUnsignedLongLongMask(PyObject *io)
Will first attempt to cast the object to a PyIntObject or PyLongObject, if it is not already one, and then return its value as unsigned long long, without checking for overflow.

New in version 2.3.
*/
func PyInt_AsUnsignedLongLongMask(self *PyObject) uint {
	return uint(C.PyInt_AsUnsignedLongLongMask(topy(self)))
}

/*
Py_ssize_t PyInt_AsSsize_t(PyObject *io)
Will first attempt to cast the object to a PyIntObject or PyLongObject, if it is not already one, and then return its value as Py_ssize_t.

New in version 2.5.
*/
func PyInt_AsSsize_t(self *PyObject) int {
	return int(C.PyInt_AsSsize_t(topy(self)))
}

/*
long PyInt_GetMax()
Return the system’s idea of the largest integer it can handle (LONG_MAX, as defined in the system header files).
*/
func PyInt_GetMax() int {
	return int(C.PyInt_GetMax())
}

/*
int PyInt_ClearFreeList()
Clear the integer free list. Return the number of items that could not be freed.

New in version 2.6.
*/
func PyInt_ClearFreeList() {
	C.PyInt_ClearFreeList()
}

// EOF
