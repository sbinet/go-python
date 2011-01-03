package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
//int _gopy_PyBool_Check(PyObject *o) { return PyBool_Check(o); }
//PyObject *_gopy_pyfalse() { return Py_False; }
//PyObject *_gopy_pytrue() { return Py_True; }
import "C"
//import "unsafe"

/*
int PyBool_Check(PyObject *o)
Return true if o is of type PyBool_Type.

New in version 2.3.
*/
func PyBool_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyBool_Check(topy(self)))
}


/*
The Python False object. This object has no methods. It needs to be treated just like any other object with respect to reference counts.
*/
var Py_False = togo(C._gopy_pyfalse())

/*
PyObject* Py_True
The Python True object. This object has no methods. It needs to be treated just like any other object with respect to reference counts.
*/
var Py_True = togo(C._gopy_pytrue())

/*
Py_RETURN_FALSE
Return Py_False from a function, properly incrementing its reference count.

New in version 2.4.

Py_RETURN_TRUE
Return Py_True from a function, properly incrementing its reference count.

New in version 2.4.
*/

/*
PyObject* PyBool_FromLong(long v)
Return value: New reference.
Return a new reference to Py_True or Py_False depending on the truth value of v.

New in version 2.3.
*/
func PyBool_FromLong(v int) *PyObject {
	return togo(C.PyBool_FromLong(C.long(v)))
}

// EOF
