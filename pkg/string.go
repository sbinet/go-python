package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
//int _gopy_PyString_Check(PyObject *o) { return PyString_Check(o); }
//Py_ssize_t _gopy_PyString_GET_SIZE(PyObject *o) { return PyString_GET_SIZE(o);}
//char* _gopy_PyString_AS_STRING(PyObject *o) { return PyString_AS_STRING(o); }
import "C"
import "unsafe"

type PyString PyObject

/*
int PyString_Check(PyObject *o)
Return true if the object o is a string object or an instance of a subtype of the string type.

Changed in version 2.2: Allowed subtypes to be accepted.
*/
func PyString_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyString_Check(self.ptr))
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
PyObject* PyString_FromStringAndSize(const char *v, Py_ssize_t len)
Return value: New reference.
Return a new string object with a copy of the string v as value and length len on success, and NULL on failure. If v is NULL, the contents of the string are uninitialized.

Changed in version 2.5: This function used an int type for len. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyString_FromStringAndSize(v string, sz int) *PyObject {
	cstr := C.CString(v)
	defer C.free(unsafe.Pointer(cstr))
	return togo(C.PyString_FromStringAndSize(cstr, C.Py_ssize_t(sz)))
}

/*
PyObject* PyString_FromFormat(const char *format, ...)
Return value: New reference.
Take a C printf()-style format string and a variable number of arguments, calculate the size of the resulting Python string and return a string with the values formatted into it. The variable arguments must be C types and must correspond exactly to the format characters in the format string. The following format characters are allowed:

Format Characters	Type	Comment
%%	n/a	The literal % character.
%c	int	A single character, represented as an C int.
%d	int	Exactly equivalent to printf("%d").
%u	unsigned int	Exactly equivalent to printf("%u").
%ld	long	Exactly equivalent to printf("%ld").
%lu	unsigned long	Exactly equivalent to printf("%lu").
%lld	long long	Exactly equivalent to printf("%lld").
%llu	unsigned long long	Exactly equivalent to printf("%llu").
%zd	Py_ssize_t	Exactly equivalent to printf("%zd").
%zu	size_t	Exactly equivalent to printf("%zu").
%i	int	Exactly equivalent to printf("%i").
%x	int	Exactly equivalent to printf("%x").
%s	char*	A null-terminated C character array.
%p	void*	The hex representation of a C pointer. Mostly equivalent to printf("%p") except that it is guaranteed to start with the literal 0x regardless of what the platform’s printf yields.
An unrecognized format character causes all the rest of the format string to be copied as-is to the result string, and any extra arguments discarded.

Note The “%lld” and “%llu” format specifiers are only available when HAVE_LONG_LONG is defined.
Changed in version 2.7: Support for “%lld” and “%llu” added.
*/
func PyString_FromFormat(format string, args ...interface{}) *PyObject {
	cfmt := C.CString(format)
	defer C.free(unsafe.Pointer(cfmt))

	// FIXME
	panic("not implemented")
	return nil
}

/*
PyObject* PyString_FromFormatV(const char *format, va_list vargs)
Return value: New reference.
Identical to PyString_FromFormat() except that it takes exactly two arguments.
*/
func PyStringFromFormatV(format string, args ...interface{}) *PyObject {
	cfmt := C.CString(format)
	defer C.free(unsafe.Pointer(cfmt))

	// FIXME
	panic("not implemented")
	return nil
}

/*
Py_ssize_t PyString_Size(PyObject *string)
Return the length of the string in string object string.

Changed in version 2.5: This function returned an int type. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyString_Size(self *PyObject) int {
	sz := C.PyString_Size(topy(self))
	return int(sz)
}

/*
Py_ssize_t PyString_GET_SIZE(PyObject *string)
Macro form of PyString_Size() but without error checking.

Changed in version 2.5: This macro returned an int type. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyString_GET_SIZE(self *PyObject) int {
	sz := C._gopy_PyString_GET_SIZE(topy(self))
	return int(sz)
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

/*
char* PyString_AS_STRING(PyObject *string)
Macro form of PyString_AsString() but without error checking. Only string objects are supported; no Unicode objects should be passed.
*/
func PyString_AS_STRING(self *PyObject) string {
	c_str := C._gopy_PyString_AS_STRING(self.ptr)
	// we dont own c_str...
	//defer C.free(unsafe.Pointer(c_str))
	return C.GoString(c_str)
}

/*
int PyString_AsStringAndSize(PyObject *obj, char **buffer, Py_ssize_t *length)
Return a NUL-terminated representation of the contents of the object obj through the output variables buffer and length.

The function accepts both string and Unicode objects as input. For Unicode objects it returns the default encoded version of the object. If length is NULL, the resulting buffer may not contain NUL characters; if it does, the function returns -1 and a TypeError is raised.

The buffer refers to an internal string buffer of obj, not a copy. The data must not be modified in any way, unless the string was just created using PyString_FromStringAndSize(NULL, size). It must not be deallocated. If string is a Unicode object, this function computes the default encoding of string and operates on that. If string is not a string object at all, PyString_AsStringAndSize() returns -1 and raises TypeError.

Changed in version 2.5: This function used an int * type for length. This might require changes in your code for properly supporting 64-bit systems.

*/
func PyString_AsStringAndSize(self *PyObject, sz int) (str string, err int) {
	// FIXME
	panic("not implemented")
}

/*
void PyString_Concat(PyObject **string, PyObject *newpart)
Create a new string object in *string containing the contents of newpart appended to string; the caller will own the new reference. The reference to the old value of string will be stolen. If the new string cannot be created, the old reference to string will still be discarded and the value of *string will be set to NULL; the appropriate exception will be set.
*/
func PyString_Concat(self, newpart *PyObject) *PyObject {
	// FIXME
	panic("not implemented")
}

