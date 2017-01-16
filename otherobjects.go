package python

// #include "go-python.h"
import "C"

import (
	"unsafe"
)

///// module /////

// int PyModule_Check(PyObject *p)
// Return true if p is a module object, or a subtype of a module object.
//
// Changed in version 2.2: Allowed subtypes to be accepted.
func PyModule_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyModule_Check(topy(self)))
}

// int PyModule_CheckExact(PyObject *p)
// Return true if p is a module object, but not a subtype of PyModule_Type.
//
// New in version 2.2.
func PyModule_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyModule_CheckExact(topy(self)))
}

// PyObject* PyModule_New(const char *name)
// Return value: New reference.
// Return a new module object with the __name__ attribute set to name. Only the module’s __doc__ and __name__ attributes are filled in; the caller is responsible for providing a __file__ attribute.
func PyModule_New(name string) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return togo(C.PyModule_New(c_name))
}

// PyObject* PyModule_GetDict(PyObject *module)
// Return value: Borrowed reference.
// Return the dictionary object that implements module‘s namespace; this object is the same as the __dict__ attribute of the module object. This function never fails. It is recommended extensions use other PyModule_*() and PyObject_*() functions rather than directly manipulate a module’s __dict__.
func PyModule_GetDict(self *PyObject) *PyObject {
	return togo(C.PyModule_GetDict(topy(self)))
}

// char* PyModule_GetName(PyObject *module)
// Return module‘s __name__ value. If the module does not provide one, or if it is not a string, SystemError is raised and NULL is returned.
func PyModule_GetName(self *PyObject) string {
	c_name := C.PyModule_GetName(topy(self))
	// we do not own c_name...
	return C.GoString(c_name)
}

// char* PyModule_GetFilename(PyObject *module)
// Return the name of the file from which module was loaded using module‘s __file__ attribute. If this is not defined, or if it is not a string, raise SystemError and return NULL.
func PyModule_GetFilename(self *PyObject) string {
	c_name := C.PyModule_GetFilename(topy(self))
	// we do not own c_name...
	return C.GoString(c_name)
}

// int PyModule_AddObject(PyObject *module, const char *name, PyObject *value)
// Add an object to module as name. This is a convenience function which can be used from the module’s initialization function. This steals a reference to value. Return -1 on error, 0 on success.
//
// New in version 2.0.
func PyModule_AddObject(self *PyObject, name string, value *PyObject) error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return int2err(C.PyModule_AddObject(topy(self), c_name, topy(value)))
}

// int PyModule_AddIntConstant(PyObject *module, const char *name, long value)
// Add an integer constant to module as name. This convenience function can be used from the module’s initialization function. Return -1 on error, 0 on success.
//
// New in version 2.0.
func PyModule_AddIntConstant(self *PyObject, name string, value int) error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return int2err(C.PyModule_AddIntConstant(topy(self), c_name, C.long(value)))
}

// int PyModule_AddStringConstant(PyObject *module, const char *name, const char *value)
// Add a string constant to module as name. This convenience function can be used from the module’s initialization function. The string value must be null-terminated. Return -1 on error, 0 on success.
//
// New in version 2.0.
func PyModule_AddStringConstant(self *PyObject, name, value string) error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	return int2err(C.PyModule_AddStringConstant(topy(self), c_name, c_value))
}

// int PyModule_AddIntMacro(PyObject *module, macro)
// Add an int constant to module. The name and the value are taken from macro. For example PyModule_AddConstant(module, AF_INET) adds the int constant AF_INET with the value of AF_INET to module. Return -1 on error, 0 on success.
//
// New in version 2.6.
func PyModule_AddIntMacro(self *PyObject, macro interface{}) error {
	//FIXME ?
	panic("not implemented")
}

// int PyModule_AddStringMacro(PyObject *module, macro)
// Add a string constant to module.
// New in version 2.6.
func PyModule_AddStringMacro(self *PyObject, macro interface{}) error {
	//FIXME ?
	panic("not implemented")
}

///// class /////

// int PyClass_Check(PyObject *o)
// Return true if the object o is a class object, including instances of types derived from the standard class object. Return false in all other cases.
func PyClass_Check(o *PyObject) bool {
	return int2bool(C._gopy_PyClass_Check(topy(o)))
}

