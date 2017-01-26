package python

// #include "go-python.h"
import "C"

import (
	"unsafe"
)

// void PyErr_PrintEx(int set_sys_last_vars)
// Print a standard traceback to sys.stderr and clear the error indicator. Call this function only when the error indicator is set. (Otherwise it will cause a fatal error!)
//
// If set_sys_last_vars is nonzero, the variables sys.last_type, sys.last_value and sys.last_traceback will be set to the type, value and traceback of the printed exception, respectively.
func PyErr_PrintEx(set_sys_last_vars bool) {
	c := bool2int(set_sys_last_vars)
	C.PyErr_PrintEx(c)
}

// void PyErr_Print()
// Alias for PyErr_PrintEx(1).
func PyErr_Print() {
	C.PyErr_Print()
}

// PyObject* PyErr_Occurred()
// Return value: Borrowed reference.
// Test whether the error indicator is set. If set, return the exception type (the first argument to the last call to one of the PyErr_Set*() functions or to PyErr_Restore()). If not set, return NULL. You do not own a reference to the return value, so you do not need to Py_DECREF() it.
//
// Note Do not compare the return value to a specific exception; use PyErr_ExceptionMatches() instead, shown below. (The comparison could easily fail since the exception may be an instance instead of a class, in the case of a class exception, or it may the a subclass of the expected exception.)
func PyErr_Occurred() *PyObject {
	return togo(C.PyErr_Occurred())
}

// int PyErr_ExceptionMatches(PyObject *exc)
// Equivalent to PyErr_GivenExceptionMatches(PyErr_Occurred(), exc). This should only be called when an exception is actually set; a memory access violation will occur if no exception has been raised.
func PyErr_ExceptionMatches(exc *PyObject) bool {
	return int2bool(C.PyErr_ExceptionMatches(topy(exc)))
}

// int PyErr_GivenExceptionMatches(PyObject *given, PyObject *exc)
// Return true if the given exception matches the exception in exc. If exc is a class object, this also returns true when given is an instance of a subclass. If exc is a tuple, all exceptions in the tuple (and recursively in subtuples) are searched for a match.
func PyErr_GivenExceptionMatches(given, exc *PyObject) bool {
	return int2bool(C.PyErr_GivenExceptionMatches(topy(given), topy(exc)))
}

// void PyErr_NormalizeException(PyObject**exc, PyObject**val, PyObject**tb)
// Under certain circumstances, the values returned by PyErr_Fetch() below can be “unnormalized”, meaning that *exc is a class object but *val is not an instance of the same class. This function can be used to instantiate the class in that case. If the values are already normalized, nothing happens. The delayed normalization is implemented to improve performance.
func PyErr_NormalizeException(exc, val, tb *PyObject) (*PyObject, *PyObject, *PyObject) {
	C.PyErr_NormalizeException(&exc.ptr, &val.ptr, &tb.ptr)
	return exc, val, tb
}

// void PyErr_Clear()
// Clear the error indicator. If the error indicator is not set, there is no effect.
func PyErr_Clear() {
	C.PyErr_Clear()
}

// void PyErr_Fetch(PyObject **ptype, PyObject **pvalue, PyObject **ptraceback)
// Retrieve the error indicator into three variables whose addresses are passed. If the error indicator is not set, set all three variables to NULL. If it is set, it will be cleared and you own a reference to each object retrieved. The value and traceback object may be NULL even when the type object is not.
//
// Note This function is normally only used by code that needs to handle exceptions or by code that needs to save and restore the error indicator temporarily.
func PyErr_Fetch() (exc, val, tb *PyObject) {
	exc = &PyObject{}
	val = &PyObject{}
	tb = &PyObject{}

	C.PyErr_Fetch(&exc.ptr, &val.ptr, &tb.ptr)
	return
}

// void PyErr_Restore(PyObject *type, PyObject *value, PyObject *traceback)
// Set the error indicator from the three objects. If the error indicator is already set, it is cleared first. If the objects are NULL, the error indicator is cleared. Do not pass a NULL type and non-NULL value or traceback. The exception type should be a class. Do not pass an invalid exception type or value. (Violating these rules will cause subtle problems later.) This call takes away a reference to each object: you must own a reference to each object before the call and after the call you no longer own these references. (If you don’t understand this, don’t use this function. I warned you.)
//
// Note This function is normally only used by code that needs to save and restore the error indicator temporarily; use PyErr_Fetch() to save the current exception state.
func PyErr_Restore(typ, value, traceback *PyObject) {
	C.PyErr_Restore(topy(typ), topy(value), topy(traceback))
}

