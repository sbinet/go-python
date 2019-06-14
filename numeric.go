package python

//#include "go-python.h"
import "C"

import (
	"unsafe"
)

// int PyInt_Check(PyObject *o)
// Return true if o is of type PyInt_Type or a subtype of PyInt_Type.
//
// Changed in version 2.2: Allowed subtypes to be accepted.
func PyInt_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyInt_Check(topy(self)))
}

// int PyInt_CheckExact(PyObject *o)
// Return true if o is of type PyInt_Type, but not a subtype of PyInt_Type.
//
// New in version 2.2.
func PyInt_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyInt_CheckExact(topy(self)))
}

// PyObject* PyInt_FromString(char *str, char **pend, int base)
// Return value: New reference.
// Return a new PyIntObject or PyLongObject based on the string value in str,
// which is interpreted according to the radix in base. If pend is non-NULL,
// *pend will point to the first character in str which follows the
// representation of the number. If base is 0, the radix will be determined
// based on the leading characters of str: if str starts with '0x' or '0X',
// radix 16 will be used; if str starts with '0', radix 8 will be used; otherwise
// radix 10 will be used. If base is not 0, it must be between 2 and 36,
// inclusive. Leading spaces are ignored. If there are no digits, ValueError will
//  be raised. If the string represents a number too large to be contained within
//  the machine’s long int type and overflow warnings are being suppressed, a
// PyLongObject will be returned. If overflow warnings are not being suppressed,
// NULL will be returned in this case.
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

// PyObject* PyInt_FromLong(long ival)
// Return value: New reference.
// Create a new integer object with a value of ival.
//
// The current implementation keeps an array of integer objects for all integers between -5 and 256, when you create an int in that range you actually just get back a reference to the existing object. So it should be possible to change the value of 1. I suspect the behaviour of Python in this case is undefined. :-)
func PyInt_FromLong(val int) *PyObject {
	return togo(C.PyInt_FromLong(C.long(val)))
}

// PyObject* PyInt_FromSsize_t(Py_ssize_t ival)
// Return value: New reference.
// Create a new integer object with a value of ival. If the value is larger than LONG_MAX or smaller than LONG_MIN, a long integer object is returned.
//
// New in version 2.5.
func PyInt_FromSsize_t(val int) *PyObject {
	return togo(C.PyInt_FromSsize_t(C.Py_ssize_t(val)))
}

// PyObject* PyInt_FromSize_t(size_t ival)
// Create a new integer object with a value of ival. If the value exceeds LONG_MAX, a long integer object is returned.
//
// New in version 2.5.
func PyInt_FromSize_t(val int) *PyObject {
	return togo(C.PyInt_FromSize_t(C.size_t(val)))
}

// long PyInt_AsLong(PyObject *io)
// Will first attempt to cast the object to a PyIntObject, if it is not already one, and then return its value. If there is an error, -1 is returned, and the caller should check PyErr_Occurred() to find out whether there was an error, or whether the value just happened to be -1.
func PyInt_AsLong(self *PyObject) int {
	return int(C.PyInt_AsLong(topy(self)))
}

// long PyInt_AS_LONG(PyObject *io)
// Return the value of the object io. No error checking is performed.
func PyInt_AS_LONG(self *PyObject) int {
	return int(C._gopy_PyInt_AS_LONG(topy(self)))
}

// unsigned long PyInt_AsUnsignedLongMask(PyObject *io)
// Will first attempt to cast the object to a PyIntObject or PyLongObject, if it is not already one, and then return its value as unsigned long. This function does not check for overflow.
//
// New in version 2.3.
func PyInt_AsUnsignedLongMask(self *PyObject) uint {
	return uint(C.PyInt_AsUnsignedLongMask(topy(self)))
}

// unsigned PY_LONG_LONG PyInt_AsUnsignedLongLongMask(PyObject *io)
// Will first attempt to cast the object to a PyIntObject or PyLongObject, if it is not already one, and then return its value as unsigned long long, without checking for overflow.
//
// New in version 2.3.
func PyInt_AsUnsignedLongLongMask(self *PyObject) uint {
	return uint(C.PyInt_AsUnsignedLongLongMask(topy(self)))
}

