package python

// #include "go-python.h"
// char *gopy_ProgName = NULL;
// char *gopy_PythonHome = NULL;
import "C"

import (
	"unsafe"
)

// Py_SetProgramName should be called before Py_Initialize() is called for
// the first time, if it is called at all.
// It tells the interpreter the value of the argv[0] argument to the main()
// function of the program. This is used by Py_GetPath() and some other
// functions below to find the Python run-time libraries relative to the
// interpreter executable. The default value is 'python'. The argument should
// point to a zero-terminated character string in static storage whose contents
// will not change for the duration of the program’s execution.
// No code in the Python interpreter will change the contents of this storage.
func Py_SetProgramName(name string) {
	C.free(unsafe.Pointer(C.gopy_ProgName))
	C.gopy_ProgName = C.CString(name)
	C.Py_SetProgramName(C.gopy_ProgName)
}

// Py_GetProgramName returns the program name set with Py_SetProgramName(),
// or the default.
// The returned string points into static storage; the caller should not
// modify its value.
func Py_GetProgramName() string {
	cname := C.Py_GetProgramName()
	return C.GoString(cname)
}

// PySys_SetArgv initializes the 'sys.argv' array in python.
func PySys_SetArgv(argv []string) {
	argc := C.int(len(argv))
	cargs := make([]*C.char, len(argv))
	for idx, arg := range argv {
		cargs[idx] = C.CString(arg)
		defer C.free(unsafe.Pointer(cargs[idx]))
	}
	C.PySys_SetArgv(argc, &cargs[0])
}

// Set the default “home” directory, that is, the location of the standard Python libraries. See PYTHONHOME for the meaning of the argument string.
//
// The argument should point to a zero-terminated character string in static storage whose contents will not change for the duration of the program’s execution. No code in the Python interpreter will change the contents of this storage.
func Py_SetPythonHome(home string) {
	C.free(unsafe.Pointer(C.gopy_PythonHome))
	C.gopy_PythonHome = C.CString(home)
	C.Py_SetPythonHome(C.gopy_PythonHome)
}

// Return the default “home”, that is, the value set by a previous call to Py_SetPythonHome(), or the value of the PYTHONHOME environment variable if it is set.
func Py_GetPythonHome() string {
	home := C.Py_GetPythonHome()
	return C.GoString(home)
}
