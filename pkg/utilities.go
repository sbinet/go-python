package python

/*
#include "Python.h"
#include <stdlib.h>
#include <string.h>

 PyObject* _gopy_PyImport_ImportModuleEx(char *name, PyObject *globals, PyObject *locals, PyObject *fromlist) { return PyImport_ImportModuleEx(name, globals, locals, fromlist); }

*/
import "C"
import "unsafe"
import "os"

////// Operating System Utilities //////

/*
int Py_FdIsInteractive(FILE *fp, const char *filename)
Return true (nonzero) if the standard I/O file fp with name filename is deemed interactive. This is the case for files for which isatty(fileno(fp)) is true. If the global flag Py_InteractiveFlag is true, this function also returns true if the filename pointer is NULL or if the name is equal to one of the strings '<stdin>' or '???'.
*/
func Py_FdIsInteractive(fp *C.FILE, fname string) bool {
	c_fname := C.CString(fname)
	defer C.free(unsafe.Pointer(c_fname))
	return int2bool(C.Py_FdIsInteractive(fp,c_fname))
}

/*
void PyOS_AfterFork()
Function to update some internal state after a process fork; this should be called in the new process if the Python interpreter will continue to be used. If a new executable is loaded into the new process, this function does not need to be called.
*/
func PyOS_AfterFork() {
	C.PyOS_AfterFork()
}

/*
int PyOS_CheckStack()
Return true when the interpreter runs out of stack space. This is a reliable check, but is only available when USE_STACKCHECK is defined (currently on Windows using the Microsoft Visual C++ compiler). USE_STACKCHECK will be defined automatically; you should never change the definition in your own code.
*/
// func PyOS_CheckStack() bool {
// 	return int2bool(C._gopy_PyOS_CheckStack())
// }

/*
PyOS_sighandler_t PyOS_getsig(int i)
Return the current signal handler for signal i. This is a thin wrapper around either sigaction() or signal(). Do not call those functions directly! PyOS_sighandler_t is a typedef alias for void (*)(int).
*/

/*
PyOS_sighandler_t PyOS_setsig(int i, PyOS_sighandler_t h)
Set the signal handler for signal i to be h; return the old signal handler. This is a thin wrapper around either sigaction() or signal(). Do not call those functions directly! PyOS_sighandler_t is a typedef alias for void (*)(int).
*/


/////// Process Control /////////

/*
void Py_FatalError(const char *message)
Print a fatal error message and kill the process. No cleanup is performed. This function should only be invoked when a condition is detected that would make it dangerous to continue using the Python interpreter; e.g., when the object administration appears to be corrupted. On Unix, the standard C library function abort() is called which will attempt to produce a core file.
*/
func Py_FatalError(message string) {
	c_msg := C.CString(message)
	defer C.free(unsafe.Pointer(c_msg))
	C.Py_FatalError(c_msg)
}

/*
void Py_Exit(int status)
Exit the current process. This calls Py_Finalize() and then calls the standard C library function exit(status).
*/
func Py_Exit(status int) {
	C.Py_Exit(C.int(status))
}

var atexit_funcs []func()

/*
int Py_AtExit(void (*func) ())
Register a cleanup function to be called by Py_Finalize(). The cleanup function will be called with no arguments and should return no value. At most 32 cleanup functions can be registered. When the registration is successful, Py_AtExit() returns 0; on failure, it returns -1. The cleanup function registered last is called first. Each cleanup function will be called at most once. Since Python’s internal finalization will have completed before the cleanup function, no Python APIs should be called by func.
*/
func Py_AtExit(fct func()) os.Error {
	atexit_funcs = append(atexit_funcs, fct)
	//c_fct := 
	// FIXME
	panic("not implemented")
	return os.NewError("C<->go callbacks are hard")
}

///// import /////

/*
PyObject* PyImport_ImportModule(const char *name)
Return value: New reference.
This is a simplified interface to PyImport_ImportModuleEx() below, leaving the globals and locals arguments set to NULL and level set to 0. When the name argument contains a dot (when it specifies a submodule of a package), the fromlist argument is set to the list ['*'] so that the return value is the named module rather than the top-level package containing it as would otherwise be the case. (Unfortunately, this has an additional side effect when name in fact specifies a subpackage instead of a submodule: the submodules specified in the package’s __all__ variable are loaded.) Return a new reference to the imported module, or NULL with an exception set on failure. Before Python 2.4, the module may still be created in the failure case — examine sys.modules to find out. Starting with Python 2.4, a failing import of a module no longer leaves the module in sys.modules.

Changed in version 2.4: Failing imports remove incomplete module objects.

Changed in version 2.6: Always uses absolute imports.
*/
func PyImport_ImportModule(name string) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return togo(C.PyImport_ImportModule(c_name))
}

