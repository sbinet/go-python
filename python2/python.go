package python2

//#cgo pkg-config: python-2.7
//#include "go-python.h"
import "C"

import (
	python "github.com/sbinet/go-python/runtime"
	"os"
	"unsafe"
)

var Runtime python.Runtime = py2Runtime{}

type py2Runtime struct{}

func (py2Runtime) GetVersion() string {
	return C.GoString(C.Py_GetVersion())
}

func (py2Runtime) GetCompiler() string {
	return C.GoString(C.Py_GetCompiler())
}

func (py2Runtime) Initialize(signals bool) {
	sig := 0
	if signals {
		sig = 1
	}
	C.Py_InitializeEx(C.int(sig))
}

func (py2Runtime) IsInitialized() bool {
	return C.Py_IsInitialized() != 0
}

func (py2Runtime) Finalize() {
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
func (py2Runtime) Main(args []string) int {
	argc := C.int(len(args))
	// no need to free. Py_Main takes ownership.
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	return int(C.Py_Main(argc, &argv[0]))
}

// RunString executes the Python source code from command in the __main__ module according to the flags argument.
// If __main__ does not already exist, it is created. Returns 0 on success or -1 if an exception was raised.
// If there was an error, there is no way to get the exception information.
func (py2Runtime) RunString(command string) int {
	cmd := C.CString(command)
	defer C.free(unsafe.Pointer(cmd))

	return int(C.PyRun_SimpleStringFlags(cmd, nil))
}

func (py2Runtime) RunFile(f *os.File) int {
	cname := C.CString(f.Name())
	defer C.free(unsafe.Pointer(cname))

	cmode := C.CString("r")
	defer C.free(unsafe.Pointer(cmode))

	cf := C.fdopen(C.int(f.Fd()), cmode)

	return int(C.PyRun_SimpleFileExFlags(cf, cname, 0, nil))
}