// int PyClass_IsSubclass(PyObject *klass, PyObject *base)
// Return true if klass is a subclass of base. Return false in all other cases.
// There are very few functions specific to instance objects.
func PyClass_IsSubclass(klass, base *PyObject) bool {
	return int2bool(C.PyClass_IsSubclass(topy(klass), topy(base)))
}

///// instance /////

// int PyInstance_Check(PyObject *obj)
// Return true if obj is an instance.
func PyInstance_Check(obj *PyObject) bool {
	return int2bool(C._gopy_PyInstance_Check(topy(obj)))
}

// PyObject* PyInstance_New(PyObject *class, PyObject *arg, PyObject *kw)
// Return value: New reference.
// Create a new instance of a specific class. The parameters arg and kw are used as the positional and keyword parameters to the object’s constructor.
func PyInstance_New(class, arg, kw *PyObject) *PyObject {
	return togo(C.PyInstance_New(topy(class), topy(arg), topy(kw)))
}

// PyObject* PyInstance_NewRaw(PyObject *class, PyObject *dict)
// Return value: New reference.
// Create a new instance of a specific class without calling its constructor. class is the class of new object. The dict parameter will be used as the object’s __dict__; if NULL, a new dictionary will be created for the instance.
func PyInstance_NewRaw(class, dict *PyObject) *PyObject {
	return togo(C.PyInstance_NewRaw(topy(class), topy(dict)))
}

///// function /////

// int PyFunction_Check(PyObject *o)
// Return true if o is a function object (has type PyFunction_Type). The parameter must not be NULL.
func PyFunction_Check(o *PyObject) bool {
	return int2bool(C._gopy_PyFunction_Check(topy(o)))
}

// PyObject* PyFunction_New(PyObject *code, PyObject *globals)
// Return value: New reference.
// Return a new function object associated with the code object code. globals must be a dictionary with the global variables accessible to the function.
//
// The function’s docstring, name and __module__ are retrieved from the code object, the argument defaults and closure are set to NULL.
func PyFunction_New(code, globals *PyObject) *PyObject {
	return togo(C.PyFunction_New(topy(code), topy(globals)))
}

// PyObject* PyFunction_GetCode(PyObject *op)
// Return value: Borrowed reference.
// Return the code object associated with the function object op.
func PyFunction_GetCode(op *PyObject) *PyObject {
	return togo(C.PyFunction_GetCode(topy(op)))
}

// PyObject* PyFunction_GetGlobals(PyObject *op)
// Return value: Borrowed reference.
// Return the globals dictionary associated with the function object op.
func PyFunction_GetGlobals(op *PyObject) *PyObject {
	return togo(C.PyFunction_GetGlobals(topy(op)))
}

// PyObject* PyFunction_GetModule(PyObject *op)
// Return value: Borrowed reference.
// Return the __module__ attribute of the function object op. This is normally a string containing the module name, but can be set to any other object by Python code.
func PyFunction_GetModule(op *PyObject) *PyObject {
	return togo(C.PyFunction_GetModule(topy(op)))
}

// PyObject* PyFunction_GetDefaults(PyObject *op)
// Return value: Borrowed reference.
// Return the argument default values of the function object op. This can be a tuple of arguments or NULL.
func PyFunction_GetDefaults(op *PyObject) *PyObject {
	return togo(C.PyFunction_GetDefaults(topy(op)))
}

// int PyFunction_SetDefaults(PyObject *op, PyObject *defaults)
// Set the argument default values for the function object op. defaults must be Py_None or a tuple.
//
// Raises SystemError and returns -1 on failure.
func PyFunction_SetDefaults(op, defaults *PyObject) error {
	return int2err(C.PyFunction_SetDefaults(topy(op), topy(defaults)))
}

// PyObject* PyFunction_GetClosure(PyObject *op)
// Return value: Borrowed reference.
// Return the closure associated with the function object op. This can be NULL or a tuple of cell objects.
func PyFunction_GetClosure(op *PyObject) *PyObject {
	return togo(C.PyFunction_GetClosure(topy(op)))
}

// int PyFunction_SetClosure(PyObject *op, PyObject *closure)
// Set the closure associated with the function object op. closure must be Py_None or a tuple of cell objects.
//
// Raises SystemError and returns -1 on failure.
func PyFunction_SetClosure(op, closure *PyObject) error {
	return int2err(C.PyFunction_SetClosure(topy(op), topy(closure)))
}

