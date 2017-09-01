package python

//#include "go-python.h"
import "C"

import (
	"unsafe"
)

// int PyDict_Check(PyObject *p)
// Return true if p is a dict object or an instance of a subtype of the dict type.
//
// Changed in version 2.2: Allowed subtypes to be accepted.
func PyDict_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyDict_Check(topy(self)))
}

// int PyDict_CheckExact(PyObject *p)
// Return true if p is a dict object, but not an instance of a subtype of the dict type.
//
// New in version 2.4.
func PyDict_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyDict_CheckExact(topy(self)))
}

// PyObject* PyDict_New()
// Return value: New reference.
// Return a new empty dictionary, or NULL on failure.
func PyDict_New() *PyObject {
	return togo(C.PyDict_New())
}

// PyObject* PyDictProxy_New(PyObject *dict)
// Return value: New reference.
// Return a proxy object for a mapping which enforces read-only behavior. This is normally used to create a proxy to prevent modification of the dictionary for non-dynamic class types.
//
// New in version 2.2.
func PyDictProxy_New(self *PyObject) *PyObject {
	return togo(C.PyDictProxy_New(topy(self)))
}

// void PyDict_Clear(PyObject *p)
// Empty an existing dictionary of all key-value pairs.
func PyDict_Clear(self *PyObject) {
	C.PyDict_Clear(topy(self))
}

// int PyDict_Contains(PyObject *p, PyObject *key)
// Determine if dictionary p contains key. If an item in p is matches key, return 1, otherwise return 0. On error, return -1. This is equivalent to the Python expression key in p.
//
// New in version 2.4.
func PyDict_Contains(self, key *PyObject) (bool, error) {
	err := C.PyDict_Contains(topy(self), topy(key))
	if err != -1 {
		return int2bool(err), nil
	}
	return false, int2err(err)
}

// PyObject* PyDict_Copy(PyObject *p)
// Return value: New reference.
// Return a new dictionary that contains the same key-value pairs as p.
//
// New in version 1.6.
func PyDict_Copy(self *PyObject) *PyObject {
	return togo(C.PyDict_Copy(topy(self)))
}

// int PyDict_SetItem(PyObject *p, PyObject *key, PyObject *val)
// Insert value into the dictionary p with a key of key. key must be hashable; if it isn’t, TypeError will be raised. Return 0 on success or -1 on failure.
// int PyDict_SetItemString(PyObject *p, const char *key, PyObject *val)
// Insert value into the dictionary p using key as a key. key should be a char*. The key object is created using PyString_FromString(key). Return 0 on success or -1 on failure.
func PyDict_SetItem(self, key, val *PyObject) error {
	err := C.PyDict_SetItem(topy(self), topy(key), topy(val))
	return int2err(err)
}

// int PyDict_DelItem(PyObject *p, PyObject *key)
// Remove the entry in dictionary p with key key. key must be hashable; if it isn’t, TypeError is raised. Return 0 on success or -1 on failure.
// int PyDict_DelItemString(PyObject *p, char *key)
// Remove the entry in dictionary p which has a key specified by the string key. Return 0 on success or -1 on failure.
func PyDict_DelItem(self, key *PyObject) error {
	err := C.PyDict_DelItem(topy(self), topy(key))
	return int2err(err)
}

// PyObject* PyDict_GetItem(PyObject *p, PyObject *key)
// Return value: Borrowed reference.
// Return the object from dictionary p which has a key key. Return NULL if the key key is not present, but without setting an exception.
func PyDict_GetItem(self, key *PyObject) *PyObject {
	return togo(C.PyDict_GetItem(topy(self), topy(key)))
}

// PyObject* PyDict_GetItemString(PyObject *p, const char *key)
// Return value: Borrowed reference.
// This is the same as PyDict_GetItem(), but key is specified as a char*, rather than a PyObject*.
func PyDict_GetItemString(self *PyObject, key string) *PyObject {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	return togo(C.PyDict_GetItemString(topy(self), c_key))
}

// PyObject* PyDict_Items(PyObject *p)
// Return value: New reference.
// Return a PyListObject containing all the items from the dictionary, as in the dictionary method dict.items().
func PyDict_Items(self *PyObject) *PyObject {
	return togo(C.PyDict_Items(topy(self)))
}

// PyObject* PyDict_Keys(PyObject *p)
// Return value: New reference.
// Return a PyListObject containing all the keys from the dictionary, as in the dictionary method dict.keys().
func PyDict_Keys(self *PyObject) *PyObject {
	return togo(C.PyDict_Keys(topy(self)))
}

