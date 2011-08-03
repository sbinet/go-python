// simplistic wrapper around the python C-API
package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
import "C"

func init() {
	// make sure the python interpreter has been initialized
	if C.Py_IsInitialized() == 0 {
		C.Py_Initialize()
	}
	if C.Py_IsInitialized() == 0 {
		panic("could not initialize the python interpreter")
	} else {
		//println("python interpreter initialized")
	}

	// make sure the GIL is correctly initialized
	if C.PyEval_ThreadsInitialized() == 0 {
		C.PyEval_InitThreads()
	}
	if C.PyEval_ThreadsInitialized() == 0 {
		panic("could not initialize the GIL")
	} else {
		//println("GIL initialized")
	}
}

// EOF
