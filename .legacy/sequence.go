package python

// #include "go-python.h"
import "C"

import (
	"unsafe"
)

//////// bytearray ////////

// int PyByteArray_Check(PyObject *o)
// Return true if the object o is a bytearray object or an instance of a subtype of the bytearray type.
func PyByteArray_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyByteArray_Check(topy(self)))
}

// int PyByteArray_CheckExact(PyObject *o)
// Return true if the object o is a bytearray object, but not an instance of a subtype of the bytearray type.
func PyByteArray_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyByteArray_CheckExact(topy(self)))
}

// PyObject* PyByteArray_FromObject(PyObject *o)
// Return a new bytearray object from any object, o, that implements the buffer protocol.
func PyByteArray_FromObject(self *PyObject) *PyObject {
	return togo(C.PyByteArray_FromObject(topy(self)))
}

// PyObject* PyByteArray_FromStringAndSize(const char *string, Py_ssize_t len)
// Create a new bytearray object from string and its length, len. On failure, NULL is returned.
func PyByteArray_FromStringAndSize(str string) *PyObject {
	//FIXME use []byte instead ?
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	return togo(C.PyByteArray_FromStringAndSize(c_str, C.Py_ssize_t(len(str))))
}

// PyObject* PyByteArray_Concat(PyObject *a, PyObject *b)
// Concat bytearrays a and b and return a new bytearray with the result.
func PyByteArray_Concat(a, b *PyObject) *PyObject {
	return togo(C.PyByteArray_Concat(topy(a), topy(b)))
}

// Py_ssize_t PyByteArray_Size(PyObject *bytearray)
// Return the size of bytearray after checking for a NULL pointer.
func PyByteArray_Size(self *PyObject) int {
	return int(C.PyByteArray_Size(topy(self)))
}

// char* PyByteArray_AsString(PyObject *bytearray)
// Return the contents of bytearray as a char array after checking for a NULL pointer.
func PyByteArray_AsString(self *PyObject) string {
	c_str := C.PyByteArray_AsString(topy(self))
	return C.GoString(c_str)
}

// PyByteArray_AsBytes returns the contents of bytearray as []bytes
func PyByteArray_AsBytes(self *PyObject) []byte {
	length := C._gopy_PyByteArray_GET_SIZE(topy(self))
	c_str := C.PyByteArray_AsString(topy(self))
	return C.GoBytes(unsafe.Pointer(c_str), C.int(length))
}

// PyByteArray_AsBytesN returns the contents of bytearray as []bytes, size length
func PyByteArray_AsBytesN(self *PyObject, length int) []byte {
	blength := int(C._gopy_PyByteArray_GET_SIZE(topy(self)))
	if (blength < length) || (length < 0) {
		panic("bytearray length out of range")
	}
	c_str := C.PyByteArray_AsString(topy(self))
	return C.GoBytes(unsafe.Pointer(c_str), C.int(length))
}

// int PyByteArray_Resize(PyObject *bytearray, Py_ssize_t len)
// Resize the internal buffer of bytearray to len.
func PyByteArray_Resize(self *PyObject, sz int) error {
	return int2err(C.PyByteArray_Resize(topy(self), C.Py_ssize_t(sz)))
}

// char* PyByteArray_AS_STRING(PyObject *bytearray)
// Macro version of PyByteArray_AsString().
func PyByteArray_AS_STRING(self *PyObject) string {
	c_str := C._gopy_PyByteArray_AS_STRING(topy(self))
	return C.GoString(c_str)
}

// Py_ssize_t PyByteArray_GET_SIZE(PyObject *bytearray)
// Macro version of PyByteArray_Size().
func PyByteArray_GET_SIZE(self *PyObject) int {
	return int(C._gopy_PyByteArray_GET_SIZE(topy(self)))
}

//////// tuple /////////

// int PyTuple_Check(PyObject *p)
// Return true if p is a tuple object or an instance of a subtype of the tuple type.
//
// Changed in version 2.2: Allowed subtypes to be accepted.
func PyTuple_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyTuple_Check(topy(self)))
}

// int PyTuple_CheckExact(PyObject *p)
// Return true if p is a tuple object, but not an instance of a subtype of the tuple type.
//
// New in version 2.2.
func PyTuple_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyTuple_CheckExact(topy(self)))
}

// PyObject* PyTuple_New(Py_ssize_t len)
// Return value: New reference.
// Return a new tuple object of size len, or NULL on failure.
//
// Changed in version 2.5: This function used an int type for len. This might require changes in your code for properly supporting 64-bit systems.
func PyTuple_New(sz int) *PyObject {
	return togo(C.PyTuple_New(C.Py_ssize_t(sz)))
}

// PyObject* PyTuple_Pack(Py_ssize_t n, ...)
// Return value: New reference.
// Return a new tuple object of size n, or NULL on failure. The tuple values are initialized to the subsequent n C arguments pointing to Python objects. PyTuple_Pack(2, a, b) is equivalent to Py_BuildValue("(OO)", a, b).
//
// New in version 2.4.
//
// Changed in version 2.5: This function used an int type for n. This might require changes in your code for properly supporting 64-bit systems.
func PyTuple_Pack(n int, objs ...*PyObject) *PyObject {
	//FIXME
	panic("not implemented")
}