///// method /////

// int PyMethod_Check(PyObject *o)
// Return true if o is a method object (has type PyMethod_Type). The parameter must not be NULL.
func PyMethod_Check(o *PyObject) bool {
	return int2bool(C._gopy_PyMethod_Check(topy(o)))
}

// PyObject* PyMethod_New(PyObject *func, PyObject *self, PyObject *class)
// Return value: New reference.
// Return a new method object, with func being any callable object; this is the function that will be called when the method is called. If this method should be bound to an instance, self should be the instance and class should be the class of self, otherwise self should be NULL and class should be the class which provides the unbound method..
func PyMethod_New(fct, self, class *PyObject) *PyObject {
	return togo(C.PyMethod_New(topy(fct), topy(self), topy(class)))
}

// PyObject* PyMethod_Class(PyObject *meth)
// Return value: Borrowed reference.
// Return the class object from which the method meth was created; if this was created from an instance, it will be the class of the instance.
func PyMethod_Class(meth *PyObject) *PyObject {
	return togo(C.PyMethod_Class(topy(meth)))
}

// PyObject* PyMethod_GET_CLASS(PyObject *meth)
// Return value: Borrowed reference.
// Macro version of PyMethod_Class() which avoids error checking.
func PyMethod_GET_CLASS(meth *PyObject) *PyObject {
	return togo(C._gopy_PyMethod_GET_CLASS(topy(meth)))
}

// PyObject* PyMethod_Function(PyObject *meth)
// Return value: Borrowed reference.
// Return the function object associated with the method meth.
func PyMethod_Function(meth *PyObject) *PyObject {
	return togo(C.PyMethod_Function(topy(meth)))
}

// PyObject* PyMethod_GET_FUNCTION(PyObject *meth)
// Return value: Borrowed reference.
// Macro version of PyMethod_Function() which avoids error checking.
func PyMethod_GET_FUNCTION(meth *PyObject) *PyObject {
	return togo(C._gopy_PyMethod_GET_FUNCTION(topy(meth)))
}

// PyObject* PyMethod_Self(PyObject *meth)
// Return value: Borrowed reference.
// Return the instance associated with the method meth if it is bound, otherwise return NULL.
func PyMethod_Self(meth *PyObject) *PyObject {
	return togo(C.PyMethod_Self(topy(meth)))
}

// PyObject* PyMethod_GET_SELF(PyObject *meth)
// Return value: Borrowed reference.
// Macro version of PyMethod_Self() which avoids error checking.
func PyMethod_GET_SELF(meth *PyObject) *PyObject {
	return togo(C._gopy_PyMethod_GET_SELF(topy(meth)))
}

// int PyMethod_ClearFreeList()
// Clear the free list. Return the total number of freed items.
//
// New in version 2.6.
func PyMethod_ClearFreeList() int {
	return int(C.PyMethod_ClearFreeList())
}

///// slice /////

type PySliceObject struct {
	ptr *C.PySliceObject
}

// int PySlice_Check(PyObject *ob)
// Return true if ob is a slice object; ob must not be NULL.
func PySlice_Check(ob *PyObject) bool {
	return int2bool(C._gopy_PySlice_Check(topy(ob)))
}

// PyObject* PySlice_New(PyObject *start, PyObject *stop, PyObject *step)
// Return value: New reference.
// Return a new slice object with the given values. The start, stop, and step parameters are used as the values of the slice object attributes of the same names. Any of the values may be NULL, in which case the None will be used for the corresponding attribute. Return NULL if the new object could not be allocated.
func PySlice_New(start, stop, step *PyObject) *PyObject {
	return togo(C.PySlice_New(topy(start), topy(stop), topy(step)))
}

