package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
//int _gopy_PyTuple_Check(PyObject *o) { return PyTuple_Check(o); }
//int _gopy_PyTuple_CheckExact(PyObject *o) { return PyTuple_CheckExact(o); }
//Py_ssize_t _gopy_PyTuple_GET_SIZE(PyObject *p) { return PyTuple_GET_SIZE(p); }
//void _gopy_PyTuple_SET_ITEM(PyObject *p, Py_ssize_t pos, PyObject *o) { PyTuple_SET_ITEM(p, pos, o); }
//PyObject* _gopy_PyTuple_GET_ITEM(PyObject *p, Py_ssize_t pos) { return PyTuple_GET_ITEM(p, pos); }
import "C"
//import "unsafe"
import "os"

/*
int PyTuple_Check(PyObject *p)
Return true if p is a tuple object or an instance of a subtype of the tuple type.

Changed in version 2.2: Allowed subtypes to be accepted.
*/
func PyTuple_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyTuple_Check(topy(self)))
}

/*
int PyTuple_CheckExact(PyObject *p)
Return true if p is a tuple object, but not an instance of a subtype of the tuple type.

New in version 2.2.
*/
func PyTuple_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyTuple_CheckExact(topy(self)))
}

/*
PyObject* PyTuple_New(Py_ssize_t len)
Return value: New reference.
Return a new tuple object of size len, or NULL on failure.

Changed in version 2.5: This function used an int type for len. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyTuple_New(sz int) *PyObject {
	return togo(C.PyTuple_New(C.Py_ssize_t(sz)))
}

/*
PyObject* PyTuple_Pack(Py_ssize_t n, ...)
Return value: New reference.
Return a new tuple object of size n, or NULL on failure. The tuple values are initialized to the subsequent n C arguments pointing to Python objects. PyTuple_Pack(2, a, b) is equivalent to Py_BuildValue("(OO)", a, b).

New in version 2.4.

Changed in version 2.5: This function used an int type for n. This might require changes in your code for properly supporting 64-bit systems.
*/

func PyTuple_Pack(n int, objs ...*PyObject) *PyObject {
	//FIXME
	panic("not implemented")
}

/*
Py_ssize_t PyTuple_Size(PyObject *p)
Take a pointer to a tuple object, and return the size of that tuple.

Changed in version 2.5: This function returned an int type. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyTuple_Size(self *PyObject) int {
	return int(C.PyTuple_Size(topy(self)))
}

/*
Py_ssize_t PyTuple_GET_SIZE(PyObject *p)
Return the size of the tuple p, which must be non-NULL and point to a tuple; no error checking is performed.

Changed in version 2.5: This function returned an int type. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyTuple_GET_SIZE(self *PyObject) int {
	return int(C._gopy_PyTuple_GET_SIZE(topy(self)))
}

/*
PyObject* PyTuple_GetItem(PyObject *p, Py_ssize_t pos)
Return value: Borrowed reference.
Return the object at position pos in the tuple pointed to by p. If pos is out of bounds, return NULL and sets an IndexError exception.

Changed in version 2.5: This function used an int type for pos. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyTuple_GetItem(self *PyObject, pos int) *PyObject {
	return togo(C.PyTuple_GetItem(topy(self), C.Py_ssize_t(pos)))
}

/*
PyObject* PyTuple_GET_ITEM(PyObject *p, Py_ssize_t pos)
Return value: Borrowed reference.
Like PyTuple_GetItem(), but does no checking of its arguments.

Changed in version 2.5: This function used an int type for pos. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyTuple_GET_ITEM(self *PyObject, pos int) *PyObject {
	return togo(C._gopy_PyTuple_GET_ITEM(topy(self), C.Py_ssize_t(pos)))
}

/*
PyObject* PyTuple_GetSlice(PyObject *p, Py_ssize_t low, Py_ssize_t high)
Return value: New reference.
Take a slice of the tuple pointed to by p from low to high and return it as a new tuple.

Changed in version 2.5: This function used an int type for low and high. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyTuple_GetSlice(self *PyObject, low, high int) *PyObject {
	return togo(C.PyTuple_GetSlice(topy(self), C.Py_ssize_t(low), C.Py_ssize_t(high)))
}

/*
int PyTuple_SetItem(PyObject *p, Py_ssize_t pos, PyObject *o)
Insert a reference to object o at position pos of the tuple pointed to by p. Return 0 on success.

Note This function “steals” a reference to o.
Changed in version 2.5: This function used an int type for pos. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyTuple_SetItem(self *PyObject, pos int, o *PyObject) os.Error {
	return int2err(C.PyTuple_SetItem(topy(self), C.Py_ssize_t(pos), topy(o)))
}

/*
void PyTuple_SET_ITEM(PyObject *p, Py_ssize_t pos, PyObject *o)
Like PyTuple_SetItem(), but does no error checking, and should only be used to fill in brand new tuples.

Note This function “steals” a reference to o.
Changed in version 2.5: This function used an int type for pos. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyTuple_SET_ITEM(self *PyObject, pos int, o *PyObject) {
	py_self := topy(self)
	py_pos  := C.Py_ssize_t(pos)
	py_o    := topy(o)
	C._gopy_PyTuple_SET_ITEM(py_self, py_pos, py_o)
}

/*
int _PyTuple_Resize(PyObject **p, Py_ssize_t newsize)
Can be used to resize a tuple. newsize will be the new length of the tuple. Because tuples are supposed to be immutable, this should only be used if there is only one reference to the object. Do not use this if the tuple may already be known to some other part of the code. The tuple will always grow or shrink at the end. Think of this as destroying the old tuple and creating a new one, only more efficiently. Returns 0 on success. Client code should never assume that the resulting value of *p will be the same as before calling this function. If the object referenced by *p is replaced, the original *p is destroyed. On failure, returns -1 and sets *p to NULL, and raises MemoryError or SystemError.

Changed in version 2.2: Removed unused third parameter, last_is_sticky.

Changed in version 2.5: This function used an int type for newsize. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyTuple_Resize(self *PyObject, newsize int) os.Error {
	py_self := topy(self)
	py_newsz:= C.Py_ssize_t(newsize)
	err := C._PyTuple_Resize(&py_self, py_newsz)
	return int2err(err)
}

/*
int PyTuple_ClearFreeList()
Clear the free list. Return the total number of freed items.

New in version 2.6.
*/
func PyTuple_ClearFreeList() {
	C.PyTuple_ClearFreeList()
}
// EOF