// Py_ssize_t PyTuple_Size(PyObject *p)
// Take a pointer to a tuple object, and return the size of that tuple.
//
// Changed in version 2.5: This function returned an int type. This might require changes in your code for properly supporting 64-bit systems.
func PyTuple_Size(self *PyObject) int {
	return int(C.PyTuple_Size(topy(self)))
}

// Py_ssize_t PyTuple_GET_SIZE(PyObject *p)
// Return the size of the tuple p, which must be non-NULL and point to a tuple; no error checking is performed.
//
// Changed in version 2.5: This function returned an int type. This might require changes in your code for properly supporting 64-bit systems.
func PyTuple_GET_SIZE(self *PyObject) int {
	return int(C._gopy_PyTuple_GET_SIZE(topy(self)))
}

// PyObject* PyTuple_GetItem(PyObject *p, Py_ssize_t pos)
// Return value: Borrowed reference.
// Return the object at position pos in the tuple pointed to by p. If pos is out of bounds, return NULL and sets an IndexError exception.
//
// Changed in version 2.5: This function used an int type for pos. This might require changes in your code for properly supporting 64-bit systems.
func PyTuple_GetItem(self *PyObject, pos int) *PyObject {
	return togo(C.PyTuple_GetItem(topy(self), C.Py_ssize_t(pos)))
}

// PyObject* PyTuple_GET_ITEM(PyObject *p, Py_ssize_t pos)
// Return value: Borrowed reference.
// Like PyTuple_GetItem(), but does no checking of its arguments.
//
// Changed in version 2.5: This function used an int type for pos. This might require changes in your code for properly supporting 64-bit systems.
func PyTuple_GET_ITEM(self *PyObject, pos int) *PyObject {
	return togo(C._gopy_PyTuple_GET_ITEM(topy(self), C.Py_ssize_t(pos)))
}

// PyObject* PyTuple_GetSlice(PyObject *p, Py_ssize_t low, Py_ssize_t high)
// Return value: New reference.
// Take a slice of the tuple pointed to by p from low to high and return it as a new tuple.
//
// Changed in version 2.5: This function used an int type for low and high. This might require changes in your code for properly supporting 64-bit systems.
func PyTuple_GetSlice(self *PyObject, low, high int) *PyObject {
	return togo(C.PyTuple_GetSlice(topy(self), C.Py_ssize_t(low), C.Py_ssize_t(high)))
}

// int PyTuple_SetItem(PyObject *p, Py_ssize_t pos, PyObject *o)
// Insert a reference to object o at position pos of the tuple pointed to by p. Return 0 on success.
//
// Note This function “steals” a reference to o.
// Changed in version 2.5: This function used an int type for pos. This might require changes in your code for properly supporting 64-bit systems.
func PyTuple_SetItem(self *PyObject, pos int, o *PyObject) error {
	return int2err(C.PyTuple_SetItem(topy(self), C.Py_ssize_t(pos), topy(o)))
}

// void PyTuple_SET_ITEM(PyObject *p, Py_ssize_t pos, PyObject *o)
// Like PyTuple_SetItem(), but does no error checking, and should only be used to fill in brand new tuples.
//
// Note This function “steals” a reference to o.
// Changed in version 2.5: This function used an int type for pos. This might require changes in your code for properly supporting 64-bit systems.
func PyTuple_SET_ITEM(self *PyObject, pos int, o *PyObject) {
	py_self := topy(self)
	py_pos := C.Py_ssize_t(pos)
	py_o := topy(o)
	C._gopy_PyTuple_SET_ITEM(py_self, py_pos, py_o)
}

// int _PyTuple_Resize(PyObject **p, Py_ssize_t newsize)
// Can be used to resize a tuple. newsize will be the new length of the tuple. Because tuples are supposed to be immutable, this should only be used if there is only one reference to the object. Do not use this if the tuple may already be known to some other part of the code. The tuple will always grow or shrink at the end. Think of this as destroying the old tuple and creating a new one, only more efficiently. Returns 0 on success. Client code should never assume that the resulting value of *p will be the same as before calling this function. If the object referenced by *p is replaced, the original *p is destroyed. On failure, returns -1 and sets *p to NULL, and raises MemoryError or SystemError.
//
// Changed in version 2.2: Removed unused third parameter, last_is_sticky.
//
// Changed in version 2.5: This function used an int type for newsize. This might require changes in your code for properly supporting 64-bit systems.
func PyTuple_Resize(self *PyObject, newsize int) error {
	py_self := topy(self)
	py_newsz := C.Py_ssize_t(newsize)
	err := C._PyTuple_Resize(&py_self, py_newsz)
	return int2err(err)
}

// int PyTuple_ClearFreeList()
// Clear the free list. Return the total number of freed items.
//
// New in version 2.6.
func PyTuple_ClearFreeList() {
	C.PyTuple_ClearFreeList()
}

/////////// list ///////////

// int PyList_Check(PyObject *p)
// Return true if p is a list object or an instance of a subtype of the list type.
//
// Changed in version 2.2: Allowed subtypes to be accepted.
func PyList_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyList_Check(topy(self)))
}