/*
PyObject* PyImport_ImportModuleNoBlock(const char *name)
This version of PyImport_ImportModule() does not block. It’s intended to be used in C functions that import other modules to execute a function. The import may block if another thread holds the import lock. The function PyImport_ImportModuleNoBlock() never blocks. It first tries to fetch the module from sys.modules and falls back to PyImport_ImportModule() unless the lock is held, in which case the function will raise an ImportError.

New in version 2.6.
*/
func PyImport_ImportModuleNoBlock(name string) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return togo(C.PyImport_ImportModuleNoBlock(c_name))
}


/*
PyObject* PyImport_ImportModuleEx(char *name, PyObject *globals, PyObject *locals, PyObject *fromlist)
Return value: New reference.
Import a module. This is best described by referring to the built-in Python function __import__(), as the standard __import__() function calls this function directly.

The return value is a new reference to the imported module or top-level package, or NULL with an exception set on failure (before Python 2.4, the module may still be created in this case). Like for __import__(), the return value when a submodule of a package was requested is normally the top-level package, unless a non-empty fromlist was given.

Changed in version 2.4: Failing imports remove incomplete module objects.

Changed in version 2.6: The function is an alias for PyImport_ImportModuleLevel() with -1 as level, meaning relative import.
*/
func PyImport_ImportModuleEx(name string, globals, locals, fromlist *PyObject) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return togo(C._gopy_PyImport_ImportModuleEx(c_name, topy(globals), topy(locals), topy(fromlist)))
}


/*
PyObject* PyImport_ImportModuleLevel(char *name, PyObject *globals, PyObject *locals, PyObject *fromlist, int level)
Return value: New reference.
Import a module. This is best described by referring to the built-in Python function __import__(), as the standard __import__() function calls this function directly.

The return value is a new reference to the imported module or top-level package, or NULL with an exception set on failure. Like for __import__(), the return value when a submodule of a package was requested is normally the top-level package, unless a non-empty fromlist was given.

New in version 2.5.
*/
func PyImport_ImportModuleLevel(name string, globals, locals, fromlist *PyObject, level int) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return togo(C.PyImport_ImportModuleLevel(c_name, topy(globals), topy(locals), topy(fromlist), C.int(level)))
}

/*
PyObject* PyImport_Import(PyObject *name)
Return value: New reference.
This is a higher-level interface that calls the current “import hook function”. It invokes the __import__() function from the __builtins__ of the current globals. This means that the import is done using whatever import hooks are installed in the current environment, e.g. by rexec or ihooks.

Changed in version 2.6: Always uses absolute imports.
*/
func PyImport_Import(name *PyObject) *PyObject {
	return togo(C.PyImport_Import(topy(name)))
}

/*
PyObject* PyImport_ReloadModule(PyObject *m)
Return value: New reference.
Reload a module. This is best described by referring to the built-in Python function reload(), as the standard reload() function calls this function directly. Return a new reference to the reloaded module, or NULL with an exception set on failure (the module still exists in this case).
*/
func PyImport_ReloadModule(m *PyObject) *PyObject {
	return togo(C.PyImport_ReloadModule(topy(m)))
}

/*
PyObject* PyImport_AddModule(const char *name)
Return value: Borrowed reference.
Return the module object corresponding to a module name. The name argument may be of the form package.module. First check the modules dictionary if there’s one there, and if not, create a new one and insert it in the modules dictionary. Return NULL with an exception set on failure.

Note This function does not load or import the module; if the module wasn’t already loaded, you will get an empty module object. Use PyImport_ImportModule() or one of its variants to import a module. Package structures implied by a dotted name for name are not created if not already present.
*/
func PyImport_AddModule(name string) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return togo(C.PyImport_AddModule(c_name))
}

/*
PyObject* PyImport_ExecCodeModule(char *name, PyObject *co)
Return value: New reference.
Given a module name (possibly of the form package.module) and a code object read from a Python bytecode file or obtained from the built-in function compile(), load the module. Return a new reference to the module object, or NULL with an exception set if an error occurred. Before Python 2.4, the module could still be created in error cases. Starting with Python 2.4, name is removed from sys.modules in error cases, and even if name was already in sys.modules on entry to PyImport_ExecCodeModule(). Leaving incompletely initialized modules in sys.modules is dangerous, as imports of such modules have no way to know that the module object is an unknown (and probably damaged with respect to the module author’s intents) state.

The module’s __file__ attribute will be set to the code object’s co_filename.

This function will reload the module if it was already imported. See PyImport_ReloadModule() for the intended way to reload a module.

If name points to a dotted name of the form package.module, any package structures not already created will still not be created.

Changed in version 2.4: name is removed from sys.modules in error cases.
*/
func PyImport_ExecCodeModule(name string, co *PyObject) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return togo(C.PyImport_ExecCodeModule(c_name, topy(co)))
}