// Py_ssize_t PyInt_AsSsize_t(PyObject *io)
// Will first attempt to cast the object to a PyIntObject or PyLongObject, if it is not already one, and then return its value as Py_ssize_t.
//
// New in version 2.5.
func PyInt_AsSsize_t(self *PyObject) int {
	return int(C.PyInt_AsSsize_t(topy(self)))
}

// long PyInt_GetMax()
// Return the system’s idea of the largest integer it can handle (LONG_MAX, as defined in the system header files).
func PyInt_GetMax() int {
	return int(C.PyInt_GetMax())
}

// int PyInt_ClearFreeList()
// Clear the integer free list. Return the number of items that could not be freed.
//
// New in version 2.6.
func PyInt_ClearFreeList() {
	C.PyInt_ClearFreeList()
}

////////// long //////////

// int PyLong_Check(PyObject *p)
// Return true if its argument is a PyLongObject or a subtype of PyLongObject.
//
// Changed in version 2.2: Allowed subtypes to be accepted.
func PyLong_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyLong_Check(topy(self)))
}

// int PyLong_CheckExact(PyObject *p)
// Return true if its argument is a PyLongObject, but not a subtype of PyLongObject.
//
// New in version 2.2.
func PyLong_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyLong_CheckExact(topy(self)))
}

// PyObject* PyLong_FromLong(long v)
// Return value: New reference.
// Return a new PyLongObject object from v, or NULL on failure.
func PyLong_FromLong(v int) *PyObject {
	return togo(C.PyLong_FromLong(C.long(v)))
}

// PyObject* PyLong_FromUnsignedLong(unsigned long v)
// Return value: New reference.
// Return a new PyLongObject object from a C unsigned long, or NULL on failure.
func PyLong_FromUnsignedLong(v uint) *PyObject {
	return togo(C.PyLong_FromUnsignedLong(C.ulong(v)))
}

// PyObject* PyLong_FromSsize_t(Py_ssize_t v)
// Return value: New reference.
// Return a new PyLongObject object from a C Py_ssize_t, or NULL on failure.
//
// New in version 2.6.
func PyLong_FromSsize_t(v int) *PyObject {
	return togo(C.PyLong_FromSsize_t(C.Py_ssize_t(v)))
}

// PyObject* PyLong_FromSize_t(size_t v)
// Return value: New reference.
// Return a new PyLongObject object from a C size_t, or NULL on failure.
//
// New in version 2.6.
func PyLong_FromSize_t(v int) *PyObject {
	return togo(C.PyLong_FromSize_t(C.size_t(v)))
}

// PyObject* PyLong_FromLongLong(PY_LONG_LONG v)
// Return value: New reference.
// Return a new PyLongObject object from a C long long, or NULL on failure.
func PyLong_FromLongLong(v int64) *PyObject {
	return togo(C.PyLong_FromLongLong(C.PY_LONG_LONG(v)))
}

// PyObject* PyLong_FromUnsignedLongLong(unsigned PY_LONG_LONG v)
// Return value: New reference.
// Return a new PyLongObject object from a C unsigned long long, or NULL on failure.
func PyLong_FromUnsignedLongLong(v uint64) *PyObject {
	return togo(C.PyLong_FromUnsignedLongLong(C.ulonglong(v)))
}

// PyObject* PyLong_FromDouble(double v)
// Return value: New reference.
// Return a new PyLongObject object from the integer part of v, or NULL on failure.
func PyLong_FromDouble(v float64) *PyObject {
	return togo(C.PyLong_FromDouble(C.double(v)))
}

// PyObject* PyLong_FromString(char *str, char **pend, int base)
// Return value: New reference.
// Return a new PyLongObject based on the string value in str, which is interpreted according to the radix in base. If pend is non-NULL, *pend will point to the first character in str which follows the representation of the number. If base is 0, the radix will be determined based on the leading characters of str: if str starts with '0x' or '0X', radix 16 will be used; if str starts with '0', radix 8 will be used; otherwise radix 10 will be used. If base is not 0, it must be between 2 and 36, inclusive. Leading spaces are ignored. If there are no digits, ValueError will be raised.
func PyLong_FromString(str string, pend, base int) *PyObject {
	//FIXME
	panic("not implemented")
}