// void PyErr_SetString(PyObject *type, const char *message)
// This is the most common way to set the error indicator. The first argument specifies the exception type; it is normally one of the standard exceptions, e.g. PyExc_RuntimeError. You need not increment its reference count. The second argument is an error message; it is converted to a string object.
func PyErr_SetString(typ *PyObject, message string) {
	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))
	C.PyErr_SetString(topy(typ), c_message)
}

// void PyErr_SetObject(PyObject *type, PyObject *value)
// This function is similar to PyErr_SetString() but lets you specify an arbitrary Python object for the “value” of the exception.
func PyErr_SetObject(typ, value *PyObject) {
	C.PyErr_SetObject(topy(typ), topy(value))
}

// PyObject* PyErr_Format(PyObject *exception, const char *format, ...)
// Return value: Always NULL.
// This function sets the error indicator and returns NULL. exception should be a Python exception class. The format and subsequent parameters help format the error message; they have the same meaning and values as in PyString_FromFormat().
func PyErr_Format(exception *PyObject, format string, args ...interface{}) *PyObject {
	//FIXME
	panic("not implemented")
}

// void PyErr_SetNone(PyObject *type)
// This is a shorthand for PyErr_SetObject(type, Py_None).
func PyErr_SetNone(typ *PyObject) {
	C.PyErr_SetNone(topy(typ))
}

// int PyErr_BadArgument()
// This is a shorthand for PyErr_SetString(PyExc_TypeError, message), where message indicates that a built-in operation was invoked with an illegal argument. It is mostly for internal use.
func PyErr_BadArgument() bool {
	return int2bool(C.PyErr_BadArgument())
}

// PyObject* PyErr_NoMemory()
// Return value: Always NULL.
// This is a shorthand for PyErr_SetNone(PyExc_MemoryError); it returns NULL so an object allocation function can write return PyErr_NoMemory(); when it runs out of memory.
func PyErr_NoMemory() *PyObject {
	//FIXME is this the right thing to do ?
	//      should we integrate better with go panic/recover ?
	return togo(C.PyErr_NoMemory())
}

// PyObject* PyErr_SetFromErrno(PyObject *type)
// Return value: Always NULL.
// This is a convenience function to raise an exception when a C library function has returned an error and set the C variable errno. It constructs a tuple object whose first item is the integer errno value and whose second item is the corresponding error message (gotten from strerror()), and then calls PyErr_SetObject(type, object). On Unix, when the errno value is EINTR, indicating an interrupted system call, this calls PyErr_CheckSignals(), and if that set the error indicator, leaves it set to that. The function always returns NULL, so a wrapper function around a system call can write return PyErr_SetFromErrno(type); when the system call returns an error.
func PyErr_SetFromErrno(typ *PyObject) *PyObject {
	//FIXME is this the right thing to do ?
	//      should we integrate better with go panic/recover ?
	return togo(C.PyErr_SetFromErrno(topy(typ)))
}

// PyObject* PyErr_SetFromErrnoWithFilename(PyObject *type, const char *filename)
// Return value: Always NULL.
// Similar to PyErr_SetFromErrno(), with the additional behavior that if filename is not NULL, it is passed to the constructor of type as a third parameter. In the case of exceptions such as IOError and OSError, this is used to define the filename attribute of the exception instance.
func PyErr_SetFromErrnoWithFilename(typ *PyObject, filename string) *PyObject {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	//FIXME is this the right thing to do ?
	//      should we integrate better with go panic/recover ?
	return togo(C.PyErr_SetFromErrnoWithFilename(topy(typ), c_filename))
}

// PyObject* PyErr_SetFromWindowsErr(int ierr)
// Return value: Always NULL.
// This is a convenience function to raise WindowsError. If called with ierr of 0, the error code returned by a call to GetLastError() is used instead. It calls the Win32 function FormatMessage() to retrieve the Windows description of error code given by ierr or GetLastError(), then it constructs a tuple object whose first item is the ierr value and whose second item is the corresponding error message (gotten from FormatMessage()), and then calls PyErr_SetObject(PyExc_WindowsError, object). This function always returns NULL. Availability: Windows.
func PyErr_SetFromWindowsErr(ierr bool) *PyObject {
	c_ierr := bool2int(ierr)
	return togo(C.PyErr_SetFromWindowsErr(c_ierr))
}

// PyObject* PyErr_SetExcFromWindowsErr(PyObject *type, int ierr)
// Return value: Always NULL.
// Similar to PyErr_SetFromWindowsErr(), with an additional parameter specifying the exception type to be raised. Availability: Windows.
//
// New in version 2.3.
func PyErr_SetExcFromWindowsErr(typ *PyObject, ierr bool) *PyObject {
	c_ierr := bool2int(ierr)
	return togo(C.PyErr_SetExcFromWindowsErr(topy(typ), c_ierr))
}

