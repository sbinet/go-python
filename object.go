package python

//#include "go-python.h"
import "C"

import (
	"fmt"
	"os"
	"strings"
	"unsafe"
)

// PyObject layer
type PyObject struct {
	ptr *C.PyObject
}

// String returns a string representation of the PyObject
func (self *PyObject) String() string {
	o := self.Str()
	defer o.DecRef()
	return PyString_AsString(o)
}

func (self *PyObject) topy() *C.PyObject {
	return self.ptr
}

func topy(self *PyObject) *C.PyObject {
	if self == nil {
		return nil
	}
	return self.ptr
}

func togo(obj *C.PyObject) *PyObject {
	switch obj {
	case nil:
		return nil
	case Py_None.ptr:
		return Py_None
	case Py_True.ptr:
		return Py_True
	case Py_False.ptr:
		return Py_False
	default:
		return &PyObject{ptr: obj}
	}
}

// PyObject_FromVoidPtr converts a PyObject from an unsafe.Pointer
func PyObject_FromVoidPtr(ptr unsafe.Pointer) *PyObject {
	return togo((*C.PyObject)(ptr))
}

func int2bool(i C.int) bool {
	switch i {
	case -1:
		return false
	case 0:
		return false
	case 1:
		return true
	default:
		return true
	}
	return false
}

func long2bool(i C.long) bool {
	switch i {
	case -1:
		return false
	case 0:
		return false
	case 1:
		return true
	default:
		return true
	}
	return false
}

func bool2int(i bool) C.int {
	if i {
		return C.int(1)
	}
	return C.int(0)
}

type gopy_err struct {
	err string
}

func (self *gopy_err) Error() string {
	return self.err
}

func int2err(i C.int) error {
	if i == 0 {
		return nil
	}
	//FIXME: also handle python exceptions ?
	return &gopy_err{fmt.Sprintf("error in C-Python (rc=%d)", int(i))}
}

func file2go(f *C.FILE) *os.File {
	return nil
}

// C.PyObject* PyObject_GetCPointer(PyObject *o)
// Returns the internal C pointer to CPython object.
func (self *PyObject) GetCPointer() *C.PyObject {
	return self.ptr
}

// void Py_IncRef(PyObject *o)
// Increment the reference count for object o. The object may be
// NULL, in which case the function has no effect.
func (self *PyObject) IncRef() {
	C.Py_IncRef(self.ptr)
}

// void Py_DecRef(PyObject *o)
// Decrement the reference count for object o. If the object is
// NULL, nothing happens. If the reference count reaches zero, the
// object’s type’s deallocation function (which must not be NULL) is
// invoked.
// WARNING: The deallocation function can cause arbitrary Python
// code to be invoked. See the warnings and instructions in the
// Python docs, and consider using Clear instead.
func (self *PyObject) DecRef() {
	C.Py_DecRef(self.ptr)
}

// void Py_CLEAR(PyObject *o)
// Clear sets the PyObject's internal pointer to nil
// before calling Py_DecRef. This avoids the potential issues with
// Python code called by the deallocator referencing invalid,
// partially-deallocated data.
func (self *PyObject) Clear() {
	tmp := self.ptr
	self.ptr = nil
	C.Py_DecRef(tmp)
}

// int PyObject_HasAttr(PyObject *o, PyObject *attr_name)
// Returns 1 if o has the attribute attr_name, and 0 otherwise. This is equivalent to the Python expression hasattr(o, attr_name). This function always succeeds.
func (self *PyObject) HasAttr(attr_name *PyObject) int {
	return int(C.PyObject_HasAttr(self.ptr, attr_name.ptr))
}

// int PyObject_HasAttrString(PyObject *o, const char *attr_name)
// Returns 1 if o has the attribute attr_name, and 0 otherwise. This is equivalent to the Python expression hasattr(o, attr_name). This function always succeeds.
func (self *PyObject) HasAttrString(attr_name string) int {
	c_attr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(c_attr_name))

	return int(C.PyObject_HasAttrString(self.ptr, c_attr_name))
}

// PyObject* PyObject_GetAttr(PyObject *o, PyObject *attr_name)
// Return value: New reference.
// Retrieve an attribute named attr_name from object o. Returns the attribute value on success, or NULL on failure. This is the equivalent of the Python expression o.attr_name.
func (self *PyObject) GetAttr(attr_name *PyObject) *PyObject {
	return togo(C.PyObject_GetAttr(self.ptr, attr_name.ptr))
}