// PyObject* PyLong_FromUnicode(Py_UNICODE *u, Py_ssize_t length, int base)
// Return value: New reference.
// Convert a sequence of Unicode digits to a Python long integer value. The first parameter, u, points to the first character of the Unicode string, length gives the number of characters, and base is the radix for the conversion. The radix must be in the range [2, 36]; if it is out of range, ValueError will be raised.
//
// New in version 1.6.
//
// Changed in version 2.5: This function used an int for length. This might require changes in your code for properly supporting 64-bit systems.
func PyLong_FromUnicode(self *C.Py_UNICODE, length, base int) {
	//FIXME
	panic("not implemented")
}

// PyObject* PyLong_FromVoidPtr(void *p)
// Return value: New reference.
// Create a Python integer or long integer from the pointer p. The pointer value can be retrieved from the resulting value using PyLong_AsVoidPtr().
//
// New in version 1.5.2.
//
// Changed in version 2.5: If the integer is larger than LONG_MAX, a positive long integer is returned.
func PyLong_FromVoidPtr(v interface{}) *PyObject {
	/*
		c_ptr := (*C.char)(unsafe.Pointer(v))
		return togo(C.PyLong_FromVoidPtr(c_ptr))
	*/
	//FIXME
	panic("not implemented")
}

// long PyLong_AsLong(PyObject *pylong)
// Return a C long representation of the contents of pylong. If pylong is greater than LONG_MAX, an OverflowError is raised and -1 will be returned.
//
// long PyLong_AsLongAndOverflow(PyObject *pylong, int *overflow)
// Return a C long representation of the contents of pylong. If pylong is greater than LONG_MAX or less than LONG_MIN, set *overflow to 1 or -1, respectively, and return -1; otherwise, set *overflow to 0. If any other exception occurs (for example a TypeError or MemoryError), then -1 will be returned and *overflow will be 0.
//
// New in version 2.7.
func PyLong_AsLong(self *PyObject) int64 {
	return int64(C.PyLong_AsLong(topy(self)))
}

// PY_LONG_LONG PyLong_AsLongLongAndOverflow(PyObject *pylong, int *overflow)
// Return a C long long representation of the contents of pylong. If pylong is greater than PY_LLONG_MAX or less than PY_LLONG_MIN, set *overflow to 1 or -1, respectively, and return -1; otherwise, set *overflow to 0. If any other exception occurs (for example a TypeError or MemoryError), then -1 will be returned and *overflow will be 0.
//
// New in version 2.7.
func PyLong_AsLongLongAndOverflow(self *PyObject) (value int64, overflow int) {
	c_overflow := C.int(0)
	value = int64(C.PyLong_AsLongLongAndOverflow(topy(self), &c_overflow))
	overflow = int(c_overflow)
	return
}

// Py_ssize_t PyLong_AsSsize_t(PyObject *pylong)
// Return a C Py_ssize_t representation of the contents of pylong. If pylong is greater than PY_SSIZE_T_MAX, an OverflowError is raised and -1 will be returned.
//
// New in version 2.6.
func PyLong_AsSsize_t(self *PyObject) int {
	return int(C.PyLong_AsSsize_t(topy(self)))
}

// unsigned long PyLong_AsUnsignedLong(PyObject *pylong)
// Return a C unsigned long representation of the contents of pylong. If pylong is greater than ULONG_MAX, an OverflowError is raised.
func PyLong_AsUnsignedLong(self *PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLong(topy(self)))
}

// PY_LONG_LONG PyLong_AsLongLong(PyObject *pylong)
// Return a C long long from a Python long integer. If pylong cannot be represented as a long long, an OverflowError is raised and -1 is returned.
//
// New in version 2.2.
func PyLong_AsLongLong(self *PyObject) int64 {
	return int64(C.PyLong_AsLongLong(topy(self)))
}

// unsigned PY_LONG_LONG PyLong_AsUnsignedLongLong(PyObject *pylong)
// Return a C unsigned long long from a Python long integer. If pylong cannot be represented as an unsigned long long, an OverflowError is raised and (unsigned long long)-1 is returned.
//
// New in version 2.2.
//
// Changed in version 2.7: A negative pylong now raises OverflowError, not TypeError.
func PyLong_AsUnsignedLongLong(self *PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLongLong(topy(self)))
}