// int PyList_CheckExact(PyObject *p)
// Return true if p is a list object, but not an instance of a subtype of the list type.
//
// New in version 2.2.
func PyList_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyList_CheckExact(topy(self)))
}

// PyObject* PyList_New(Py_ssize_t len)
// Return value: New reference.
// Return a new list of length len on success, or NULL on failure.
//
// Note If length is greater than zero, the returned list object’s items are set to NULL. Thus you cannot use abstract API functions such as PySequence_SetItem() or expose the object to Python code before setting all items to a real object with PyList_SetItem().
// Changed in version 2.5: This function used an int for size. This might require changes in your code for properly supporting 64-bit systems.
func PyList_New(sz int) *PyObject {
	return togo(C.PyList_New(C.Py_ssize_t(sz)))
}

// Py_ssize_t PyList_Size(PyObject *list)
// Return the length of the list object in list; this is equivalent to len(list) on a list object.
//
// Changed in version 2.5: This function returned an int. This might require changes in your code for properly supporting 64-bit systems.
func PyList_Size(self *PyObject) int {
	return int(C.PyList_Size(topy(self)))
}

// Py_ssize_t PyList_GET_SIZE(PyObject *list)
// Macro form of PyList_Size() without error checking.
//
// Changed in version 2.5: This macro returned an int. This might require changes in your code for properly supporting 64-bit systems.
func PyList_GET_SIZE(self *PyObject) int {
	return int(C._gopy_PyList_GET_SIZE(topy(self)))
}

// PyObject* PyList_GetItem(PyObject *list, Py_ssize_t index)
// Return value: Borrowed reference.
// Return the object at position pos in the list pointed to by p. The position must be positive, indexing from the end of the list is not supported. If pos is out of bounds, return NULL and set an IndexError exception.
//
// Changed in version 2.5: This function used an int for index. This might require changes in your code for properly supporting 64-bit systems.
func PyList_GetItem(self *PyObject, index int) *PyObject {
	return togo(C.PyList_GetItem(topy(self), C.Py_ssize_t(index)))
}

// PyObject* PyList_GET_ITEM(PyObject *list, Py_ssize_t i)
// Return value: Borrowed reference.
// Macro form of PyList_GetItem() without error checking.
//
// Changed in version 2.5: This macro used an int for i. This might require changes in your code for properly supporting 64-bit systems.
func PyList_GET_ITEM(self *PyObject, index int) *PyObject {
	return togo(C._gopy_PyList_GET_ITEM(topy(self), C.Py_ssize_t(index)))
}

// int PyList_SetItem(PyObject *list, Py_ssize_t index, PyObject *item)
// Set the item at index index in list to item. Return 0 on success or -1 on failure.
//
// Note This function “steals” a reference to item and discards a reference to an item already in the list at the affected position.
// Changed in version 2.5: This function used an int for index. This might require changes in your code for properly supporting 64-bit systems.
func PyList_SetItem(self *PyObject, index int, item *PyObject) error {
	err := C.PyList_SetItem(topy(self), C.Py_ssize_t(index), topy(item))
	return int2err(err)
}

// void PyList_SET_ITEM(PyObject *list, Py_ssize_t i, PyObject *o)
// Macro form of PyList_SetItem() without error checking. This is normally only used to fill in new lists where there is no previous content.
//
// Note This macro “steals” a reference to item, and, unlike PyList_SetItem(), does not discard a reference to any item that it being replaced; any reference in list at position i will be leaked.
// Changed in version 2.5: This macro used an int for i. This might require changes in your code for properly supporting 64-bit systems.
func PyList_SET_ITEM(self *PyObject, index int, o *PyObject) {
	C._gopy_PyList_SET_ITEM(topy(self), C.Py_ssize_t(index), topy(o))
}

// int PyList_Insert(PyObject *list, Py_ssize_t index, PyObject *item)
// Insert the item item into list list in front of index index. Return 0 if successful; return -1 and set an exception if unsuccessful. Analogous to list.insert(index, item).
//
// Changed in version 2.5: This function used an int for index. This might require changes in your code for properly supporting 64-bit systems.
func PyList_Insert(self *PyObject, index int, item *PyObject) error {
	err := C.PyList_Insert(topy(self), C.Py_ssize_t(index), topy(item))
	return int2err(err)
}

// int PyList_Append(PyObject *list, PyObject *item)
// Append the object item at the end of list list. Return 0 if successful; return -1 and set an exception if unsuccessful. Analogous to list.append(item).
// PyObject* PyList_GetSlice(PyObject *list, Py_ssize_t low, Py_ssize_t high)
// Return value: New reference.
// Return a list of the objects in list containing the objects between low and high. Return NULL and set an exception if unsuccessful. Analogous to list[low:high]. Negative indices, as when slicing from Python, are not supported.
//
// Changed in version 2.5: This function used an int for low and high. This might require changes in your code for properly supporting 64-bit systems.
func PyList_Append(self, item *PyObject) error {
	err := C.PyList_Append(topy(self), topy(item))
	return int2err(err)
}