// PyObject* PyObject_Dir()
// Return value: New reference.
// This is equivalent to the Python expression dir(o), returning a (possibly empty) list of strings appropriate for the object argument, or NULL if there was an error. If the argument is NULL, this is like the Python dir(), returning the names of the current locals; in this case, if no execution frame is active then NULL is returned but PyErr_Occurred() will return false.
func (self *PyObject) PyObject_Dir() *PyObject {
	return togo(C.PyObject_Dir(self.ptr))

}

// PyObject* PyObject_GetAttrString(PyObject *o, const char *attr_name)
// Return value: New reference.
// Retrieve an attribute named attr_name from object o. Returns the attribute value on success, or NULL on failure. This is the equivalent of the Python expression o.attr_name.
func (self *PyObject) GetAttrString(attr_name string) *PyObject {
	c_attr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(c_attr_name))
	return togo(C.PyObject_GetAttrString(self.ptr, c_attr_name))
}

// PyObject* PyObject_GenericGetAttr(PyObject *o, PyObject *name)
// Generic attribute getter function that is meant to be put into a type object’s tp_getattro slot. It looks for a descriptor in the dictionary of classes in the object’s MRO as well as an attribute in the object’s __dict__ (if present). As outlined in Implementing Descriptors, data descriptors take preference over instance attributes, while non-data descriptors don’t. Otherwise, an AttributeError is raised.
func (self *PyObject) GenericGetAttr(name *PyObject) *PyObject {
	return togo(C.PyObject_GenericGetAttr(self.ptr, name.ptr))
}

// int PyObject_SetAttr(PyObject *o, PyObject *attr_name, PyObject *v)
// Set the value of the attribute named attr_name, for object o, to the value v. Returns -1 on failure. This is the equivalent of the Python statement o.attr_name = v.
func (self *PyObject) SetAttr(attr_name, v *PyObject) int {
	return int(C.PyObject_SetAttr(self.ptr, attr_name.ptr, v.ptr))
}

// int PyObject_SetAttrString(PyObject *o, const char *attr_name, PyObject *v)
// Set the value of the attribute named attr_name, for object o, to the value v. Returns -1 on failure. This is the equivalent of the Python statement o.attr_name = v.
func (self *PyObject) SetAttrString(attr_name string, v *PyObject) int {
	c_attr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(c_attr_name))
	return int(C.PyObject_SetAttrString(self.ptr, c_attr_name, v.ptr))
}

// int PyObject_GenericSetAttr(PyObject *o, PyObject *name, PyObject *value)
// Generic attribute setter function that is meant to be put into a type object’s tp_setattro slot. It looks for a data descriptor in the dictionary of classes in the object’s MRO, and if found it takes preference over setting the attribute in the instance dictionary. Otherwise, the attribute is set in the object’s __dict__ (if present). Otherwise, an AttributeError is raised and -1 is returned.
func (self *PyObject) GenericSetAttr(name, value *PyObject) int {
	return int(C.PyObject_GenericSetAttr(self.ptr, name.ptr, value.ptr))
}

// int PyObject_DelAttr(PyObject *o, PyObject *attr_name)
// Delete attribute named attr_name, for object o. Returns -1 on failure. This is the equivalent of the Python statement del o.attr_name.
func (self *PyObject) DelAttr(attr_name *PyObject) int {
	return int(C._gopy_PyObject_DelAttr(self.ptr, attr_name.ptr))
}

// int PyObject_DelAttrString(PyObject *o, const char *attr_name)
// Delete attribute named attr_name, for object o. Returns -1 on failure. This is the equivalent of the Python statement del o.attr_name.
func (self *PyObject) DelAttrString(attr_name string) int {
	c_attr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(c_attr_name))
	return int(C._gopy_PyObject_DelAttrString(self.ptr, c_attr_name))
}

type Py_OPID C.int

const (
	Py_LT Py_OPID = C.Py_LT
	Py_LE Py_OPID = C.Py_LE
	Py_EQ Py_OPID = C.Py_EQ
	Py_NE Py_OPID = C.Py_NE
	Py_GT Py_OPID = C.Py_GT
	Py_GE Py_OPID = C.Py_GE
)

// PyObject* PyObject_RichCompare(PyObject *o1, PyObject *o2, int opid)
// Return value: New reference.
// Compare the values of o1 and o2 using the operation specified by opid, which must be one of Py_LT, Py_LE, Py_EQ, Py_NE, Py_GT, or Py_GE, corresponding to <, <=, ==, !=, >, or >= respectively. This is the equivalent of the Python expression o1 op o2, where op is the operator corresponding to opid. Returns the value of the comparison on success, or NULL on failure.
func (self *PyObject) RichCompare(o2 *PyObject, opid Py_OPID) *PyObject {
	return togo(C.PyObject_RichCompare(self.ptr, o2.ptr, C.int(opid)))
}