// unsigned long PyLong_AsUnsignedLongMask(PyObject *io)
// Return a C unsigned long from a Python long integer, without checking for overflow.
//
// New in version 2.3.
func PyLong_AsUnsignedLongMask(self *PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLongMask(topy(self)))
}

// unsigned PY_LONG_LONG PyLong_AsUnsignedLongLongMask(PyObject *io)
// Return a C unsigned long long from a Python long integer, without checking for overflow.
//
// New in version 2.3.
func PyLong_AsUnsignedLongLongMask(self *PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLongLongMask(topy(self)))
}

// double PyLong_AsDouble(PyObject *pylong)
// Return a C double representation of the contents of pylong. If pylong cannot be approximately represented as a double, an OverflowError exception is raised and -1.0 will be returned.
func PyLong_AsDouble(self *PyObject) float64 {
	return float64(C.PyLong_AsDouble(topy(self)))
}

// void* PyLong_AsVoidPtr(PyObject *pylong)
// Convert a Python integer or long integer pylong to a C void pointer. If pylong cannot be converted, an OverflowError will be raised. This is only assured to produce a usable void pointer for values created with PyLong_FromVoidPtr().
//
// New in version 1.5.2.
//
// Changed in version 2.5: For values outside 0..LONG_MAX, both signed and unsigned integers are accepted.
func PyLong_AsVoidPtr(self *PyObject) *C.char {
	//FIXME
	panic("not implemented")
}

//////////// bool ////////////

// int PyBool_Check(PyObject *o)
// Return true if o is of type PyBool_Type.
//
// New in version 2.3.
func PyBool_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyBool_Check(topy(self)))
}

// The Python False object. This object has no methods.
// It needs to be treated just like any other object with respect to
// reference counts.
var Py_False = &PyObject{ptr: C._gopy_pyfalse()}

// PyObject* Py_True
// The Python True object. This object has no methods.
// It needs to be treated just like any other object with respect to
// reference counts.
var Py_True = &PyObject{ptr: C._gopy_pytrue()}

/*
Py_RETURN_FALSE
Return Py_False from a function, properly incrementing its reference count.

New in version 2.4.

Py_RETURN_TRUE
Return Py_True from a function, properly incrementing its reference count.

New in version 2.4.
*/

// PyObject* PyBool_FromLong(long v)
// Return value: New reference.
// Return a new reference to Py_True or Py_False depending on the truth value of v.
//
// New in version 2.3.
func PyBool_FromLong(v int) *PyObject {
	return togo(C.PyBool_FromLong(C.long(v)))
}

////////// float ////////////

// int PyFloat_Check(PyObject *p)
// Return true if its argument is a PyFloatObject or a subtype of PyFloatObject.
//
// Changed in version 2.2: Allowed subtypes to be accepted.
func PyFloat_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyFloat_Check(topy(self)))
}

// int PyFloat_CheckExact(PyObject *p)
// Return true if its argument is a PyFloatObject, but not a subtype of PyFloatObject.
//
// New in version 2.2.
func PyFloat_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyFloat_CheckExact(topy(self)))
}

// PyObject* PyFloat_FromString(PyObject *str, char **pend)
// Return value: New reference.
// Create a PyFloatObject object based on the string value in str, or NULL on failure. The pend argument is ignored. It remains only for backward compatibility.
func PyFloat_FromString(str *PyObject) *PyObject {
	return togo(C.PyFloat_FromString(topy(str), nil))
}

// PyObject* PyFloat_FromDouble(double v)
// Return value: New reference.
// Create a PyFloatObject object from v, or NULL on failure.
func PyFloat_FromDouble(v float64) *PyObject {
	return togo(C.PyFloat_FromDouble(C.double(v)))
}

// double PyFloat_AsDouble(PyObject *pyfloat)
// Return a C double representation of the contents of pyfloat. If pyfloat is not a Python floating point object but has a __float__() method, this method will first be called to convert pyfloat into a float.
func PyFloat_AsDouble(self *PyObject) float64 {
	return float64(C.PyFloat_AsDouble(topy(self)))
}