// int PySlice_GetIndices(PySliceObject *slice, Py_ssize_t length, Py_ssize_t *start, Py_ssize_t *stop, Py_ssize_t *step)
// Retrieve the start, stop and step indices from the slice object slice, assuming a sequence of length length. Treats indices greater than length as errors.
//
// Returns 0 on success and -1 on error with no exception set (unless one of the indices was not None and failed to be converted to an integer, in which case -1 is returned with an exception set).
//
// You probably do not want to use this function. If you want to use slice objects in versions of Python prior to 2.3, you would probably do well to incorporate the source of PySlice_GetIndicesEx(), suitably renamed, in the source of your extension.
//
// Changed in version 2.5: This function used an int type for length and an int * type for start, stop, and step. This might require changes in your code for properly supporting 64-bit systems.
func PySlice_GetIndices(slice *PySliceObject, length int) (start, stop, step int, err error) {
	c_start := C.Py_ssize_t(0)
	c_stop := C.Py_ssize_t(0)
	c_step := C.Py_ssize_t(0)

	err = int2err(C.PySlice_GetIndices(slice.ptr, C.Py_ssize_t(length),
		&c_start, &c_stop, &c_step))

	start = int(c_start)
	stop = int(c_stop)
	step = int(c_step)

	return
}

// int PySlice_GetIndicesEx(PySliceObject *slice, Py_ssize_t length, Py_ssize_t *start, Py_ssize_t *stop, Py_ssize_t *step, Py_ssize_t *slicelength)
// Usable replacement for PySlice_GetIndices(). Retrieve the start, stop, and step indices from the slice object slice assuming a sequence of length length, and store the length of the slice in slicelength. Out of bounds indices are clipped in a manner consistent with the handling of normal slices.
//
// Returns 0 on success and -1 on error with exception set.
//
// New in version 2.3.
//
// Changed in version 2.5: This function used an int type for length and an int * type for start, stop, step, and slicelength. This might require changes in your code for properly supporting 64-bit systems.
func PySlice_GetIndicesEx(slice *PySliceObject, length int) (start, stop, step, slicelength int, err error) {

	c_start := C.Py_ssize_t(0)
	c_stop := C.Py_ssize_t(0)
	c_step := C.Py_ssize_t(0)
	c_slice := C.Py_ssize_t(0)

	err = int2err(C.PySlice_GetIndicesEx(slice.ptr, C.Py_ssize_t(length),
		&c_start, &c_stop, &c_step, &c_slice))

	start = int(c_start)
	stop = int(c_stop)
	step = int(c_step)
	slicelength = int(c_slice)

	return
}

///// capsule /////
type PyCapsule_Destructor func(*PyObject)

// int PyCapsule_CheckExact(PyObject *p)
// Return true if its argument is a PyCapsule.
func PyCapsule_CheckExact(p *PyObject) bool {
	return int2bool(C._gopy_PyCapsule_CheckExact(topy(p)))
}

// PyObject* PyCapsule_New(void *pointer, const char *name, PyCapsule_Destructor destructor)
// Return value: New reference.
// Create a PyCapsule encapsulating the pointer. The pointer argument may not be NULL.
//
// On failure, set an exception and return NULL.
//
// The name string may either be NULL or a pointer to a valid C string. If non-NULL, this string must outlive the capsule. (Though it is permitted to free it inside the destructor.)
//
// If the destructor argument is not NULL, it will be called with the capsule as its argument when it is destroyed.
//
// If this capsule will be stored as an attribute of a module, the name should be specified as modulename.attributename. This will enable other modules to import the capsule using PyCapsule_Import().
func PyCapsule_New(pointer *C.char, name string, dtor C.PyCapsule_Destructor) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//FIXME use a go PyCapsule_Destructor ?
	//FIXME use an interface{} instead of *C.char ?
	return togo(C.PyCapsule_New(unsafe.Pointer(pointer), c_name, dtor))
}

// void* PyCapsule_GetPointer(PyObject *capsule, const char *name)
// Retrieve the pointer stored in the capsule. On failure, set an exception and return NULL.
//
// The name parameter must compare exactly to the name stored in the capsule. If the name stored in the capsule is NULL, the name passed in must also be NULL. Python uses the C function strcmp() to compare capsule names.
func PyCapsule_GetPointer(capsule *PyObject, name string) *C.char {

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//FIXME use an interface{} instead of *C.char ?
	ptr := C.PyCapsule_GetPointer(topy(capsule), c_name)
	return (*C.char)(ptr)
}

