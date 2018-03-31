package python3

//#cgo pkg-config: python-3.6
//#include "go-python.h"
import "C"

import (
	"github.com/sbinet/go-python"
	"os"
	"unsafe"
)

var Runtime python.Runtime = py3Runtime{}

type py3Runtime struct{}

func (py3Runtime) Initialize() {
	C.Py_Initialize()
}

func (py3Runtime) IsInitialized() bool {
	return C.Py_IsInitialized() != 0
}

func (py3Runtime) EvalInitThreads() {
	C.PyEval_InitThreads()
}

func (py3Runtime) EvalThreadsInitialized() bool {
	return C.PyEval_ThreadsInitialized() != 0
}

func (py3Runtime) Finalize() {
	C.Py_Finalize()
}

// Main runs the main program for the standard interpreter. This is made available for programs which embed Python.
// The args parameters should be prepared exactly as those which are passed to a C program’s main() function.
// It is important to note that the argument list may be modified (but the contents of the strings pointed to
// by the argument list are not). The return value will be the integer passed to the sys.exit() function,
// 1 if the interpreter exits due to an exception, or 2 if the parameter list does not represent a valid Python command line.
//
// Note that if an otherwise unhandled SystemError is raised, this function will not return 1, but exit the process,
// as long as Py_InspectFlag is not set.
func (py3Runtime) Main(args []string) int {
	argc := C.int(len(args))
	// no need to free. Py_Main takes ownership.
	argv := make([]*C.wchar_t, argc)
	for i, arg := range args {
		argv[i], _ = stringToWcharT(arg)
	}
	return int(C.Py_Main(argc, &argv[0]))
}

// RunString executes the Python source code from command in the __main__ module according to the flags argument.
// If __main__ does not already exist, it is created. Returns 0 on success or -1 if an exception was raised.
// If there was an error, there is no way to get the exception information.
func (py3Runtime) RunString(command string) int {
	cmd := C.CString(command)
	defer C.free(unsafe.Pointer(cmd))

	return int(C.PyRun_SimpleStringFlags(cmd, nil))
}

func (py3Runtime) RunFile(f *os.File) int {
	cname := C.CString(f.Name())
	defer C.free(unsafe.Pointer(cname))

	cmode := C.CString("r")
	defer C.free(unsafe.Pointer(cmode))

	cf := C.fdopen(C.int(f.Fd()), cmode)

	return int(C.PyRun_SimpleFileExFlags(cf, cname, 0, nil))
}