// PyObject* PyErr_SetFromWindowsErrWithFilename(int ierr, const char *filename)
// Return value: Always NULL.
// Similar to PyErr_SetFromWindowsErr(), with the additional behavior that if filename is not NULL, it is passed to the constructor of WindowsError as a third parameter. Availability: Windows.
func PyErr_SetFromWindowsErrWithFilename(ierr bool, filename string) *PyObject {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	c_ierr := bool2int(ierr)
	return togo(C.PyErr_SetFromWindowsErrWithFilename(c_ierr, c_filename))
}

// PyObject* PyErr_SetExcFromWindowsErrWithFilename(PyObject *type, int ierr, char *filename)
// Return value: Always NULL.
// Similar to PyErr_SetFromWindowsErrWithFilename(), with an additional parameter specifying the exception type to be raised. Availability: Windows.
//
// New in version 2.3.
func PyErr_SetExcFromWindowsErrWithFilename(typ *PyObject, ierr bool, filename string) *PyObject {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	c_ierr := bool2int(ierr)
	return togo(C.PyErr_SetExcFromWindowsErrWithFilename(topy(typ), c_ierr, c_filename))
}

// void PyErr_BadInternalCall()
// This is a shorthand for PyErr_SetString(PyExc_SystemError, message), where message indicates that an internal operation (e.g. a Python/C API function) was invoked with an illegal argument. It is mostly for internal use.
func PyErr_BadInternalCall() {
	C.PyErr_BadInternalCall()
}

// int PyErr_WarnEx(PyObject *category, char *message, int stacklevel)
// Issue a warning message. The category argument is a warning category (see below) or NULL; the message argument is a message string. stacklevel is a positive number giving a number of stack frames; the warning will be issued from the currently executing line of code in that stack frame. A stacklevel of 1 is the function calling PyErr_WarnEx(), 2 is the function above that, and so forth.
//
// This function normally prints a warning message to sys.stderr; however, it is also possible that the user has specified that warnings are to be turned into errors, and in that case this will raise an exception. It is also possible that the function raises an exception because of a problem with the warning machinery (the implementation imports the warnings module to do the heavy lifting). The return value is 0 if no exception is raised, or -1 if an exception is raised. (It is not possible to determine whether a warning message is actually printed, nor what the reason is for the exception; this is intentional.) If an exception is raised, the caller should do its normal exception handling (for example, Py_DECREF() owned references and return an error value).
//
// Warning categories must be subclasses of Warning; the default warning category is RuntimeWarning. The standard Python warning categories are available as global variables whose names are PyExc_ followed by the Python exception name. These have the type PyObject*; they are all class objects. Their names are PyExc_Warning, PyExc_UserWarning, PyExc_UnicodeWarning, PyExc_DeprecationWarning, PyExc_SyntaxWarning, PyExc_RuntimeWarning, and PyExc_FutureWarning. PyExc_Warning is a subclass of PyExc_Exception; the other warning categories are subclasses of PyExc_Warning.
//
// For information about warning control, see the documentation for the warnings module and the -W option in the command line documentation. There is no C API for warning control.
func PyErr_WarnEx(category *PyObject, message string, stacklevel int) error {
	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	return int2err(C.PyErr_WarnEx(topy(category), c_message, C.Py_ssize_t(stacklevel)))
}

/*
int PyErr_Warn(PyObject *category, char *message)
Issue a warning message. The category argument is a warning category (see below) or NULL; the message argument is a message string. The warning will appear to be issued from the function calling PyErr_Warn(), equivalent to calling PyErr_WarnEx() with a stacklevel of 1.

Deprecated; use PyErr_WarnEx() instead.
*/

// int PyErr_WarnExplicit(PyObject *category, const char *message, const char *filename, int lineno, const char *module, PyObject *registry)
// Issue a warning message with explicit control over all warning attributes. This is a straightforward wrapper around the Python function warnings.warn_explicit(), see there for more information. The module and registry arguments may be set to NULL to get the default effect described there.
func PyErr_WarnExplicit(category *PyObject, message, filename string, lineno int, module string, registry *PyObject) error {
	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_module := C.CString(module)
	defer C.free(unsafe.Pointer(c_module))

	return int2err(C.PyErr_WarnExplicit(topy(category), c_message, c_filename,
		C.int(lineno), c_module, topy(registry)))
}

// int PyErr_WarnPy3k(char *message, int stacklevel)
// Issue a DeprecationWarning with the given message and stacklevel if the Py_Py3kWarningFlag flag is enabled.
// New in version 2.6.
func PyErr_WarnPy3k(message string, stacklevel int) error {
	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	return int2err(C._gopy_PyErr_WarnPy3k(c_message, C.int(stacklevel)))
}

