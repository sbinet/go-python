package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
//int _gopy_PyRun_SimpleString(const char *command)
//{return PyRun_SimpleString(command);}
import "C"

import (
	"unsafe"
)

// very high level interface

// int Py_Main(int argc, char **argv)
// The main program for the standard interpreter. This is made available for programs which embed Python. The argc and argv parameters should be prepared exactly as those which are passed to a C program’s main() function. It is important to note that the argument list may be modified (but the contents of the strings pointed to by the argument list are not). The return value will be the integer passed to the sys.exit() function, 1 if the interpreter exits due to an exception, or 2 if the parameter list does not represent a valid Python command line.
//
// Note that if an otherwise unhandled SystemError is raised, this function will not return 1, but exit the process, as long as Py_InspectFlag is not set.
func Py_Main(args []string) int {
	var argc C.int = C.int(len(args))
	var argv []*C.char = make([]*C.char, argc)
	for idx, arg := range args {
		argv[idx] = C.CString(arg)
		// no need to free. Py_Main takes owner ship.
		//defer C.free(unsafe.Pointer(argv[idx]))
	}
	return int(C.Py_Main(argc, &argv[0]))
}

// int PyRun_SimpleString(const char *command)
// This is a simplified interface to PyRun_SimpleStringFlags() below, leaving the PyCompilerFlags* argument set to NULL.
func PyRun_SimpleString(command string) int {
	c_cmd := C.CString(command)
	defer C.free(unsafe.Pointer(c_cmd))
	return int(C._gopy_PyRun_SimpleString(c_cmd))
}

// EOF
