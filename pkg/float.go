package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
//int _gopy_PyFloat_Check(PyObject *o) { return PyFloat_Check(o); }
//int _gopy_PyFloat_CheckExact(PyObject *o) { return PyFloat_CheckExact(o); }
//double _gopy_PyFloat_AS_DOUBLE(PyObject *pyfloat) { return PyFloat_AS_DOUBLE(pyfloat); }
//int _gopy_PyComplex_Check(PyObject *o) { return PyComplex_Check(o); }
//int _gopy_PyComplex_CheckExact(PyObject *o) { return PyComplex_CheckExact(o); }
import "C"
//import "unsafe"

/*
int PyFloat_Check(PyObject *p)
Return true if its argument is a PyFloatObject or a subtype of PyFloatObject.

Changed in version 2.2: Allowed subtypes to be accepted.
*/
func PyFloat_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyFloat_Check(topy(self)))
}

/*
int PyFloat_CheckExact(PyObject *p)
Return true if its argument is a PyFloatObject, but not a subtype of PyFloatObject.

New in version 2.2.
*/
func PyFloat_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyFloat_CheckExact(topy(self)))
}

/*
PyObject* PyFloat_FromString(PyObject *str, char **pend)
Return value: New reference.
Create a PyFloatObject object based on the string value in str, or NULL on failure. The pend argument is ignored. It remains only for backward compatibility.
*/
func PyFloat_FromString(str *PyObject) *PyObject {
	return togo(C.PyFloat_FromString(topy(str), nil))
}

/*
PyObject* PyFloat_FromDouble(double v)
Return value: New reference.
Create a PyFloatObject object from v, or NULL on failure.
*/
func PyFloat_FromDouble(v float) *PyObject {
	return togo(C.PyFloat_FromDouble(C.double(v)))
}

/*
double PyFloat_AsDouble(PyObject *pyfloat)
Return a C double representation of the contents of pyfloat. If pyfloat is not a Python floating point object but has a __float__() method, this method will first be called to convert pyfloat into a float.
*/
func PyFloat_AsDouble(self *PyObject) float {
	return float(C.PyFloat_AsDouble(topy(self)))
}

/*
double PyFloat_AS_DOUBLE(PyObject *pyfloat)
Return a C double representation of the contents of pyfloat, but without error checking.
*/
func PyFloat_AS_DOUBLE(self *PyObject) float {
	return float(C._gopy_PyFloat_AS_DOUBLE(topy(self)))
}

/*
PyObject* PyFloat_GetInfo(void)
Return a structseq instance which contains information about the precision, minimum and maximum values of a float. It’s a thin wrapper around the header file float.h.

New in version 2.6.
*/
func PyFloat_GetInfo() *PyObject {
	return togo(C.PyFloat_GetInfo())
}

/*
double PyFloat_GetMax()
Return the maximum representable finite float DBL_MAX as C double.

New in version 2.6.
*/
func PyFloat_GetMax() float {
	return float(C.PyFloat_GetMax())
}

/*
double PyFloat_GetMin()
Return the minimum normalized positive float DBL_MIN as C double.

New in version 2.6.
*/
func PyFloat_GetMin() float {
	return float(C.PyFloat_GetMin())
}

/*
int PyFloat_ClearFreeList()
Clear the float free list. Return the number of items that could not be freed.

New in version 2.6.
*/
func PyFloat_ClearFreeList() int {
	return int(C.PyFloat_ClearFreeList())
}

/*
void PyFloat_AsString(char *buf, PyFloatObject *v)
Convert the argument v to a string, using the same rules as str(). The length of buf should be at least 100.

This function is unsafe to call because it writes to a buffer whose length it does not know.

Deprecated since version 2.7: Use PyObject_Str() or PyOS_double_to_string() instead.
*/
func PyFloat_AsString(buf []byte, v *C.PyFloatObject) {
	//FIXME ?
	panic("not implemented")
}

/*
void PyFloat_AsReprString(char *buf, PyFloatObject *v)
Same as PyFloat_AsString, except uses the same rules as repr(). The length of buf should be at least 100.

This function is unsafe to call because it writes to a buffer whose length it does not know.

Deprecated since version 2.7: Use PyObject_Repr() or PyOS_double_to_string() instead.
*/
func PyFloat_AsReprString(buf []byte, v *C.PyFloatObject) {
	//FIXME ?
	panic("not implemented")
}


/////////// complex ///////////

/*
int PyComplex_Check(PyObject *p)
Return true if its argument is a PyComplexObject or a subtype of PyComplexObject.

Changed in version 2.2: Allowed subtypes to be accepted.
*/
func PyComplex_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyComplex_Check(topy(self)))
}

/*
int PyComplex_CheckExact(PyObject *p)
Return true if its argument is a PyComplexObject, but not a subtype of PyComplexObject.

New in version 2.2.
*/
func PyComplex_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyComplex_CheckExact(topy(self)))
}

/*
PyObject* PyComplex_FromCComplex(Py_complex v)
Return value: New reference.
Create a new Python complex number object from a C Py_complex value.
*/
func PyComplex_FromCComplex(v C.Py_complex) *PyObject {
	//FIXME ? use go-complex ?
	return togo(C.PyComplex_FromCComplex(v))
}

/*
PyObject* PyComplex_FromDoubles(double real, double imag)
Return value: New reference.
Return a new PyComplexObject object from real and imag.
*/
func PyComplex_FromDoubles(real, imag float64) *PyObject {
	return togo(C.PyComplex_FromDoubles(C.double(real), C.double(imag)))
}

/*
double PyComplex_RealAsDouble(PyObject *op)
Return the real part of op as a C double.
*/
func PyComplex_RealAsDouble(op *PyObject) float64 {
	return float64(C.PyComplex_RealAsDouble(topy(op)))
}

/*
double PyComplex_ImagAsDouble(PyObject *op)
Return the imaginary part of op as a C double.
*/
func PyComplex_ImagAsDouble(op *PyObject) float64 {
	return float64(C.PyComplex_ImagAsDouble(topy(op)))
}

/*
Py_complex PyComplex_AsCComplex(PyObject *op)
Return the Py_complex value of the complex number op.

Changed in version 2.6: If op is not a Python complex number object but has a __complex__() method, this method will first be called to convert op to a Python complex number object.
*/
func PyComplex_AsCComplex(op *PyObject) C.Py_complex {
	// FIXME ? use go-complex ?
	return C.PyComplex_AsCComplex(topy(op))
}

// EOF