// double PyFloat_AS_DOUBLE(PyObject *pyfloat)
// Return a C double representation of the contents of pyfloat, but without error checking.
func PyFloat_AS_DOUBLE(self *PyObject) float64 {
	return float64(C._gopy_PyFloat_AS_DOUBLE(topy(self)))
}

// PyObject* PyFloat_GetInfo(void)
// Return a structseq instance which contains information about the precision, minimum and maximum values of a float. It’s a thin wrapper around the header file float.h.
//
// New in version 2.6.
func PyFloat_GetInfo() *PyObject {
	return togo(C.PyFloat_GetInfo())
}

// double PyFloat_GetMax()
// Return the maximum representable finite float DBL_MAX as C double.
//
// New in version 2.6.
func PyFloat_GetMax() float64 {
	return float64(C.PyFloat_GetMax())
}

// double PyFloat_GetMin()
// Return the minimum normalized positive float DBL_MIN as C double.
//
// New in version 2.6.
func PyFloat_GetMin() float64 {
	return float64(C.PyFloat_GetMin())
}

// int PyFloat_ClearFreeList()
// Clear the float free list. Return the number of items that could not be freed.
//
// New in version 2.6.
func PyFloat_ClearFreeList() int {
	return int(C.PyFloat_ClearFreeList())
}

// void PyFloat_AsString(char *buf, PyFloatObject *v)
// Convert the argument v to a string, using the same rules as str(). The length of buf should be at least 100.
//
// This function is unsafe to call because it writes to a buffer whose length it does not know.
//
// Deprecated since version 2.7: Use PyObject_Str() or PyOS_double_to_string() instead.
func PyFloat_AsString(buf []byte, v *C.PyFloatObject) {
	//FIXME ?
	panic("not implemented")
}

// void PyFloat_AsReprString(char *buf, PyFloatObject *v)
// Same as PyFloat_AsString, except uses the same rules as repr(). The length of buf should be at least 100.
//
// This function is unsafe to call because it writes to a buffer whose length it does not know.
//
// Deprecated since version 2.7: Use PyObject_Repr() or PyOS_double_to_string() instead.
func PyFloat_AsReprString(buf []byte, v *C.PyFloatObject) {
	//FIXME ?
	panic("not implemented")
}

/////////// complex ///////////

// int PyComplex_Check(PyObject *p)
// Return true if its argument is a PyComplexObject or a subtype of PyComplexObject.
//
// Changed in version 2.2: Allowed subtypes to be accepted.
func PyComplex_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyComplex_Check(topy(self)))
}

// int PyComplex_CheckExact(PyObject *p)
// Return true if its argument is a PyComplexObject, but not a subtype of PyComplexObject.
//
// New in version 2.2.
func PyComplex_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyComplex_CheckExact(topy(self)))
}

// PyObject* PyComplex_FromCComplex(Py_complex v)
// Return value: New reference.
// Create a new Python complex number object from a C Py_complex value.
func PyComplex_FromCComplex(v C.Py_complex) *PyObject {
	//FIXME ? use go-complex ?
	return togo(C.PyComplex_FromCComplex(v))
}

// PyObject* PyComplex_FromDoubles(double real, double imag)
// Return value: New reference.
// Return a new PyComplexObject object from real and imag.
func PyComplex_FromDoubles(real, imag float64) *PyObject {
	return togo(C.PyComplex_FromDoubles(C.double(real), C.double(imag)))
}

// double PyComplex_RealAsDouble(PyObject *op)
// Return the real part of op as a C double.
func PyComplex_RealAsDouble(op *PyObject) float64 {
	return float64(C.PyComplex_RealAsDouble(topy(op)))
}

// double PyComplex_ImagAsDouble(PyObject *op)
// Return the imaginary part of op as a C double.
func PyComplex_ImagAsDouble(op *PyObject) float64 {
	return float64(C.PyComplex_ImagAsDouble(topy(op)))
}

// Py_complex PyComplex_AsCComplex(PyObject *op)
// Return the Py_complex value of the complex number op.
//
// Changed in version 2.6: If op is not a Python complex number object but has a __complex__() method, this method will first be called to convert op to a Python complex number object.
func PyComplex_AsCComplex(op *PyObject) C.Py_complex {
	// FIXME ? use go-complex ?
	return C.PyComplex_AsCComplex(topy(op))
}

// EOF
