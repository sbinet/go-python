package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
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
// EOF
