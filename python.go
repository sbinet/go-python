// simplistic wrapper around the python C-API
package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
import "C"

import (
	"fmt"
)

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

// EOF