// int PyObject_RichCompareBool(PyObject *o1, PyObject *o2, int opid)
// Compare the values of o1 and o2 using the operation specified by opid, which must be one of Py_LT, Py_LE, Py_EQ, Py_NE, Py_GT, or Py_GE, corresponding to <, <=, ==, !=, >, or >= respectively. Returns -1 on error, 0 if the result is false, 1 otherwise. This is the equivalent of the Python expression o1 op o2, where op is the operator corresponding to opid.
func (self *PyObject) RichCompareBool(o2 *PyObject, opid Py_OPID) int {
	return int(C.PyObject_RichCompareBool(self.ptr, o2.ptr, C.int(opid)))
}

// int PyObject_Cmp(PyObject *o1, PyObject *o2, int *result)
// Compare the values of o1 and o2 using a routine provided by o1, if one exists, otherwise with a routine provided by o2. The result of the comparison is returned in result. Returns -1 on failure. This is the equivalent of the Python statement result = cmp(o1, o2).
func (self *PyObject) Cmp(o2 *PyObject) (err, result int) {
	var c_result C.int = -1
	var c_err C.int = -1

	c_err = C.PyObject_Cmp(self.ptr, o2.ptr, &c_result)
	return int(c_err), int(c_result)
}

// int PyObject_Compare(PyObject *o1, PyObject *o2)
// Compare the values of o1 and o2 using a routine provided by o1, if one exists, otherwise with a routine provided by o2. Returns the result of the comparison on success. On error, the value returned is undefined; use PyErr_Occurred() to detect an error. This is equivalent to the Python expression cmp(o1, o2).
func (self *PyObject) Compare(o2 *PyObject) int {
	return int(C.PyObject_Compare(self.ptr, o2.ptr))
}

// PyObject* PyObject_Repr(PyObject *o)
// Return value: New reference.
// Compute a string representation of object o. Returns the string representation on success, NULL on failure. This is the equivalent of the Python expression repr(o). Called by the repr() built-in function and by reverse quotes.
func (self *PyObject) Repr() *PyObject {
	return togo(C.PyObject_Repr(self.ptr))
}

// PyObject* PyObject_Str(PyObject *o)
// Return value: New reference.
// Compute a string representation of object o. Returns the string representation on success, NULL on failure. This is the equivalent of the Python expression str(o). Called by the str() built-in function and by the print statement.
func (self *PyObject) Str() *PyObject {
	return togo(C.PyObject_Str(self.ptr))
}

// PyObject* PyObject_Bytes(PyObject *o)
// Compute a bytes representation of object o. In 2.x, this is just a alias for PyObject_Str().
func (self *PyObject) Bytes() *PyObject {
	return togo(C.PyObject_Bytes(self.ptr))
}

// PyObject* PyObject_Unicode(PyObject *o)
// Return value: New reference.
// Compute a Unicode string representation of object o. Returns the Unicode string representation on success, NULL on failure. This is the equivalent of the Python expression unicode(o). Called by the unicode() built-in function.
func (self *PyObject) Unicode() *PyObject {
	return togo(C.PyObject_Unicode(self.ptr))
}

// int PyObject_IsInstance(PyObject *inst, PyObject *cls)
// Returns 1 if inst is an instance of the class cls or a subclass of cls, or 0 if not. On error, returns -1 and sets an exception. If cls is a type object rather than a class object, PyObject_IsInstance() returns 1 if inst is of type cls. If cls is a tuple, the check will be done against every entry in cls. The result will be 1 when at least one of the checks returns 1, otherwise it will be 0. If inst is not a class instance and cls is neither a type object, nor a class object, nor a tuple, inst must have a __class__ attribute — the class relationship of the value of that attribute with cls will be used to determine the result of this function.
//
// New in version 2.1.
//
// Changed in version 2.2: Support for a tuple as the second argument added.
//
// Subclass determination is done in a fairly straightforward way, but includes a wrinkle that implementors of extensions to the class system may want to be aware of. If A and B are class objects, B is a subclass of A if it inherits from A either directly or indirectly. If either is not a class object, a more general mechanism is used to determine the class relationship of the two objects. When testing if B is a subclass of A, if A is B, PyObject_IsSubclass() returns true. If A and B are different objects, B‘s __bases__ attribute is searched in a depth-first fashion for A — the presence of the __bases__ attribute is considered sufficient for this determination.
func (self *PyObject) IsInstance(cls *PyObject) int {
	return int(C.PyObject_IsInstance(self.ptr, cls.ptr))
}