// int PyList_SetSlice(PyObject *list, Py_ssize_t low, Py_ssize_t high, PyObject *itemlist)
// Set the slice of list between low and high to the contents of itemlist. Analogous to list[low:high] = itemlist. The itemlist may be NULL, indicating the assignment of an empty list (slice deletion). Return 0 on success, -1 on failure. Negative indices, as when slicing from Python, are not supported.
//
// Changed in version 2.5: This function used an int for low and high. This might require changes in your code for properly supporting 64-bit systems.
func PyList_SetSlice(self *PyObject, low, high int, itemlist *PyObject) error {
	err := C.PyList_SetSlice(
		topy(self), C.Py_ssize_t(low), C.Py_ssize_t(high), topy(itemlist),
	)
	return int2err(err)
}

// int PyList_Sort(PyObject *list)
// Sort the items of list in place. Return 0 on success, -1 on failure. This is equivalent to list.sort().
func PyList_Sort(self *PyObject) error {
	err := C.PyList_Sort(topy(self))
	return int2err(err)
}

// int PyList_Reverse(PyObject *list)
// Reverse the items of list in place. Return 0 on success, -1 on failure. This is the equivalent of list.reverse().
func PyList_Reverse(self *PyObject) error {
	err := C.PyList_Reverse(topy(self))
	return int2err(err)
}

// PyObject* PyList_AsTuple(PyObject *list)
// Return value: New reference.
// Return a new tuple object containing the contents of list; equivalent to tuple(list).
func PyList_AsTuple(self *PyObject) *PyObject {
	return togo(C.PyList_AsTuple(topy(self)))
}

/////// string ////////
type PyString PyObject

// int PyString_Check(PyObject *o)
// Return true if the object o is a string object or an instance of a subtype of the string type.
//
// Changed in version 2.2: Allowed subtypes to be accepted.
func PyString_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyString_Check(self.ptr))
}

// func (self *PyString) Check() int {
// 	return int(C.PyString_Check(self.topy()))
// }

// PyObject* PyString_FromString(const char *v)
// Return value: New reference.
// Return a new string object with a copy of the string v as value on success, and NULL on failure. The parameter v must not be NULL; it will not be checked.
func PyString_FromString(v string) *PyObject {
	cstr := C.CString(v)
	defer C.free(unsafe.Pointer(cstr))
	return togo(C.PyString_FromString(cstr))
}

// PyObject* PyString_FromStringAndSize(const char *v, Py_ssize_t len)
// Return value: New reference.
// Return a new string object with a copy of the string v as value and length len on success, and NULL on failure. If v is NULL, the contents of the string are uninitialized.
//
// Changed in version 2.5: This function used an int type for len. This might require changes in your code for properly supporting 64-bit systems.
func PyString_FromStringAndSize(v string, sz int) *PyObject {
	cstr := C.CString(v)
	defer C.free(unsafe.Pointer(cstr))
	return togo(C.PyString_FromStringAndSize(cstr, C.Py_ssize_t(sz)))
}

// PyObject* PyString_FromFormat(const char *format, ...)
// Return value: New reference.
// Take a C printf()-style format string and a variable number of arguments, calculate the size of the resulting Python string and return a string with the values formatted into it. The variable arguments must be C types and must correspond exactly to the format characters in the format string. The following format characters are allowed:
//
// Format Characters	Type	Comment
// %%	n/a	The literal % character.
// %c	int	A single character, represented as an C int.
// %d	int	Exactly equivalent to printf("%d").
// %u	unsigned int	Exactly equivalent to printf("%u").
// %ld	long	Exactly equivalent to printf("%ld").
// %lu	unsigned long	Exactly equivalent to printf("%lu").
// %lld	long long	Exactly equivalent to printf("%lld").
// %llu	unsigned long long	Exactly equivalent to printf("%llu").
// %zd	Py_ssize_t	Exactly equivalent to printf("%zd").
// %zu	size_t	Exactly equivalent to printf("%zu").
// %i	int	Exactly equivalent to printf("%i").
// %x	int	Exactly equivalent to printf("%x").
// %s	char*	A null-terminated C character array.
// %p	void*	The hex representation of a C pointer. Mostly equivalent to printf("%p") except that it is guaranteed to start with the literal 0x regardless of what the platform’s printf yields.
// An unrecognized format character causes all the rest of the format string to be copied as-is to the result string, and any extra arguments discarded.
//
// Note The “%lld” and “%llu” format specifiers are only available when HAVE_LONG_LONG is defined.
// Changed in version 2.7: Support for “%lld” and “%llu” added.
func PyString_FromFormat(format string, args ...interface{}) *PyObject {
	cfmt := C.CString(format)
	defer C.free(unsafe.Pointer(cfmt))

	// FIXME
	panic("not implemented")
	return nil
}

// PyObject* PyString_FromFormatV(const char *format, va_list vargs)
// Return value: New reference.
// Identical to PyString_FromFormat() except that it takes exactly two arguments.
func PyStringFromFormatV(format string, args ...interface{}) *PyObject {
	cfmt := C.CString(format)
	defer C.free(unsafe.Pointer(cfmt))

	// FIXME
	panic("not implemented")
	return nil
}