// PyCapsule_Destructor PyCapsule_GetDestructor(PyObject *capsule)
// Return the current destructor stored in the capsule. On failure, set an exception and return NULL.
//
// It is legal for a capsule to have a NULL destructor. This makes a NULL return code somewhat ambiguous; use PyCapsule_IsValid() or PyErr_Occurred() to disambiguate.
func PyCapsule_GetDestructor(capsule *PyObject) C.PyCapsule_Destructor {
	return C.PyCapsule_GetDestructor(topy(capsule))
}

// void* PyCapsule_GetContext(PyObject *capsule)
// Return the current context stored in the capsule. On failure, set an exception and return NULL.
//
// It is legal for a capsule to have a NULL context. This makes a NULL return code somewhat ambiguous; use PyCapsule_IsValid() or PyErr_Occurred() to disambiguate.
func PyCapsule_GetContext(capsule *PyObject) *C.char {

	//FIXME use an interface{} instead of *C.char ?
	ptr := C.PyCapsule_GetContext(topy(capsule))
	return (*C.char)(ptr)
}

// const char* PyCapsule_GetName(PyObject *capsule)
// Return the current name stored in the capsule. On failure, set an exception and return NULL.
//
// It is legal for a capsule to have a NULL name. This makes a NULL return code somewhat ambiguous; use PyCapsule_IsValid() or PyErr_Occurred() to disambiguate.
func PyCapsule_GetName(capsule *PyObject) string {

	c_name := C.PyCapsule_GetName(topy(capsule))
	return C.GoString(c_name)
}

// void* PyCapsule_Import(const char *name, int no_block)
// Import a pointer to a C object from a capsule attribute in a module. The name parameter should specify the full name to the attribute, as in module.attribute. The name stored in the capsule must match this string exactly. If no_block is true, import the module without blocking (using PyImport_ImportModuleNoBlock()). If no_block is false, import the module conventionally (using PyImport_ImportModule()).
//
// Return the capsule’s internal pointer on success. On failure, set an exception and return NULL. However, if PyCapsule_Import() failed to import the module, and no_block was true, no exception is set.
func PyCapsule_Import(name string, no_block bool) *C.char {

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_no_block := C.int(0)
	if no_block {
		c_no_block = C.int(1)
	}

	//FIXME use an interface{} instead of *C.char ?
	ptr := C.PyCapsule_Import(c_name, c_no_block)
	return (*C.char)(ptr)
}

// int PyCapsule_IsValid(PyObject *capsule, const char *name)
// Determines whether or not capsule is a valid capsule. A valid capsule is non-NULL, passes PyCapsule_CheckExact(), has a non-NULL pointer stored in it, and its internal name matches the name parameter. (See PyCapsule_GetPointer() for information on how capsule names are compared.)
//
// In other words, if PyCapsule_IsValid() returns a true value, calls to any of the accessors (any function starting with PyCapsule_Get()) are guaranteed to succeed.
//
// Return a nonzero value if the object is valid and matches the name passed in. Return 0 otherwise. This function will not fail.
func PyCapsule_IsValid(capsule *PyObject, name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return int2bool(C.PyCapsule_IsValid(topy(capsule), c_name))
}

// int PyCapsule_SetContext(PyObject *capsule, void *context)
// Set the context pointer inside capsule to context.
//
// Return 0 on success. Return nonzero and set an exception on failure.
func PyCapsule_SetContext(capsule *PyObject, context *C.char) error {
	//FIXME use interface{} instead of *C.char ?
	return int2err(C.PyCapsule_SetContext(topy(capsule), unsafe.Pointer(context)))
}

// int PyCapsule_SetDestructor(PyObject *capsule, PyCapsule_Destructor destructor)
// Set the destructor inside capsule to destructor.
//
// Return 0 on success. Return nonzero and set an exception on failure.
func PyCapsule_SetDestructor(capsule *PyObject, dtor C.PyCapsule_Destructor) error {
	//FIXME use go-PyCapsule_Destructor instead of cgo one ?
	return int2err(C.PyCapsule_SetDestructor(topy(capsule), dtor))
}

// int PyCapsule_SetName(PyObject *capsule, const char *name)
// Set the name inside capsule to name. If non-NULL, the name must outlive the capsule. If the previous name stored in the capsule was not NULL, no attempt is made to free it.
//
// Return 0 on success. Return nonzero and set an exception on failure.
func PyCapsule_SetName(capsule *PyObject, name string) error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return int2err(C.PyCapsule_SetName(topy(capsule), c_name))
}