// int PyObject_IsSubclass(PyObject *derived, PyObject *cls)
// Returns 1 if the class derived is identical to or derived from the class cls, otherwise returns 0. In case of an error, returns -1. If cls is a tuple, the check will be done against every entry in cls. The result will be 1 when at least one of the checks returns 1, otherwise it will be 0. If either derived or cls is not an actual class object (or tuple), this function uses the generic algorithm described above.
//
// New in version 2.1.
//
// Changed in version 2.3: Older versions of Python did not support a tuple as the second argument.
func (self *PyObject) IsSubclass(cls *PyObject) int {
	return int(C.PyObject_IsSubclass(self.ptr, cls.ptr))
}

// int PyCallable_Check(PyObject *o)
// Determine if the object o is callable. Return 1 if the object is callable and 0 otherwise. This function always succeeds.
// PyObject* PyObject_Call(PyObject *callable_object, PyObject *args, PyObject *kw)
// Return value: New reference.
// Call a callable Python object callable_object, with arguments given by the tuple args, and named arguments given by the dictionary kw. If no named arguments are needed, kw may be NULL. args must not be NULL, use an empty tuple if no arguments are needed. Returns the result of the call on success, or NULL on failure. This is the equivalent of the Python expression apply(callable_object, args, kw) or callable_object(*args, **kw).
//
// New in version 2.2.
func (self *PyObject) Check_Callable() bool {
	return int2bool(C.PyCallable_Check(self.ptr))
}

// PyObject* PyObject_Call(PyObject *callable_object, PyObject *args, PyObject *kw)
// Return value: New reference.
// Call a callable Python object callable_object, with arguments given by the tuple args, and named arguments given by the dictionary kw. If no named arguments are needed, kw may be NULL. args must not be NULL, use an empty tuple if no arguments are needed. Returns the result of the call on success, or NULL on failure. This is the equivalent of the Python expression apply(callable_object, args, kw) or callable_object(*args, **kw).
func (self *PyObject) Call(args, kw *PyObject) *PyObject {
	return togo(C.PyObject_Call(self.ptr, args.ptr, kw.ptr))
}

// PyObject* PyObject_CallObject(PyObject *callable_object, PyObject *args)
// Return value: New reference.
// Call a callable Python object callable_object, with arguments given by the tuple args. If no arguments are needed, then args may be NULL. Returns the result of the call on success, or NULL on failure. This is the equivalent of the Python expression apply(callable_object, args) or callable_object(*args).
func (self *PyObject) CallObject(args *PyObject) *PyObject {
	return togo(C.PyObject_CallObject(self.ptr, args.ptr))
}

// PyObject* PyObject_CallFunction(PyObject *callable, char *format, ...)
// Return value: New reference.
// Call a callable Python object callable, with a variable number of C arguments. The C arguments are described using a Py_BuildValue() style format string. The format may be NULL, indicating that no arguments are provided. Returns the result of the call on success, or NULL on failure. This is the equivalent of the Python expression apply(callable, args) or callable(*args). Note that if you only pass PyObject * args, PyObject_CallFunctionObjArgs() is a faster alternative.
func (self *PyObject) CallFunction(args ...interface{}) *PyObject {
	if len(args) > int(C._gopy_max_varargs) {
		panic(fmt.Errorf(
			"gopy: maximum number of varargs (%d) exceeded (%d)",
			int(C._gopy_max_varargs),
			len(args),
		))
	}

	types := make([]string, 0, len(args))
	cargs := make([]unsafe.Pointer, 0, len(args))

	for _, arg := range args {
		ptr, typ := pyfmt(arg)
		types = append(types, typ)
		cargs = append(cargs, ptr)
		if typ == "s" {
			defer func(ptr unsafe.Pointer) {
				C.free(ptr)
			}(ptr)
		}
	}

	if len(args) <= 0 {
		o := C._gopy_PyObject_CallFunction(self.ptr, 0, nil, nil)
		return togo(o)
	}

	pyfmt := C.CString(strings.Join(types, ""))
	defer C.free(unsafe.Pointer(pyfmt))
	o := C._gopy_PyObject_CallFunction(
		self.ptr,
		C.int(len(args)),
		pyfmt,
		unsafe.Pointer(&cargs[0]),
	)

	return togo(o)

}

