// simplistic wrapper around the python C-API
package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
import "C"
import "unsafe"

func init() {
	// make sure the python interpreter has been initialized
	if C.Py_IsInitialized() == 0 {
		C.Py_Initialize()
	}
	if C.Py_IsInitialized() == 0 {
		panic("could not initialize the python interpreter")
	}

	// make sure the GIL is correctly initialized
	if C.PyEval_ThreadsInitialized() == 0 {
		C.PyEval_InitThreads()
	}
	if C.PyEval_ThreadsInitialized() == 0 {
		panic("could not initialize the GIL")
	}
}

// 
/*
int Py_Main(int argc, char **argv)
The main program for the standard interpreter. This is made available for programs which embed Python. The argc and argv parameters should be prepared exactly as those which are passed to a C program’s main() function. It is important to note that the argument list may be modified (but the contents of the strings pointed to by the argument list are not). The return value will be the integer passed to the sys.exit() function, 1 if the interpreter exits due to an exception, or 2 if the parameter list does not represent a valid Python command line.

Note that if an otherwise unhandled SystemError is raised, this function will not return 1, but exit the process, as long as Py_InspectFlag is not set.
*/
func Main(args []string) int {
	var argc C.int = C.int(len(args))
	var argv []*C.char = make([]*C.char, argc)
	for idx,arg := range args {
		argv[idx] = C.CString(arg)
	}
	defer func() {
		for idx,_ := range argv {
			C.free(unsafe.Pointer(argv[idx]))
		}
	}()
	return int(C.Py_Main(argc, &argv[0]))
}

/*
int PyRun_SimpleString(const char *command)
This is a simplified interface to PyRun_SimpleStringFlags() below, leaving the PyCompilerFlags* argument set to NULL.
*/
func Run_SimpleString(command string) int {
	c_cmd := C.CString(command)
	defer C.free(unsafe.Pointer(c_cmd))
	return int(C.PyRun_SimpleString(c_cmd))
}

// PyObject layer
type PyObject C.PyObject

func (self *PyObject) topy() *C.PyObject {
	var o C.PyObject = C.PyObject(*self)
	return &o
}

// int PyObject_HasAttr(PyObject *o, PyObject *attr_name)
// Returns 1 if o has the attribute attr_name, and 0 otherwise. This is equivalent to the Python expression hasattr(o, attr_name). This function always succeeds.
func (self *PyObject) HasAttr(attr_name *PyObject) int {
	return int(C.PyObject_HasAttr(self.topy(), attr_name.topy()))
}

/*
int PyObject_HasAttrString(PyObject *o, const char *attr_name)
Returns 1 if o has the attribute attr_name, and 0 otherwise. This is equivalent to the Python expression hasattr(o, attr_name). This function always succeeds.
*/
func (self *PyObject) HasAttrString(attr_name string) int {
	c_attr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(c_attr_name))

	return int(C.PyObject_HasAttrString(self.topy(),c_attr_name))
}

// EOF

