package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
//int _gopy_PyLong_Check(PyObject *o) { return PyLong_Check(o); }
//int _gopy_PyLong_CheckExact(PyObject *o) { return PyLong_CheckExact(o); }
import "C"
//import "unsafe"

/*
int PyLong_Check(PyObject *p)
Return true if its argument is a PyLongObject or a subtype of PyLongObject.

Changed in version 2.2: Allowed subtypes to be accepted.
*/
func PyLong_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyLong_Check(topy(self)))
}

/*
int PyLong_CheckExact(PyObject *p)
Return true if its argument is a PyLongObject, but not a subtype of PyLongObject.

New in version 2.2.
*/
func PyLong_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyLong_CheckExact(topy(self)))
}

/*
PyObject* PyLong_FromLong(long v)
Return value: New reference.
Return a new PyLongObject object from v, or NULL on failure.
*/
func PyLong_FromLong(v int) *PyObject {
	return togo(C.PyLong_FromLong(C.long(v)))
}

/*
PyObject* PyLong_FromUnsignedLong(unsigned long v)
Return value: New reference.
Return a new PyLongObject object from a C unsigned long, or NULL on failure.
*/
func PyLong_FromUnsignedLong(v uint) *PyObject {
	return togo(C.PyLong_FromUnsignedLong(C.ulong(v)))
}

/*
PyObject* PyLong_FromSsize_t(Py_ssize_t v)
Return value: New reference.
Return a new PyLongObject object from a C Py_ssize_t, or NULL on failure.

New in version 2.6.
*/
func PyLong_FromSsize_t(v int) *PyObject {
	return togo(C.PyLong_FromSsize_t(C.Py_ssize_t(v)))
}

/*
PyObject* PyLong_FromSize_t(size_t v)
Return value: New reference.
Return a new PyLongObject object from a C size_t, or NULL on failure.

New in version 2.6.
*/
func PyLong_FromSize_t(v int) *PyObject {
	return togo(C.PyLong_FromSize_t(C.size_t(v)))
}

/*
PyObject* PyLong_FromLongLong(PY_LONG_LONG v)
Return value: New reference.
Return a new PyLongObject object from a C long long, or NULL on failure.
*/
func PyLong_FromLongLong(v int64) *PyObject {
	return togo(C.PyLong_FromLongLong(C.PY_LONG_LONG(v)))
}

/*
PyObject* PyLong_FromUnsignedLongLong(unsigned PY_LONG_LONG v)
Return value: New reference.
Return a new PyLongObject object from a C unsigned long long, or NULL on failure.
*/
func PyLong_FromUnsignedLongLong(v uint64) *PyObject {
	return togo(C.PyLong_FromUnsignedLongLong(C.ulonglong(v)))
}

/*
PyObject* PyLong_FromDouble(double v)
Return value: New reference.
Return a new PyLongObject object from the integer part of v, or NULL on failure.
*/
func PyLong_FromDouble(v float64) *PyObject {
	return togo(C.PyLong_FromDouble(C.double(v)))
}

/*
PyObject* PyLong_FromString(char *str, char **pend, int base)
Return value: New reference.
Return a new PyLongObject based on the string value in str, which is interpreted according to the radix in base. If pend is non-NULL, *pend will point to the first character in str which follows the representation of the number. If base is 0, the radix will be determined based on the leading characters of str: if str starts with '0x' or '0X', radix 16 will be used; if str starts with '0', radix 8 will be used; otherwise radix 10 will be used. If base is not 0, it must be between 2 and 36, inclusive. Leading spaces are ignored. If there are no digits, ValueError will be raised.
*/
func PyLong_FromString(str string, pend, base int) *PyObject {
	//FIXME
	panic("not implemented")
}

/*
PyObject* PyLong_FromUnicode(Py_UNICODE *u, Py_ssize_t length, int base)
Return value: New reference.
Convert a sequence of Unicode digits to a Python long integer value. The first parameter, u, points to the first character of the Unicode string, length gives the number of characters, and base is the radix for the conversion. The radix must be in the range [2, 36]; if it is out of range, ValueError will be raised.

New in version 1.6.

Changed in version 2.5: This function used an int for length. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyLong_FromUnicode(self *C.Py_UNICODE, length, base int) {
	//FIXME
	panic("not implemented")
}

/*
PyObject* PyLong_FromVoidPtr(void *p)
Return value: New reference.
Create a Python integer or long integer from the pointer p. The pointer value can be retrieved from the resulting value using PyLong_AsVoidPtr().

New in version 1.5.2.

Changed in version 2.5: If the integer is larger than LONG_MAX, a positive long integer is returned.
*/
func PyLong_FromVoidPtr(v interface{}) *PyObject {
	/*
	c_ptr := (*C.char)(unsafe.Pointer(v))
	return togo(C.PyLong_FromVoidPtr(c_ptr))
	 */
	//FIXME
	panic("not implemented")
}