// PyObject* PyObject_CallMethod(PyObject *o, char *method, char *format, ...)
// Return value: New reference.
// Call the method named method of object o with a variable number of C arguments. The C arguments are described by a Py_BuildValue() format string that should produce a tuple. The format may be NULL, indicating that no arguments are provided. Returns the result of the call on success, or NULL on failure. This is the equivalent of the Python expression o.method(args). Note that if you only pass PyObject * args, PyObject_CallMethodObjArgs() is a faster alternative.
func (self *PyObject) CallMethod(method string, args ...interface{}) *PyObject {
	if len(args) > int(C._gopy_max_varargs) {
		panic(fmt.Errorf(
			"gopy: maximum number of varargs (%d) exceeded (%d)",
			int(C._gopy_max_varargs),
			len(args),
		))
	}

	cmethod := C.CString(method)
	defer C.free(unsafe.Pointer(cmethod))

	types := make([]string, 0, len(args))
	cargs := make([]unsafe.Pointer, 0, len(args))

	for _, arg := range args {
		ptr, typ := pyfmt(arg)
		types = append(types, typ)
		cargs = append(cargs, ptr)
		if typ == "s" {
			defer func(ptr unsafe.Pointer) {
				C.free(ptr)
			}(ptr)
		}
	}

	if len(args) <= 0 {
		o := C._gopy_PyObject_CallMethod(self.ptr, cmethod, 0, nil, nil)
		return togo(o)
	}

	pyfmt := C.CString(strings.Join(types, ""))
	defer C.free(unsafe.Pointer(pyfmt))
	o := C._gopy_PyObject_CallMethod(
		self.ptr,
		cmethod,
		C.int(len(args)),
		pyfmt,
		unsafe.Pointer(&cargs[0]),
	)

	return togo(o)
}

/*
PyObject* PyObject_CallFunctionObjArgs(PyObject *callable, ..., NULL)
Return value: New reference.
Call a callable Python object callable, with a variable number of PyObject* arguments. The arguments are provided as a variable number of parameters followed by NULL. Returns the result of the call on success, or NULL on failure.

New in version 2.2.
*/
func (self *PyObject) CallFunctionObjArgs(format string, args ...interface{}) *PyObject {
	return self.CallFunction(args...)
}

/*
PyObject* PyObject_CallMethodObjArgs(PyObject *o, PyObject *name, ..., NULL)
Return value: New reference.
Calls a method of the object o, where the name of the method is given as a Python string object in name. It is called with a variable number of PyObject* arguments. The arguments are provided as a variable number of parameters followed by NULL. Returns the result of the call on success, or NULL on failure.

New in version 2.2.
*/
func (self *PyObject) CallMethodObjArgs(method string, args ...interface{}) *PyObject {
	return self.CallMethod(method, args...)
}

// long PyObject_Hash(PyObject *o)
// Compute and return the hash value of an object o. On failure, return -1. This is the equivalent of the Python expression hash(o).
func (self *PyObject) Hash() int64 {
	return int64(C.PyObject_Hash(topy(self)))
}

// long PyObject_HashNotImplemented(PyObject *o)
// Set a TypeError indicating that type(o) is not hashable and return -1. This function receives special treatment when stored in a tp_hash slot, allowing a type to explicitly indicate to the interpreter that it is not hashable.
//
// New in version 2.6.
func (self *PyObject) HashNotImplemented() bool {
	return long2bool(C.PyObject_HashNotImplemented(topy(self)))
}

// int PyObject_IsTrue(PyObject *o)
// Returns 1 if the object o is considered to be true, and 0 otherwise. This is equivalent to the Python expression not not o. On failure, return -1.
func (self *PyObject) IsTrue() bool {
	return int2bool(C.PyObject_IsTrue(topy(self)))
}

// int PyObject_Not(PyObject *o)
// Returns 0 if the object o is considered to be true, and 1 otherwise. This is equivalent to the Python expression not o. On failure, return -1.
func (self *PyObject) Not() bool {
	return int2bool(C.PyObject_Not(topy(self)))
}

// PyObject* PyObject_Type(PyObject *o)
// Return value: New reference.
// When o is non-NULL, returns a type object corresponding to the object type of object o. On failure, raises SystemError and returns NULL. This is equivalent to the Python expression type(o). This function increments the reference count of the return value. There’s really no reason to use this function instead of the common expression o->ob_type, which returns a pointer of type PyTypeObject*, except when the incremented reference count is needed.
func (self *PyObject) Type() *PyObject {
	return togo(C.PyObject_Type(topy(self)))
}

// EOF