// int PyCapsule_SetPointer(PyObject *capsule, void *pointer)
// Set the void pointer inside capsule to pointer. The pointer may not be NULL.
//
// Return 0 on success. Return nonzero and set an exception on failure.
func PyCapsule_SetPointer(capsule *PyObject, pointer *C.char) error {
	//FIXME use interface{} instead of *C.char ?
	return int2err(C.PyCapsule_SetPointer(topy(capsule), unsafe.Pointer(pointer)))
}

///// generator /////

type PyFrameObject struct {
	ptr *C.PyFrameObject
}

// int PyGen_Check(PyObject *ob)
// Return true if ob is a generator object; ob must not be NULL.
func PyGen_Check(ob *PyObject) bool {
	return int2bool(C._gopy_PyGen_Check(topy(ob)))
}

// int PyGen_CheckExact(ob)
// Return true if ob‘s type is PyGen_Type is a generator object; ob must not be NULL.
func PyGen_CheckExact(ob *PyObject) bool {
	return int2bool(C._gopy_PyGen_CheckExact(topy(ob)))
}

// PyObject* PyGen_New(PyFrameObject *frame)
// Return value: New reference.
// Create and return a new generator object based on the frame object. A reference to frame is stolen by this function. The parameter must not be NULL.
func PyGen_New(frame *PyFrameObject) *PyObject {
	return togo(C.PyGen_New((*C.struct__frame)(frame.ptr)))
}

///// iterator /////

// int PySeqIter_Check(op)
// Return true if the type of op is PySeqIter_Type.
//
// New in version 2.2.
func PySeqIter_Check(op *PyObject) bool {
	return int2bool(C._gopy_PySeqIter_Check(topy(op)))
}

// PyObject* PyIter_Next(PyObject *o)
// Return value: New reference.
//
// Return the next value from the iteration o.  The object must be an iterator (it is up to the caller to check this).  If there are no remaining values, returns NULL with no exception set.  If an error occurs while retrieving the item, returns NULL and passes along the exception.
func PyIter_Next(op *PyObject) *PyObject {
  return togo(C.PyIter_Next(topy(op)))
}

// PyObject* PySeqIter_New(PyObject *seq)
// Return value: New reference.
// Return an iterator that works with a general sequence object, seq. The iteration ends when the sequence raises IndexError for the subscripting operation.
//
// New in version 2.2.
//
// PyTypeObject PyCallIter_Type
// Type object for iterator objects returned by PyCallIter_New() and the two-argument form of the iter() built-in function.
//
// New in version 2.2.
func PySeqIter_New(seq *PyObject) *PyObject {
	return togo(C.PySeqIter_New(topy(seq)))
}

// int PyCallIter_Check(op)
// Return true if the type of op is PyCallIter_Type.
//
// New in version 2.2.
func PyCallIter_Check(op *PyObject) bool {
	return int2bool(C._gopy_PyCallIter_Check(topy(op)))
}

// PyObject* PyCallIter_New(PyObject *callable, PyObject *sentinel)
// Return value: New reference.
// Return a new iterator. The first parameter, callable, can be any Python callable object that can be called with no parameters; each call to it should return the next item in the iteration. When callable returns a value equal to sentinel, the iteration will be terminated.
//
// New in version 2.2.
func PyCallIter_New(callable, sentinel *PyObject) *PyObject {
	return togo(C.PyCallIter_New(topy(callable), topy(sentinel)))
}

// PyCodeObject* PyCode_NewEmpty(char *filename, char *funcname, int firstlineno)
// Return value: New reference.
// Return an empty CodePythonObject that corresponds to the file name, func name and line number in the source file it points to.
func PyCode_NewEmpty(filename string, funcname string, firstlineno int) *PyObject {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	c_funcname := C.CString(funcname)
	defer C.free(unsafe.Pointer(c_funcname))

	o := C.PyCode_NewEmpty(c_filename, c_funcname, C.int(firstlineno))

	// need to (unsafe-ly) cast to *C.PyObject as o is a *C.PyCodeObject
	return togo((*C.PyObject)(unsafe.Pointer(o)))
}

// EOF