// Py_ssize_t PyString_Size(PyObject *string)
// Return the length of the string in string object string.
//
// Changed in version 2.5: This function returned an int type. This might require changes in your code for properly supporting 64-bit systems.
func PyString_Size(self *PyObject) int {
	sz := C.PyString_Size(topy(self))
	return int(sz)
}

// Py_ssize_t PyString_GET_SIZE(PyObject *string)
// Macro form of PyString_Size() but without error checking.
//
// Changed in version 2.5: This macro returned an int type. This might require changes in your code for properly supporting 64-bit systems.
func PyString_GET_SIZE(self *PyObject) int {
	sz := C._gopy_PyString_GET_SIZE(topy(self))
	return int(sz)
}

// char* PyString_AsString(PyObject *string)
// Return a NUL-terminated representation of the contents of string. The pointer refers to the internal buffer of string, not a copy. The data must not be modified in any way, unless the string was just created using PyString_FromStringAndSize(NULL, size). It must not be deallocated. If string is a Unicode object, this function computes the default encoding of string and operates on that. If string is not a string object at all, PyString_AsString() returns NULL and raises TypeError.
func PyString_AsString(self *PyObject) string {
	c_str := C.PyString_AsString(self.ptr)
	// we dont own c_str...
	//defer C.free(unsafe.Pointer(c_str))
	return C.GoString(c_str)
}

// char* PyString_AS_STRING(PyObject *string)
// Macro form of PyString_AsString() but without error checking. Only string objects are supported; no Unicode objects should be passed.
func PyString_AS_STRING(self *PyObject) string {
	c_str := C._gopy_PyString_AS_STRING(self.ptr)
	// we dont own c_str...
	//defer C.free(unsafe.Pointer(c_str))
	return C.GoString(c_str)
}

// int PyString_AsStringAndSize(PyObject *obj, char **buffer, Py_ssize_t *length)
// Return a NUL-terminated representation of the contents of the object obj through the output variables buffer and length.
//
// The function accepts both string and Unicode objects as input. For Unicode objects it returns the default encoded version of the object. If length is NULL, the resulting buffer may not contain NUL characters; if it does, the function returns -1 and a TypeError is raised.
//
// The buffer refers to an internal string buffer of obj, not a copy. The data must not be modified in any way, unless the string was just created using PyString_FromStringAndSize(NULL, size). It must not be deallocated. If string is a Unicode object, this function computes the default encoding of string and operates on that. If string is not a string object at all, PyString_AsStringAndSize() returns -1 and raises TypeError.
//
// Changed in version 2.5: This function used an int * type for length. This might require changes in your code for properly supporting 64-bit systems.
func PyString_AsStringAndSize(self *PyObject, sz int) (str string, err int) {
	// FIXME
	panic("not implemented")
}

// void PyString_Concat(PyObject **string, PyObject *newpart)
// Create a new string object in *string containing the contents of newpart appended to string; the caller will own the new reference. The reference to the old value of string will be stolen. If the new string cannot be created, the old reference to string will still be discarded and the value of *string will be set to NULL; the appropriate exception will be set.
func PyString_Concat(self, newpart *PyObject) *PyObject {
	// FIXME
	panic("not implemented")
}

// void PyString_ConcatAndDel(PyObject **string, PyObject *newpart)
// Create a new string object in *string containing the contents of newpart appended to string. This version decrements the reference count of newpart.
// int _PyString_Resize(PyObject **string, Py_ssize_t newsize)
// A way to resize a string object even though it is “immutable”. Only use this to build up a brand new string object; don’t use this if the string may already be known in other parts of the code. It is an error to call this function if the refcount on the input string object is not one. Pass the address of an existing string object as an lvalue (it may be written into), and the new size desired. On success, *string holds the resized string object and 0 is returned; the address in *string may differ from its input value. If the reallocation fails, the original string object at *string is deallocated, *string is set to NULL, a memory exception is set, and -1 is returned.
//
// Changed in version 2.5: This function used an int type for newsize. This might require changes in your code for properly supporting 64-bit systems.
func PyString_ConcatAndDel(self, newpart *PyObject) *PyObject {
	// FIXME
	panic("not implemented")
}

// PyObject* PyString_Format(PyObject *format, PyObject *args)
// Return value: New reference.
// Return a new string object from format and args. Analogous to format % args. The args argument must be a tuple.
func PyString_Format(format, args *PyObject) *PyObject {
	py_format := topy(format)
	py_args := topy(args)
	return togo(C.PyString_Format(py_format, py_args))
}

// void PyString_InternInPlace(PyObject **string)
// Intern the argument *string in place. The argument must be the address of a pointer variable pointing to a Python string object. If there is an existing interned string that is the same as *string, it sets *string to it (decrementing the reference count of the old string object and incrementing the reference count of the interned string object), otherwise it leaves *string alone and interns it (incrementing its reference count). (Clarification: even though there is a lot of talk about reference counts, think of this function as reference-count-neutral; you own the object after the call if and only if you owned it before the call.)
//
// Note This function is not available in 3.x and does not have a PyBytes alias.
func PyString_InternInPlace(self *PyObject) {
	//FIXME check everything is OK...
	s := topy(self)
	C.PyString_InternInPlace(&s)
}