/*
PyObject* PyImport_ExecCodeModuleEx(char *name, PyObject *co, char *pathname)
Return value: New reference.
Like PyImport_ExecCodeModule(), but the __file__ attribute of the module object is set to pathname if it is non-NULL.
*/
func PyImport_ExecCodeModuleEx(name string, co *PyObject, pathname string) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_pname := C.CString(pathname)
	defer C.free(unsafe.Pointer(c_pname))

	return togo(C.PyImport_ExecCodeModuleEx(c_name, topy(co), c_pname))
}


/*
long PyImport_GetMagicNumber()
Return the magic number for Python bytecode files (a.k.a. .pyc and .pyo files). The magic number should be present in the first four bytes of the bytecode file, in little-endian byte order.
*/
func PyImport_GetMagicNumber() int64 {
	return int64(C.PyImport_GetMagicNumber())
}

/*
PyObject* PyImport_GetModuleDict()
Return value: Borrowed reference.
Return the dictionary used for the module administration (a.k.a. sys.modules). Note that this is a per-interpreter variable.
*/
func PyImport_GetModuleDict() *PyObject {
	return togo(C.PyImport_GetModuleDict())
}

/*
PyObject* PyImport_GetImporter(PyObject *path)
Return an importer object for a sys.path/pkg.__path__ item path, possibly by fetching it from the sys.path_importer_cache dict. If it wasn’t yet cached, traverse sys.path_hooks until a hook is found that can handle the path item. Return None if no hook could; this tells our caller it should fall back to the built-in import mechanism. Cache the result in sys.path_importer_cache. Return a new reference to the importer object.

New in version 2.6.
*/
func PyImport_GetImporter(path *PyObject) *PyObject {
	return togo(C.PyImport_GetImporter(topy(path)))
}

/*
void _PyImport_Init()
Initialize the import mechanism. For internal use only.
void PyImport_Cleanup()
Empty the module table. For internal use only.
void _PyImport_Fini()
Finalize the import mechanism. For internal use only.
PyObject* _PyImport_FindExtension(char *, char *)
For internal use only.
PyObject* _PyImport_FixupExtension(char *, char *)
For internal use only.
*/

/*
int PyImport_ImportFrozenModule(char *name)
Load a frozen module named name. Return 1 for success, 0 if the module is not found, and -1 with an exception set if the initialization failed. To access the imported module on a successful load, use PyImport_ImportModule(). (Note the misnomer — this function would reload the module if it was already imported.)
*/
func PyImport_ImportFrozenModule(name string) os.Error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return int2err(C.PyImport_ImportFrozenModule(c_name))
}

/*
struct _frozen
This is the structure type definition for frozen module descriptors, as generated by the freeze utility (see Tools/freeze/ in the Python source distribution). Its definition, found in Include/import.h, is:

struct _frozen {
    char *name;
    unsigned char *code;
    int size;
};
struct _frozen* PyImport_FrozenModules
This pointer is initialized to point to an array of struct _frozen records, terminated by one whose members are all NULL or zero. When a frozen module is imported, it is searched in this table. Third-party code could play tricks with this to provide a dynamically created collection of frozen modules.
int PyImport_AppendInittab(const char *name, void (*initfunc)(void))
Add a single module to the existing table of built-in modules. This is a convenience wrapper around PyImport_ExtendInittab(), returning -1 if the table could not be extended. The new module can be imported by the name name, and uses the function initfunc as the initialization function called on the first attempted import. This should be called before Py_Initialize().
struct _inittab
Structure describing a single entry in the list of built-in modules. Each of these structures gives the name and initialization function for a module built into the interpreter. Programs which embed Python may use an array of these structures in conjunction with PyImport_ExtendInittab() to provide additional built-in modules. The structure is defined in Include/import.h as:

struct _inittab {
    char *name;
    void (*initfunc)(void);
};
int PyImport_ExtendInittab(struct _inittab *newtab)
Add a collection of modules to the table of built-in modules. The newtab array must end with a sentinel entry which contains NULL for the name field; failure to provide the sentinel value can result in a memory fault. Returns 0 on success or -1 if insufficient memory could be allocated to extend the internal table. In the event of failure, no modules are added to the internal table. This should be called before Py_Initialize().
*/

// EOF
