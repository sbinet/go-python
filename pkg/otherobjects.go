package python

/*
#include "Python.h"
#include <stdlib.h>
#include <string.h>

int _gopy_PyModule_Check(PyObject *p) { return PyModule_Check(p); }
int _gopy_PyModule_CheckExact(PyObject *p) { return PyModule_CheckExact(p); }

*/
import "C"
//import "unsafe"

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

PyObject* PyModule_GetDict(PyObject *module)
Return value: Borrowed reference.
Return the dictionary object that implements module‘s namespace; this object is the same as the __dict__ attribute of the module object. This function never fails. It is recommended extensions use other PyModule_*() and PyObject_*() functions rather than directly manipulate a module’s __dict__.

char* PyModule_GetName(PyObject *module)
Return module‘s __name__ value. If the module does not provide one, or if it is not a string, SystemError is raised and NULL is returned.

char* PyModule_GetFilename(PyObject *module)
Return the name of the file from which module was loaded using module‘s __file__ attribute. If this is not defined, or if it is not a string, raise SystemError and return NULL.

int PyModule_AddObject(PyObject *module, const char *name, PyObject *value)
Add an object to module as name. This is a convenience function which can be used from the module’s initialization function. This steals a reference to value. Return -1 on error, 0 on success.

New in version 2.0.

int PyModule_AddIntConstant(PyObject *module, const char *name, long value)
Add an integer constant to module as name. This convenience function can be used from the module’s initialization function. Return -1 on error, 0 on success.

New in version 2.0.

int PyModule_AddStringConstant(PyObject *module, const char *name, const char *value)
Add a string constant to module as name. This convenience function can be used from the module’s initialization function. The string value must be null-terminated. Return -1 on error, 0 on success.

New in version 2.0.

int PyModule_AddIntMacro(PyObject *module, macro)
Add an int constant to module. The name and the value are taken from macro. For example PyModule_AddConstant(module, AF_INET) adds the int constant AF_INET with the value of AF_INET to module. Return -1 on error, 0 on success.

New in version 2.6.

int PyModule_AddStringMacro(PyObject *module, macro)
Add a string constant to module.
New in version 2.6.
*/

// EOF