// PyObject* PyString_InternFromString(const char *v)
// Return value: New reference.
// A combination of PyString_FromString() and PyString_InternInPlace(), returning either a new string object that has been interned, or a new (“owned”) reference to an earlier interned string object with the same value.
func PyString_InternFromString(v string) *PyObject {
	cstr := C.CString(v)
	defer C.free(unsafe.Pointer(cstr))

	return togo(C.PyString_InternFromString(cstr))
}

// Note This function is not available in 3.x and does not have a PyBytes alias.
// PyObject* PyString_Decode(const char *s, Py_ssize_t size, const char *encoding, const char *errors)
// Return value: New reference.
// Create an object by decoding size bytes of the encoded buffer s using the codec registered for encoding. encoding and errors have the same meaning as the parameters of the same name in the unicode() built-in function. The codec to be used is looked up using the Python codec registry. Return NULL if an exception was raised by the codec.
//
// Note This function is not available in 3.x and does not have a PyBytes alias.
// Changed in version 2.5: This function used an int type for size. This might require changes in your code for properly supporting 64-bit systems.
func PyString_Decode(s string, sz int, encoding, errors string) *PyObject {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))

	c_encoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(c_encoding))

	c_errors := C.CString(errors)
	defer C.free(unsafe.Pointer(c_errors))

	return togo(C.PyString_Decode(c_s, C.Py_ssize_t(sz), c_encoding, c_errors))
}

// PyObject* PyString_AsDecodedObject(PyObject *str, const char *encoding, const char *errors)
// Return value: New reference.
// Decode a string object by passing it to the codec registered for encoding and return the result as Python object. encoding and errors have the same meaning as the parameters of the same name in the string encode() method. The codec to be used is looked up using the Python codec registry. Return NULL if an exception was raised by the codec.

// Note This function is not available in 3.x and does not have a PyBytes alias.
func PyString_AsDecodedObject(self *PyObject, encoding, errors string) *PyObject {
	c_encoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(c_encoding))

	c_errors := C.CString(errors)
	defer C.free(unsafe.Pointer(c_errors))

	return togo(C.PyString_AsDecodedObject(topy(self), c_encoding, c_errors))
}

// PyObject* PyString_Encode(const char *s, Py_ssize_t size, const char *encoding, const char *errors)
// Return value: New reference.
// Encode the char buffer of the given size by passing it to the codec registered for encoding and return a Python object. encoding and errors have the same meaning as the parameters of the same name in the string encode() method. The codec to be used is looked up using the Python codec registry. Return NULL if an exception was raised by the codec.
//
// Note This function is not available in 3.x and does not have a PyBytes alias.
// Changed in version 2.5: This function used an int type for size. This might require changes in your code for properly supporting 64-bit systems.
func PyString_Encode(s, encoding, errors string) *PyObject {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))

	c_encoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(c_encoding))

	c_errors := C.CString(errors)
	defer C.free(unsafe.Pointer(c_errors))

	// FIXME should check if len is len of rune or of string
	return togo(C.PyString_Encode(c_s, C.Py_ssize_t(len(s)), c_encoding, c_errors))
}

// PyObject* PyString_AsEncodedObject(PyObject *str, const char *encoding, const char *errors)
// Return value: New reference.
// Encode a string object using the codec registered for encoding and return the result as Python object. encoding and errors have the same meaning as the parameters of the same name in the string encode() method. The codec to be used is looked up using the Python codec registry. Return NULL if an exception was raised by the codec.
//
// Note This function is not available in 3.x and does not have a PyBytes alias.
func PyString_AsEncodedObject(self *PyObject, encoding, errors string) *PyObject {
	c_encoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(c_encoding))

	c_errors := C.CString(errors)
	defer C.free(unsafe.Pointer(c_errors))

	return togo(C.PyString_AsEncodedObject(topy(self), c_encoding, c_errors))
}

/////////// buffer ///////////

// Py_buffer layer
type Py_buffer struct {
	ptr *C.Py_buffer
}

type PyBUF_Flag int

const (
	PyBUF_SIMPLE       = PyBUF_Flag(C.PyBUF_SIMPLE)
	PyBUF_WRITABLE     = PyBUF_Flag(C.PyBUF_WRITABLE)
	PyBUF_STRIDES      = PyBUF_Flag(C.PyBUF_STRIDES)
	PyBUF_ND           = PyBUF_Flag(C.PyBUF_ND)
	PyBUF_C_CONTIGUOUS = PyBUF_Flag(C.PyBUF_C_CONTIGUOUS)
	PyBUF_INDIRECT     = PyBUF_Flag(C.PyBUF_INDIRECT)
	PyBUF_FORMAT       = PyBUF_Flag(C.PyBUF_FORMAT)
	PyBUF_STRIDED      = PyBUF_Flag(C.PyBUF_STRIDED)
	PyBUF_STRIDED_RO   = PyBUF_Flag(C.PyBUF_STRIDED_RO)
	PyBUF_RECORDS      = PyBUF_Flag(C.PyBUF_RECORDS)
	PyBUF_RECORDS_RO   = PyBUF_Flag(C.PyBUF_RECORDS_RO)
	PyBUF_FULL         = PyBUF_Flag(C.PyBUF_FULL)
	PyBUF_FULL_RO      = PyBUF_Flag(C.PyBUF_FULL_RO)
	PyBUF_CONTIG       = PyBUF_Flag(C.PyBUF_CONTIG)
	PyBUF_CONTIG_RO    = PyBUF_Flag(C.PyBUF_CONTIG_RO)
)

