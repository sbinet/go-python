// simplistic wrapper around the python C-API
package python

//#include "go-python.h"
import "C"

import (
	"fmt"
)

// PyGILState is the Go alias for the PyGILState_STATE enum
type PyGILState C.PyGILState_STATE

// PyThreadState layer
type PyThreadState struct {
	ptr *C.PyThreadState
}

// Initialize initializes the python interpreter and its GIL
func Initialize() error {
	// make sure the python interpreter has been initialized
	if C.Py_IsInitialized() == 0 {
		C.Py_Initialize()
	}
	if C.Py_IsInitialized() == 0 {
		return fmt.Errorf("python: could not initialize the python interpreter")
	}

	// make sure the GIL is correctly initialized
	if C.PyEval_ThreadsInitialized() == 0 {
		C.PyEval_InitThreads()
	}
	if C.PyEval_ThreadsInitialized() == 0 {
		return fmt.Errorf("python: could not initialize the GIL")
	}

	return nil
}

// Finalize shutdowns the python interpreter
func Finalize() error {
	C.Py_Finalize()
	return nil
}

// PyThreadState* PyEval_SaveThread()
// Release the global interpreter lock (if it has been created and thread
// support is enabled) and reset the thread state to NULL, returning the
// previous thread state (which is not NULL). If the lock has been created,
// the current thread must have acquired it. (This function is available even
// when thread support is disabled at compile time.)
func PyEval_SaveThread() *PyThreadState {
	state := C.PyEval_SaveThread()
	return &PyThreadState{ptr: state}
}

// void PyEval_RestoreThread(PyThreadState *tstate)
// Acquire the global interpreter lock (if it has been created and thread
// support is enabled) and set the thread state to tstate, which must not be
// NULL. If the lock has been created, the current thread must not have
// acquired it, otherwise deadlock ensues. (This function is available even
// when thread support is disabled at compile time.)
func PyEval_RestoreThread(state *PyThreadState) {
	C.PyEval_RestoreThread(state.ptr)
}

// Ensure that the current thread is ready to call the Python C API regardless
// of the current state of Python, or of the global interpreter lock. This may
// be called as many times as desired by a thread as long as each call is
// matched with a call to PyGILState_Release(). In general, other thread-related
// APIs may be used between PyGILState_Ensure() and PyGILState_Release() calls
// as long as the thread state is restored to its previous state before the
// Release(). For example, normal usage of the Py_BEGIN_ALLOW_THREADS and
// Py_END_ALLOW_THREADS macros is acceptable.
//
// The return value is an opaque “handle” to the thread state when
// PyGILState_Ensure() was called, and must be passed to PyGILState_Release()
// to ensure Python is left in the same state. Even though recursive calls are
// allowed, these handles cannot be shared - each unique call to
// PyGILState_Ensure() must save the handle for its call to PyGILState_Release().
//
// When the function returns, the current thread will hold the GIL and be able
// to call arbitrary Python code. Failure is a fatal error.
//
// New in version 2.3.
func PyGILState_Ensure() PyGILState {
	return PyGILState(C.PyGILState_Ensure())
}

// void PyGILState_Release(PyGILState_STATE)
// Release any resources previously acquired. After this call, Python’s state
// will be the same as it was prior to the corresponding PyGILState_Ensure()
// call (but generally this state will be unknown to the caller, hence the use
// of the GILState API).
//
// Every call to PyGILState_Ensure() must be matched by a call to
// PyGILState_Release() on the same thread.
//
// New in version 2.3.
func PyGILState_Release(state PyGILState) {
	C.PyGILState_Release(C.PyGILState_STATE(state))
}

// EOF