// PyObject* PyDict_Values(PyObject *p)
// Return value: New reference.
// Return a PyListObject containing all the values from the dictionary p, as in the dictionary method dict.values().
func PyDict_Values(self *PyObject) *PyObject {
	return togo(C.PyDict_Values(topy(self)))
}

// Py_ssize_t PyDict_Size(PyObject *p)
// Return the number of items in the dictionary. This is equivalent to len(p) on a dictionary.

// Changed in version 2.5: This function returned an int type. This might require changes in your code for properly supporting 64-bit systems.
func PyDict_Size(self *PyObject) int {
	return int(C.PyDict_Size(topy(self)))
}

// int PyDict_Next(PyObject *p, Py_ssize_t *ppos, PyObject **pkey, PyObject **pvalue)
// Iterate over all key-value pairs in the dictionary p. The Py_ssize_t referred to by ppos must be initialized to 0 prior to the first call to this function to start the iteration; the function returns true for each pair in the dictionary, and false once all pairs have been reported. The parameters pkey and pvalue should either point to PyObject* variables that will be filled in with each key and value, respectively, or may be NULL. Any references returned through them are borrowed. ppos should not be altered during iteration. Its value represents offsets within the internal dictionary structure, and since the structure is sparse, the offsets are not consecutive.
//
// For example:
//
// PyObject *key, *value;
// Py_ssize_t pos = 0;
//
// while (PyDict_Next(self->dict, &pos, &key, &value)) {
//     // do something interesting with the values...
//     ...
// }
// The dictionary p should not be mutated during iteration. It is safe (since Python 2.1) to modify the values of the keys as you iterate over the dictionary, but only so long as the set of keys does not change. For example:
//
// PyObject *key, *value;
// Py_ssize_t pos = 0;
//
// while (PyDict_Next(self->dict, &pos, &key, &value)) {
//     int i = PyInt_AS_LONG(value) + 1;
//     PyObject *o = PyInt_FromLong(i);
//     if (o == NULL)
//         return -1;
//     if (PyDict_SetItem(self->dict, key, o) < 0) {
//         Py_DECREF(o);
//         return -1;
//     }
//     Py_DECREF(o);
// }
// Changed in version 2.5: This function used an int * type for ppos. This might require changes in your code for properly supporting 64-bit systems.
func PyDict_Next(self *PyObject, pos *int, key, value **PyObject) bool {
	if pos == nil {
		return false
	}

	c_pos := C.Py_ssize_t(*pos)
	c_key := topy(*key)
	c_val := topy(*value)

	err := C.PyDict_Next(topy(self), &c_pos, &c_key, &c_val)

	*pos = int(c_pos)
	*key = togo(c_key)
	*value = togo(c_val)

	return int2bool(err)
}

// int PyDict_Merge(PyObject *a, PyObject *b, int override)
// Iterate over mapping object b adding key-value pairs to dictionary a. b may be a dictionary, or any object supporting PyMapping_Keys() and PyObject_GetItem(). If override is true, existing pairs in a will be replaced if a matching key is found in b, otherwise pairs will only be added if there is not a matching key in a. Return 0 on success or -1 if an exception was raised.
//
// New in version 2.2.
func PyDict_Merge(a, b *PyObject, override int) error {
	err := C.PyDict_Merge(topy(a), topy(b), C.int(override))
	return int2err(err)
}

// int PyDict_Update(PyObject *a, PyObject *b)
// This is the same as PyDict_Merge(a, b, 1) in C, or a.update(b) in Python. Return 0 on success or -1 if an exception was raised.
//
// New in version 2.2.
func PyDict_Update(a, b *PyObject) error {
	err := C.PyDict_Update(topy(a), topy(b))
	return int2err(err)
}

// int PyDict_MergeFromSeq2(PyObject *a, PyObject *seq2, int override)
// Update or merge into dictionary a, from the key-value pairs in seq2. seq2 must be an iterable object producing iterable objects of length 2, viewed as key-value pairs. In case of duplicate keys, the last wins if override is true, else the first wins. Return 0 on success or -1 if an exception was raised. Equivalent Python (except for the return value):
//
// def PyDict_MergeFromSeq2(a, seq2, override):
//     for key, value in seq2:
//         if override or key not in a:
//             a[key] = value
// New in version 2.2.
func PyDict_MergeFromSeq2(a, seq2 *PyObject, override int) error {
	err := C.PyDict_MergeFromSeq2(topy(a), topy(seq2), C.int(override))
	return int2err(err)
}

// EOF