// int PyObject_CheckBuffer(PyObject *obj)
// Return 1 if obj supports the buffer interface otherwise 0.
func PyObject_CheckBuffer(self *PyObject) bool {
	return int2bool(C._gopy_PyObject_CheckBuffer(topy(self)))
}

// int PyObject_GetBuffer(PyObject *obj, Py_buffer *view, int flags)
// Export obj into a Py_buffer, view. These arguments must never be NULL. The flags argument is a bit field indicating what kind of buffer the caller is prepared to deal with and therefore what kind of buffer the exporter is allowed to return. The buffer interface allows for complicated memory sharing possibilities, but some caller may not be able to handle all the complexity but may want to see if the exporter will let them take a simpler view to its memory.
//
// Some exporters may not be able to share memory in every possible way and may need to raise errors to signal to some consumers that something is just not possible. These errors should be a BufferError unless there is another error that is actually causing the problem. The exporter can use flags information to simplify how much of the Py_buffer structure is filled in with non-default values and/or raise an error if the object can’t support a simpler view of its memory.
//
// 0 is returned on success and -1 on error.
//
// The following table gives possible values to the flags arguments.
//
// Flag	Description
// PyBUF_SIMPLE	This is the default flag state. The returned buffer may or may not have writable memory. The format of the data will be assumed to be unsigned bytes. This is a “stand-alone” flag constant. It never needs to be ‘|’d to the others. The exporter will raise an error if it cannot provide such a contiguous buffer of bytes.
// PyBUF_WRITABLE	The returned buffer must be writable. If it is not writable, then raise an error.
// PyBUF_STRIDES	This implies PyBUF_ND. The returned buffer must provide strides information (i.e. the strides cannot be NULL). This would be used when the consumer can handle strided, discontiguous arrays. Handling strides automatically assumes you can handle shape. The exporter can raise an error if a strided representation of the data is not possible (i.e. without the suboffsets).
// PyBUF_ND	The returned buffer must provide shape information. The memory will be assumed C-style contiguous (last dimension varies the fastest). The exporter may raise an error if it cannot provide this kind of contiguous buffer. If this is not given then shape will be NULL.
// PyBUF_C_CONTIGUOUS PyBUF_F_CONTIGUOUS PyBUF_ANY_CONTIGUOUS	These flags indicate that the contiguity returned buffer must be respectively, C-contiguous (last dimension varies the fastest), Fortran contiguous (first dimension varies the fastest) or either one. All of these flags imply PyBUF_STRIDES and guarantee that the strides buffer info structure will be filled in correctly.
// PyBUF_INDIRECT	This flag indicates the returned buffer must have suboffsets information (which can be NULL if no suboffsets are needed). This can be used when the consumer can handle indirect array referencing implied by these suboffsets. This implies PyBUF_STRIDES.
// PyBUF_FORMAT	The returned buffer must have true format information if this flag is provided. This would be used when the consumer is going to be checking for what ‘kind’ of data is actually stored. An exporter should always be able to provide this information if requested. If format is not explicitly requested then the format must be returned as NULL (which means 'B', or unsigned bytes)
// PyBUF_STRIDED	This is equivalent to (PyBUF_STRIDES | PyBUF_WRITABLE).
// PyBUF_STRIDED_RO	This is equivalent to (PyBUF_STRIDES).
// PyBUF_RECORDS	This is equivalent to (PyBUF_STRIDES | PyBUF_FORMAT | PyBUF_WRITABLE).
// PyBUF_RECORDS_RO	This is equivalent to (PyBUF_STRIDES | PyBUF_FORMAT).
// PyBUF_FULL	This is equivalent to (PyBUF_INDIRECT | PyBUF_FORMAT | PyBUF_WRITABLE).
// PyBUF_FULL_RO	This is equivalent to (PyBUF_INDIRECT | PyBUF_FORMAT).
// PyBUF_CONTIG	This is equivalent to (PyBUF_ND | PyBUF_WRITABLE).
// PyBUF_CONTIG_RO	This is equivalent to (PyBUF_ND).
func PyObject_GetBuffer(self *PyObject, flags PyBUF_Flag) (buf *Py_buffer, err error) {
	buf.ptr = &C.Py_buffer{}
	err = int2err(C.PyObject_GetBuffer(topy(self), buf.ptr, C.int(flags)))
	return
}

// void PyBuffer_Release(Py_buffer *view)
// Release the buffer view. This should be called when the buffer is no longer being used as it may free memory from it.
func PyBuffer_Release(self *Py_buffer) {
	C.PyBuffer_Release(self.ptr)
}