/*
void PyString_ConcatAndDel(PyObject **string, PyObject *newpart)
Create a new string object in *string containing the contents of newpart appended to string. This version decrements the reference count of newpart.
int _PyString_Resize(PyObject **string, Py_ssize_t newsize)
A way to resize a string object even though it is “immutable”. Only use this to build up a brand new string object; don’t use this if the string may already be known in other parts of the code. It is an error to call this function if the refcount on the input string object is not one. Pass the address of an existing string object as an lvalue (it may be written into), and the new size desired. On success, *string holds the resized string object and 0 is returned; the address in *string may differ from its input value. If the reallocation fails, the original string object at *string is deallocated, *string is set to NULL, a memory exception is set, and -1 is returned.

Changed in version 2.5: This function used an int type for newsize. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyString_ConcatAndDel(self, newpart *PyObject) *PyObject {
	// FIXME
	panic("not implemented")
}

/*
PyObject* PyString_Format(PyObject *format, PyObject *args)
Return value: New reference.
Return a new string object from format and args. Analogous to format % args. The args argument must be a tuple.
*/
func PyString_Format(format, args *PyObject) *PyObject {
	py_format := topy(format)
	py_args := topy(args)
	return togo(C.PyString_Format(py_format, py_args))
}

/*
void PyString_InternInPlace(PyObject **string)
Intern the argument *string in place. The argument must be the address of a pointer variable pointing to a Python string object. If there is an existing interned string that is the same as *string, it sets *string to it (decrementing the reference count of the old string object and incrementing the reference count of the interned string object), otherwise it leaves *string alone and interns it (incrementing its reference count). (Clarification: even though there is a lot of talk about reference counts, think of this function as reference-count-neutral; you own the object after the call if and only if you owned it before the call.)

Note This function is not available in 3.x and does not have a PyBytes alias.
*/
func PyString_InternInPlance(self *PyObject) {
	//FIXME check everything is OK...
	s := topy(self)
	C.PyString_InternInPlace(&s)
}

/*
PyObject* PyString_InternFromString(const char *v)
Return value: New reference.
A combination of PyString_FromString() and PyString_InternInPlace(), returning either a new string object that has been interned, or a new (“owned”) reference to an earlier interned string object with the same value.
*/
func PyString_InternFromString(v string) *PyObject {
	cstr := C.CString(v)
	defer C.free(unsafe.Pointer(cstr))

	return togo(C.PyString_InternFromString(cstr))
}

/*
Note This function is not available in 3.x and does not have a PyBytes alias.
PyObject* PyString_Decode(const char *s, Py_ssize_t size, const char *encoding, const char *errors)
Return value: New reference.
Create an object by decoding size bytes of the encoded buffer s using the codec registered for encoding. encoding and errors have the same meaning as the parameters of the same name in the unicode() built-in function. The codec to be used is looked up using the Python codec registry. Return NULL if an exception was raised by the codec.

Note This function is not available in 3.x and does not have a PyBytes alias.
Changed in version 2.5: This function used an int type for size. This might require changes in your code for properly supporting 64-bit systems.
*/
func PyString_Decode(s string, sz int, encoding, errors string) *PyObject {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))

	c_encoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(c_encoding))

	c_errors := C.CString(errors)
	defer C.free(unsafe.Pointer(c_errors))

	return togo(C.PyString_Decode(c_s, C.Py_ssize_t(sz), c_encoding, c_errors))
}

/*
PyObject* PyString_AsDecodedObject(PyObject *str, const char *encoding, const char *errors)
Return value: New reference.
Decode a string object by passing it to the codec registered for encoding and return the result as Python object. encoding and errors have the same meaning as the parameters of the same name in the string encode() method. The codec to be used is looked up using the Python codec registry. Return NULL if an exception was raised by the codec.

Note This function is not available in 3.x and does not have a PyBytes alias.
*/
func PyString_AsDecodedObject(self *PyObject, encoding, errors string) *PyObject {
	c_encoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(c_encoding))

	c_errors := C.CString(errors)
	defer C.free(unsafe.Pointer(c_errors))

	return togo(C.PyString_AsDecodedObject(topy(self), c_encoding, c_errors))
}

/*
PyObject* PyString_Encode(const char *s, Py_ssize_t size, const char *encoding, const char *errors)
Return value: New reference.
Encode the char buffer of the given size by passing it to the codec registered for encoding and return a Python object. encoding and errors have the same meaning as the parameters of the same name in the string encode() method. The codec to be used is looked up using the Python codec registry. Return NULL if an exception was raised by the codec.

Note This function is not available in 3.x and does not have a PyBytes alias.
Changed in version 2.5: This function used an int type for size. This might require changes in your code for properly supporting 64-bit systems.
*/
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

/*
PyObject* PyString_AsEncodedObject(PyObject *str, const char *encoding, const char *errors)
Return value: New reference.
Encode a string object using the codec registered for encoding and return the result as Python object. encoding and errors have the same meaning as the parameters of the same name in the string encode() method. The codec to be used is looked up using the Python codec registry. Return NULL if an exception was raised by the codec.

Note This function is not available in 3.x and does not have a PyBytes alias.
*/
func PyString_AsEncodedObject(self *PyObject, encoding, errors string) *PyObject {
	c_encoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(c_encoding))

	c_errors := C.CString(errors)
	defer C.free(unsafe.Pointer(c_errors))

	return togo(C.PyString_AsEncodedObject(topy(self), c_encoding, c_errors))
}
// EOF
