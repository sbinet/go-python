package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
//int _gopy_PyList_Check(PyObject *o) { return PyList_Check(o); }
//int _gopy_PyList_CheckExact(PyObject *o) { return PyList_CheckExact(o); }
//Py_ssize_t _gopy_PyList_GET_SIZE(PyObject *o) { return PyList_GET_SIZE(o); }
//PyObject* _gopy_PyList_GET_ITEM(PyObject *list, Py_ssize_t i) { return PyList_GET_ITEM(list, i); }
//void _gopy_PyList_SET_ITEM(PyObject *list, Py_ssize_t i, PyObject *o) { PyList_SET_ITEM(list, i, o); }
import "C"
//import "unsafe"
import "os"

/*
int PyList_Check(PyObject *p)
Return true if p is a list object or an instance of a subtype of the list type.

Changed in version 2.2: Allowed subtypes to be accepted.
*/
func PyList_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyList_Check(topy(self)))
}

/*
int PyList_CheckExact(PyObject *p)
Return true if p is a list object, but not an instance of a subtype of the list type.

New in version 2.2.
*/
func PyList_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyList_CheckExact(topy(self)))
}

/*
PyObject* PyList_New(Py_ssize_t len)
Return value: New reference.
Return a new list of length len on success, or NULL on failure.

Note If length is greater than zero, the returned list object’s items are set to NULL. Thus you cannot use abstract API functions such as PySequence_SetItem() or expose the object to Python code before setting all items to a real object with PyList_SetItem().
Changed in version 2.5: This function used an int for size. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyList_New(sz int) *PyObject {
	return togo(C.PyList_New(C.Py_ssize_t(sz)))
}

/*
Py_ssize_t PyList_Size(PyObject *list)
Return the length of the list object in list; this is equivalent to len(list) on a list object.

Changed in version 2.5: This function returned an int. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyList_Size(self *PyObject) int {
	return int(C.PyList_Size(topy(self)))
}

/*
Py_ssize_t PyList_GET_SIZE(PyObject *list)
Macro form of PyList_Size() without error checking.

Changed in version 2.5: This macro returned an int. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyList_GET_SIZE(self *PyObject) int {
	return int(C._gopy_PyList_GET_SIZE(topy(self)))
}

/*
PyObject* PyList_GetItem(PyObject *list, Py_ssize_t index)
Return value: Borrowed reference.
Return the object at position pos in the list pointed to by p. The position must be positive, indexing from the end of the list is not supported. If pos is out of bounds, return NULL and set an IndexError exception.

Changed in version 2.5: This function used an int for index. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyList_GetItem(self *PyObject, index int) *PyObject {
	return togo(C.PyList_GetItem(topy(self), C.Py_ssize_t(index)))
}

/*
PyObject* PyList_GET_ITEM(PyObject *list, Py_ssize_t i)
Return value: Borrowed reference.
Macro form of PyList_GetItem() without error checking.

Changed in version 2.5: This macro used an int for i. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyList_GET_ITEM(self *PyObject, index int) *PyObject {
	return togo(C._gopy_PyList_GET_ITEM(topy(self), C.Py_ssize_t(index)))
}

/*
int PyList_SetItem(PyObject *list, Py_ssize_t index, PyObject *item)
Set the item at index index in list to item. Return 0 on success or -1 on failure.

Note This function “steals” a reference to item and discards a reference to an item already in the list at the affected position.
Changed in version 2.5: This function used an int for index. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyList_SetItem(self *PyObject, index int, item *PyObject) os.Error {
	err := C.PyList_SetItem(topy(self), C.Py_ssize_t(index), topy(item))
	return int2err(err)
}

/*
void PyList_SET_ITEM(PyObject *list, Py_ssize_t i, PyObject *o)
Macro form of PyList_SetItem() without error checking. This is normally only used to fill in new lists where there is no previous content.

Note This macro “steals” a reference to item, and, unlike PyList_SetItem(), does not discard a reference to any item that it being replaced; any reference in list at position i will be leaked.
Changed in version 2.5: This macro used an int for i. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyList_SET_ITEM(self *PyObject, index int, o *PyObject) {
	C._gopy_PyList_SET_ITEM(topy(self), C.Py_ssize_t(index), topy(o))
}

/*
int PyList_Insert(PyObject *list, Py_ssize_t index, PyObject *item)
Insert the item item into list list in front of index index. Return 0 if successful; return -1 and set an exception if unsuccessful. Analogous to list.insert(index, item).

Changed in version 2.5: This function used an int for index. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyList_Insert(self *PyObject, index int, item *PyObject) os.Error {
	err := C.PyList_Insert(topy(self), C.Py_ssize_t(index), topy(item))
	return int2err(err)
}

/*
int PyList_Append(PyObject *list, PyObject *item)
Append the object item at the end of list list. Return 0 if successful; return -1 and set an exception if unsuccessful. Analogous to list.append(item).
PyObject* PyList_GetSlice(PyObject *list, Py_ssize_t low, Py_ssize_t high)
Return value: New reference.
Return a list of the objects in list containing the objects between low and high. Return NULL and set an exception if unsuccessful. Analogous to list[low:high]. Negative indices, as when slicing from Python, are not supported.

Changed in version 2.5: This function used an int for low and high. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyList_Append(self, item *PyObject) os.Error {
	err := C.PyList_Append(topy(self), topy(item))
	return int2err(err)
}

/*
int PyList_SetSlice(PyObject *list, Py_ssize_t low, Py_ssize_t high, PyObject *itemlist)
Set the slice of list between low and high to the contents of itemlist. Analogous to list[low:high] = itemlist. The itemlist may be NULL, indicating the assignment of an empty list (slice deletion). Return 0 on success, -1 on failure. Negative indices, as when slicing from Python, are not supported.

Changed in version 2.5: This function used an int for low and high. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyList_SetSlice(self *PyObject, low, high int, itemlist *PyObject) os.Error {
	err := C.PyList_SetSlice(topy(self), C.Py_ssize_t(low), C.Py_ssize_t(high),
		topy(itemlist))
	return int2err(err)
}

/*
int PyList_Sort(PyObject *list)
Sort the items of list in place. Return 0 on success, -1 on failure. This is equivalent to list.sort().
*/
func PyList_Sort(self *PyObject) os.Error {
	err := C.PyList_Sort(topy(self))
	return int2err(err)
}

/*
int PyList_Reverse(PyObject *list)
Reverse the items of list in place. Return 0 on success, -1 on failure. This is the equivalent of list.reverse().
*/
func PyList_Reverse(self *PyObject) os.Error {
	err := C.PyList_Reverse(topy(self))
	return int2err(err)
}

/*
PyObject* PyList_AsTuple(PyObject *list)
Return value: New reference.
Return a new tuple object containing the contents of list; equivalent to tuple(list).
*/
func PyList_AsTuple(self *PyObject) *PyObject {
	return togo(C.PyList_AsTuple(topy(self)))
}

// EOF