// Py_ssize_t PyBuffer_SizeFromFormat(const char *)
// Return the implied ~Py_buffer.itemsize from the struct-stype ~Py_buffer.format.
func PyBuffer_SizeFromFormat(self *Py_buffer) int {
	//FIXME
	panic("not implemented")
}

// int PyObject_CopyToObject(PyObject *obj, void *buf, Py_ssize_t len, char fortran)
// Copy len bytes of data pointed to by the contiguous chunk of memory pointed to by buf into the buffer exported by obj. The buffer must of course be writable. Return 0 on success and return -1 and raise an error on failure. If the object does not have a writable buffer, then an error is raised. If fortran is 'F', then if the object is multi-dimensional, then the data will be copied into the array in Fortran-style (first dimension varies the fastest). If fortran is 'C', then the data will be copied into the array in C-style (last dimension varies the fastest). If fortran is 'A', then it does not matter and the copy will be made in whatever way is more efficient.
func PyObject_CopyToObject(self *PyObject, buf []byte, fortran string) error {
	/*
		c_buf := (*C.char)(unsafe.Pointer(&buf[0]))
		c_for := C.char(fortran[0])

		py_self := self.ptr
		//defer C.free(unsafe.Pointer(c_for))
		err := C.PyObject_CopyToObject(py_self, c_buf, C.Py_ssize_t(len(buf)), c_for)
	*/
	//FIXME
	panic("not implemented")
}

// int PyBuffer_IsContiguous(Py_buffer *view, char fortran)
// Return 1 if the memory defined by the view is C-style (fortran is 'C') or Fortran-style (fortran is 'F') contiguous or either one (fortran is 'A'). Return 0 otherwise.
func PyBuffer_IsContiguous(self *Py_buffer, fortran string) bool {
	c_fortran := C.char(fortran[0])
	return int2bool(C.PyBuffer_IsContiguous(self.ptr, c_fortran))
}

// void PyBuffer_FillContiguousStrides(int ndim, Py_ssize_t *shape, Py_ssize_t *strides, Py_ssize_t itemsize, char fortran)
// Fill the strides array with byte-strides of a contiguous (C-style if fortran is 'C' or Fortran-style if fortran is 'F' array of the given shape with the given number of bytes per element.
func PyBuffer_FillContiguousStrides(ndim int, shape, strides []int, itemsize int, fortran string) {
	//FIXME
	panic("not implemented")
}

// int PyBuffer_FillInfo(Py_buffer *view, PyObject *obj, void *buf, Py_ssize_t len, int readonly, int infoflags)
// Fill in a buffer-info structure, view, correctly for an exporter that can only share a contiguous chunk of memory of “unsigned bytes” of the given length. Return 0 on success and -1 (with raising an error) on error.
func PyBuffer_FillInfo(self *PyObject, buf []byte, readonly bool, infoflags int) (buffer *Py_buffer, err error) {
	//FIXME
	panic("not implemented")
}

///// memoryview /////

// A memoryview object exposes the new C level buffer interface as a Python object which can then be passed around like any other object.
//
// PyObject *PyMemoryView_FromObject(PyObject *obj)
// Create a memoryview object from an object that defines the new buffer interface.
func PyMemoryView_FromObject(obj *PyObject) *PyObject {
	return togo(C.PyMemoryView_FromObject(topy(obj)))
}

// PyObject *PyMemoryView_FromBuffer(Py_buffer *view)
// Create a memoryview object wrapping the given buffer-info structure view. The memoryview object then owns the buffer, which means you shouldn’t try to release it yourself: it will be released on deallocation of the memoryview object.
func PyMemoryView_FromBuffer(view *Py_buffer) *PyObject {
	return togo(C.PyMemoryView_FromBuffer(view.ptr))
}

// PyObject *PyMemoryView_GetContiguous(PyObject *obj, int buffertype, char order)
// Create a memoryview object to a contiguous chunk of memory (in either ‘C’ or ‘F’ortran order) from an object that defines the buffer interface. If memory is contiguous, the memoryview object points to the original memory. Otherwise copy is made and the memoryview points to a new bytes object.
func PyMemoryView_GetContiguous(obj *PyObject, buffertype int, order string) *PyObject {
	c_order := C.char(order[0])
	return togo(C.PyMemoryView_GetContiguous(topy(obj), C.int(buffertype), c_order))
}

// int PyMemoryView_Check(PyObject *obj)
// Return true if the object obj is a memoryview object. It is not currently allowed to create subclasses of memoryview.
func PyMemoryView_Check(obj *PyObject) bool {
	return int2bool(C._gopy_PyMemoryView_Check(topy(obj)))
}

// Py_buffer *PyMemoryView_GET_BUFFER(PyObject *obj)
// Return a pointer to the buffer-info structure wrapped by the given object. The object must be a memoryview instance; this macro doesn’t check its type, you must do it yourself or you will risk crashes.
func PyMemoryView_GET_BUFFER(obj *PyObject) *Py_buffer {
	buf := C._gopy_PyMemoryView_GET_BUFFER(topy(obj))
	return &Py_buffer{ptr: buf}
}

// EOF