/*
long PyLong_AsLong(PyObject *pylong)
Return a C long representation of the contents of pylong. If pylong is greater than LONG_MAX, an OverflowError is raised and -1 will be returned.

long PyLong_AsLongAndOverflow(PyObject *pylong, int *overflow)
Return a C long representation of the contents of pylong. If pylong is greater than LONG_MAX or less than LONG_MIN, set *overflow to 1 or -1, respectively, and return -1; otherwise, set *overflow to 0. If any other exception occurs (for example a TypeError or MemoryError), then -1 will be returned and *overflow will be 0.

New in version 2.7.
*/
func PyLong_AsLong(self *PyObject) int64 {
	return int64(C.PyLong_AsLong(topy(self)))
}

/*
PY_LONG_LONG PyLong_AsLongLongAndOverflow(PyObject *pylong, int *overflow)
Return a C long long representation of the contents of pylong. If pylong is greater than PY_LLONG_MAX or less than PY_LLONG_MIN, set *overflow to 1 or -1, respectively, and return -1; otherwise, set *overflow to 0. If any other exception occurs (for example a TypeError or MemoryError), then -1 will be returned and *overflow will be 0.

New in version 2.7.
*/
func PyLong_AsLongLongAndOverflow(self *PyObject) (value int64, overflow int) {
	c_overflow := C.int(0)
	value = int64(C.PyLong_AsLongLongAndOverflow(topy(self), &c_overflow))
	overflow = int(c_overflow)
	return
}

/*
Py_ssize_t PyLong_AsSsize_t(PyObject *pylong)
Return a C Py_ssize_t representation of the contents of pylong. If pylong is greater than PY_SSIZE_T_MAX, an OverflowError is raised and -1 will be returned.

New in version 2.6.
*/
func PyLong_AsSsize_t(self *PyObject) int {
	return int(C.PyLong_AsSsize_t(topy(self)))
}

/*
unsigned long PyLong_AsUnsignedLong(PyObject *pylong)
Return a C unsigned long representation of the contents of pylong. If pylong is greater than ULONG_MAX, an OverflowError is raised.
*/
func PyLong_AsUnsignedLong(self *PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLong(topy(self)))
}

/*
PY_LONG_LONG PyLong_AsLongLong(PyObject *pylong)
Return a C long long from a Python long integer. If pylong cannot be represented as a long long, an OverflowError is raised and -1 is returned.

New in version 2.2.
*/
func PyLong_AsLongLong(self *PyObject) int64 {
	return int64(C.PyLong_AsLongLong(topy(self)))
}

/*
unsigned PY_LONG_LONG PyLong_AsUnsignedLongLong(PyObject *pylong)
Return a C unsigned long long from a Python long integer. If pylong cannot be represented as an unsigned long long, an OverflowError is raised and (unsigned long long)-1 is returned.

New in version 2.2.

Changed in version 2.7: A negative pylong now raises OverflowError, not TypeError.
*/
func PyLong_AsUnsignedLongLong(self *PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLongLong(topy(self)))
}


/*
unsigned long PyLong_AsUnsignedLongMask(PyObject *io)
Return a C unsigned long from a Python long integer, without checking for overflow.

New in version 2.3.
*/
func PyLong_AsUnsignedLongMask(self *PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLongMask(topy(self)))
}


/*
unsigned PY_LONG_LONG PyLong_AsUnsignedLongLongMask(PyObject *io)
Return a C unsigned long long from a Python long integer, without checking for overflow.

New in version 2.3.
*/
func PyLong_AsUnsignedLongLongMask(self *PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLongLongMask(topy(self)))
}


/*
double PyLong_AsDouble(PyObject *pylong)
Return a C double representation of the contents of pylong. If pylong cannot be approximately represented as a double, an OverflowError exception is raised and -1.0 will be returned.
*/
func PyLong_AsDouble(self *PyObject) float64 {
	return float64(C.PyLong_AsDouble(topy(self)))
}


/*
void* PyLong_AsVoidPtr(PyObject *pylong)
Convert a Python integer or long integer pylong to a C void pointer. If pylong cannot be converted, an OverflowError will be raised. This is only assured to produce a usable void pointer for values created with PyLong_FromVoidPtr().

New in version 1.5.2.

Changed in version 2.5: For values outside 0..LONG_MAX, both signed and unsigned integers are accepted.
*/
func PyLong_AsVoidPtr(self *PyObject) *C.char {
	//FIXME
	panic("not implemented")
}

// EOF