// int PyErr_CheckSignals()
// This function interacts with Python’s signal handling. It checks whether a signal has been sent to the processes and if so, invokes the corresponding signal handler. If the signal module is supported, this can invoke a signal handler written in Python. In all cases, the default effect for SIGINT is to raise the KeyboardInterrupt exception. If an exception is raised the error indicator is set and the function returns -1; otherwise the function returns 0. The error indicator may or may not be cleared if it was previously set.
func PyErr_CheckSignals() bool {
	return int2bool(C.PyErr_CheckSignals())
}

// void PyErr_SetInterrupt()
// This function simulates the effect of a SIGINT signal arriving — the next time PyErr_CheckSignals() is called, KeyboardInterrupt will be raised. It may be called without holding the interpreter lock.
func PyErr_SetInterrupt() {
	C.PyErr_SetInterrupt()
}

// PyObject* PyErr_NewException(char *name, PyObject *base, PyObject *dict)
// Return value: New reference.
// This utility function creates and returns a new exception object. The name argument must be the name of the new exception, a C string of the form module.class. The base and dict arguments are normally NULL. This creates a class object derived from Exception (accessible in C as PyExc_Exception).
//
// The __module__ attribute of the new class is set to the first part (up to the last dot) of the name argument, and the class name is set to the last part (after the last dot). The base argument can be used to specify alternate base classes; it can either be only one class or a tuple of classes. The dict argument can be used to specify a dictionary of class variables and methods.
func PyErr_NewException(name string, base, dict *PyObject) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return togo(C.PyErr_NewException(c_name, topy(base), topy(dict)))
}

// PyObject* PyErr_NewExceptionWithDoc(char *name, char *doc, PyObject *base, PyObject *dict)
// Return value: New reference.
// Same as PyErr_NewException(), except that the new exception class can easily be given a docstring: If doc is non-NULL, it will be used as the docstring for the exception class.
//
// New in version 2.7.
func PyErr_NewExceptionWithDoc(name, doc string, base, dict *PyObject) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_doc := C.CString(doc)
	defer C.free(unsafe.Pointer(c_doc))

	return togo(C.PyErr_NewExceptionWithDoc(c_name, c_doc, topy(base), topy(dict)))
}

// void PyErr_WriteUnraisable(PyObject *obj)
// This utility function prints a warning message to sys.stderr when an exception has been set but it is impossible for the interpreter to actually raise the exception. It is used, for example, when an exception occurs in an __del__() method.
// The function is called with a single argument obj that identifies the context in which the unraisable exception occurred. The repr of obj will be printed in the warning message.
func PyErr_WriteUnraisable(obj *PyObject) {
	C.PyErr_WriteUnraisable(topy(obj))
}

///// exception instances /////

var (
	PyExc_BaseException       = togo(C.PyExc_BaseException)
	PyExc_Exception           = togo(C.PyExc_Exception)
	PyExc_StandardError       = togo(C.PyExc_StandardError)
	PyExc_ArithmeticError     = togo(C.PyExc_ArithmeticError)
	PyExc_LookupError         = togo(C.PyExc_LookupError)
	PyExc_AssertionError      = togo(C.PyExc_AssertionError)
	PyExc_AttributeError      = togo(C.PyExc_AttributeError)
	PyExc_EOFError            = togo(C.PyExc_EOFError)
	PyExc_EnvironmentError    = togo(C.PyExc_EnvironmentError)
	PyExc_FloatingPointError  = togo(C.PyExc_FloatingPointError)
	PyExc_IOError             = togo(C.PyExc_IOError)
	PyExc_ImportError         = togo(C.PyExc_ImportError)
	PyExc_IndexError          = togo(C.PyExc_IndexError)
	PyExc_KeyError            = togo(C.PyExc_KeyError)
	PyExc_KeyboardInterrupt   = togo(C.PyExc_KeyboardInterrupt)
	PyExc_MemoryError         = togo(C.PyExc_MemoryError)
	PyExc_NameError           = togo(C.PyExc_NameError)
	PyExc_NotImplementedError = togo(C.PyExc_NotImplementedError)
	PyExc_OSError             = togo(C.PyExc_OSError)
	PyExc_OverflowError       = togo(C.PyExc_OverflowError)
	PyExc_ReferenceError      = togo(C.PyExc_ReferenceError)
	PyExc_RuntimeError        = togo(C.PyExc_RuntimeError)
	PyExc_SyntaxError         = togo(C.PyExc_SyntaxError)
	PyExc_SystemError         = togo(C.PyExc_SystemError)
	PyExc_SystemExit          = togo(C.PyExc_SystemExit)
	PyExc_TypeError           = togo(C.PyExc_TypeError)
	PyExc_ValueError          = togo(C.PyExc_ValueError)
	//FIXME: this should go into an exceptions_windows.go file
	//PyExc_WindowsError = togo(C.PyExc_WindowsError)

	PyExc_ZeroDivisionError = togo(C.PyExc_ZeroDivisionError)
)

// EOF
