package python

/*
#include "Python.h"
#include <stdlib.h>
#include <string.h>

int _gopy_PyModule_Check(PyObject *p) { return PyModule_Check(p); }
int _gopy_PyModule_CheckExact(PyObject *p) { return PyModule_CheckExact(p); }

 int _gopy_PyClass_Check(PyObject *o) { return PyClass_Check(o); }

 int _gopy_PyInstance_Check(PyObject *obj) { return PyInstance_Check(obj); }

*/
import "C"
import "unsafe"
import "os"

///// module /////

/*
int PyModule_Check(PyObject *p)
Return true if p is a module object, or a subtype of a module object.

Changed in version 2.2: Allowed subtypes to be accepted.
*/
func PyModule_Check(self *PyObject) bool {
	return int2bool(C._gopy_PyModule_Check(topy(self)))
}

/*
int PyModule_CheckExact(PyObject *p)
Return true if p is a module object, but not a subtype of PyModule_Type.

New in version 2.2.
*/
func PyModule_CheckExact(self *PyObject) bool {
	return int2bool(C._gopy_PyModule_CheckExact(topy(self)))
}

/*
PyObject* PyModule_New(const char *name)
Return value: New reference.
Return a new module object with the __name__ attribute set to name. Only the module’s __doc__ and __name__ attributes are filled in; the caller is responsible for providing a __file__ attribute.
*/
func PyModule_New(name string) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return togo(C.PyModule_New(c_name))
}

/*
PyObject* PyModule_GetDict(PyObject *module)
Return value: Borrowed reference.
Return the dictionary object that implements module‘s namespace; this object is the same as the __dict__ attribute of the module object. This function never fails. It is recommended extensions use other PyModule_*() and PyObject_*() functions rather than directly manipulate a module’s __dict__.
*/
func PyModule_GetDict(self *PyObject) *PyObject {
	return togo(C.PyModule_GetDict(topy(self)))
}

/*
char* PyModule_GetName(PyObject *module)
Return module‘s __name__ value. If the module does not provide one, or if it is not a string, SystemError is raised and NULL is returned.
*/
func PyModule_GetName(self *PyObject) string {
	c_name := C.PyModule_GetName(topy(self))
	// we do not own c_name...
	return C.GoString(c_name)
}

/*
char* PyModule_GetFilename(PyObject *module)
Return the name of the file from which module was loaded using module‘s __file__ attribute. If this is not defined, or if it is not a string, raise SystemError and return NULL.
*/
func PyModule_GetFilename(self *PyObject) string {
	c_name := C.PyModule_GetFilename(topy(self))
	// we do not own c_name...
	return C.GoString(c_name)
}

/*
int PyModule_AddObject(PyObject *module, const char *name, PyObject *value)
Add an object to module as name. This is a convenience function which can be used from the module’s initialization function. This steals a reference to value. Return -1 on error, 0 on success.

New in version 2.0.
*/
func PyModule_AddObject(self *PyObject, name string, value *PyObject) os.Error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return int2err(C.PyModule_AddObject(topy(self), c_name, topy(value)))
}

/*
int PyModule_AddIntConstant(PyObject *module, const char *name, long value)
Add an integer constant to module as name. This convenience function can be used from the module’s initialization function. Return -1 on error, 0 on success.

New in version 2.0.
*/
func PyModule_AddIntConstant(self *PyObject, name string, value int) os.Error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return int2err(C.PyModule_AddIntConstant(topy(self), c_name, C.long(value)))
}

/*
int PyModule_AddStringConstant(PyObject *module, const char *name, const char *value)
Add a string constant to module as name. This convenience function can be used from the module’s initialization function. The string value must be null-terminated. Return -1 on error, 0 on success.

New in version 2.0.
*/
func PyModule_AddStringConstant(self *PyObject, name, value string) os.Error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	return int2err(C.PyModule_AddStringConstant(topy(self), c_name, c_value))
}

/*
int PyModule_AddIntMacro(PyObject *module, macro)
Add an int constant to module. The name and the value are taken from macro. For example PyModule_AddConstant(module, AF_INET) adds the int constant AF_INET with the value of AF_INET to module. Return -1 on error, 0 on success.

New in version 2.6.
*/
func PyModule_AddIntMacro(self *PyObject, macro interface{}) os.Error {
	//FIXME ?
	panic("not implemented")
}

/*
int PyModule_AddStringMacro(PyObject *module, macro)
Add a string constant to module.
New in version 2.6.
*/
func PyModule_AddStringMacro(self *PyObject, macro interface{}) os.Error {
	//FIXME ?
	panic("not implemented")
}

///// class /////

/*
int PyClass_Check(PyObject *o)
Return true if the object o is a class object, including instances of types derived from the standard class object. Return false in all other cases.
*/
func PyClass_Check(o *PyObject) bool {
	return int2bool(C._gopy_PyClass_Check(topy(o)))
}

/*
int PyClass_IsSubclass(PyObject *klass, PyObject *base)
Return true if klass is a subclass of base. Return false in all other cases.
There are very few functions specific to instance objects.
*/
func PyClass_IsSubclass(klass, base *PyObject) bool {
	return int2bool(C.PyClass_IsSubclass(topy(klass), topy(base)))
}

///// class /////

/*
int PyInstance_Check(PyObject *obj)
Return true if obj is an instance.
*/
func PyInstance_Check(obj *PyObject) bool {
	return int2bool(C._gopy_PyInstance_Check(topy(obj)))
}

/*
PyObject* PyInstance_New(PyObject *class, PyObject *arg, PyObject *kw)
Return value: New reference.
Create a new instance of a specific class. The parameters arg and kw are used as the positional and keyword parameters to the object’s constructor.
*/
func PyInstance_New(class, arg, kw *PyObject) *PyObject {
	return togo(C.PyInstance_New(topy(class), topy(arg), topy(kw)))
}

/*
PyObject* PyInstance_NewRaw(PyObject *class, PyObject *dict)
Return value: New reference.
Create a new instance of a specific class without calling its constructor. class is the class of new object. The dict parameter will be used as the object’s __dict__; if NULL, a new dictionary will be created for the instance.
*/
func PyInstance_NewRaw(class, dict *PyObject) *PyObject {
	return togo(C.PyInstance_NewRaw(topy(class), topy(dict)))
}

// EOF
